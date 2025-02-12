package Config

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"govenv/pkg/common/PyHelpers"

	"gopkg.in/ini.v1"
)

var notUpdatableConfigProperties = []string{"COMMON_FILE_PATH", "Name", "MemConfig", "MemConfigPath", "logger", "Parser", "ParentOnMemConfUpdate", "withConfigServer", "Handler", "loggerLog"}

var configTemplateTest string = `
[COMMON]

# Config name
NAME = common
MAIN_WAIT_BEAT = 0.01

# Log server
LG_IP = 127.0.0.1
LG_PORT = 9020

# Config server
CF_IP = 127.0.0.2
CF_PORT = 1026
CF_REFRESH = 300

# Telegram Notifs
NT_IP = 127.0.0.3
NT_PORT = 10080

# Telebot Remote Control
TB_TOKEN = !!! FILL ME TOKEN !!!
TB_CHATID = !!! FILL ME CHATID !!!
TB_IP = 127.0.0.4
TB_PORT = 31337

# Scheduler
SC_IP = 127.0.0.5
SC_PORT = 5001

# Monitoring
MT_IP = 127.0.0.6
MT_PORT = 5000

# Matrix Q
AE_IP = 127.0.0.7
AE_INTPORT = 9002
AE_EXTPORT = 9003

# Database
DB_IP = 127.0.0.8
DB_PORT = 5120
DB_ENDPOINT = http://db:5220
DB_SSLCERT = false
DB_BACKEND = ./backend.db

# datas on file system
FS_TEMP = ./fs_temp
FS_DATA = ./fs_data

# Jupyter Lab Server
JP_IP = 127.0.0.9
JP_PORT = 8888

# Reset True, remove backendDb/history
RESET = false
`

var testing_common_map = map[string]string{
	"NAME": "common", "MAIN_WAIT_BEAT": "0.01",
	"LG_IP": "127.0.0.1", "LG_PORT": "9020",
	"CF_IP": "127.0.0.2", "CF_PORT": "1026", "CF_REFRESH": "300",
	"NT_IP": "127.0.0.3", "NT_PORT": "10080",
	"TB_IP": "127.0.0.4", "TB_PORT": "31337", "TB_TOKEN": "!!! FILL ME TOKEN !!!", "TB_CHATID": "!!! FILL ME CHATID !!!",
	"SC_IP": "127.0.0.5", "SC_PORT": "5001",
	"MT_IP": "127.0.0.6", "MT_PORT": "5000",
	"AE_IP": "127.0.0.7", "AE_INTPORT": "9002", "AE_EXTPORT": "9003",
	"DB_IP": "127.0.0.8", "DB_PORT": "5120", "DB_ENDPOINT": "http://db:5220", "DB_SSLCERT": "false", "DB_BACKEND": "./backend.db",
	"FS_TEMP": "./fs_temp", "FS_DATA": "./fs_data",
	"JP_IP": "127.0.0.9", "JP_PORT": "8888",
	"RESET": "false",
}

func Test_1_check_template(t *testing.T) {
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking templates discrepancies for COMMON section... (test_1_check_template)\n\n")
	parser, _ := ini.Load(bytes.NewBufferString(configTemplate))
	parserTest, _ := ini.Load(bytes.NewBufferString(configTemplateTest))
	var section1 []string
	for _, key := range parserTest.Section(COMMON_SECTION).Keys() {
		section1 = append(section1, key.Name())
	}
	var section2 []string
	for _, key := range parser.Section(COMMON_SECTION).Keys() {
		section2 = append(section2, key.Name())
	}
	var more_option_test []string
	for _, key1 := range section1 {
		if !PyHelpers.InStringSlice(key1, section2) {
			more_option_test = append(more_option_test, key1)
		}
	}
	var more_option_standard []string
	for _, key2 := range section2 {
		if !PyHelpers.InStringSlice(key2, section1) {
			more_option_standard = append(more_option_standard, key2)
		}
	}
	// compare sections
	if len(more_option_test) != 0 {
		t.Errorf("Test template_config doesn't match to standard template_config, more params in test template_config : %d", len(more_option_test))
	}
	if len(more_option_standard) != 0 {
		t.Errorf("Test template_config doesn't match to standard template_config, more params in standard template_config : %d", len(more_option_standard))
	}
}

func Test_2_loading_from_env_var_without_env_var(t *testing.T) {
	// should failed
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking the loading of config with only env vars with missing env vars (test_2_loading_from_env_var_without_env_var)\n\n")
	// Mock ExitFunc
	exitCode := 0
	mockExit := func(code int) {
		exitCode = code
		panic("mock exit") // Simulate program termination
	}
	// Replace ExitFunc with mockExit
	ExitFunc = mockExit
	// Reset ExitFunc after the test
	defer func() { ExitFunc = nil }()

	defer func() {
		if r := recover(); r != "mock exit" {
			t.Fatalf("Expected mock exit, got: %v", r)
		}
		if exitCode != 1 {
			t.Errorf("Expected exit code 1, got: %d", exitCode)
		}
	}()
	_ = NewConfig("test from template", "", "", true)
}

var default_file_cfg string = "config.cfg"

func checkConf(t *testing.T, config *Config) {
	selfConfItems := reflect.ValueOf(config).Elem()
	selfConfTypes := reflect.TypeOf(config).Elem()
	// loop over config object properties and update it with new value (=> "COMMON" section)
	for i := 0; i < selfConfItems.NumField(); i++ {
		if !PyHelpers.InStringSlice(selfConfTypes.Field(i).Name, notUpdatableConfigProperties) {
			if selfConfItems.Field(i).String() != testing_common_map[selfConfTypes.Field(i).Name] {
				t.Errorf("%s : value '%s' doesn't match to '%s'", selfConfTypes.Field(i).Name, selfConfItems.Field(i).String(), testing_common_map[selfConfTypes.Field(i).Name])
			}
		}
	}
}

func Test_3_loading_with_config_file_not_found(t *testing.T) {
	// should failed
	configTemplate = configTemplateTest
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking template config generation and config file reloading from it...  (test_3_loading_with_config_file_not_found)\n\n")
	// Mock ExitFunc
	exitCode := 0
	mockExit := func(code int) {
		exitCode = code
		panic("mock exit") // Simulate program termination
	}
	// Replace ExitFunc with mockExit
	ExitFunc = mockExit
	// Reset ExitFunc after the test
	defer func() { ExitFunc = nil }()

	defer func() {
		if r := recover(); r != "mock exit" {
			t.Fatalf("Expected mock exit, got: %v", r)
		}
		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got: %d", exitCode)
		}
	}()
	curDir, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(curDir, default_file_cfg)); err == nil {
		_ = os.Remove(filepath.Join(curDir, default_file_cfg))
	}
	// should failed
	_ = NewConfig("test from template", fmt.Sprintf("./%s", default_file_cfg), "", true)
}

func TestReloadFromGeneratedTemplate(t *testing.T) {
	// retry with the generated config file
	load_from_generate_template_conf := NewConfig("test from template", fmt.Sprintf("./%s", default_file_cfg), "", true)
	checkConf(t, load_from_generate_template_conf)
}

func LoadMapIntoINI(sectionName string, data map[string]string, parser *ini.File) *ini.File {
	for key, value := range data {
		section, err := parser.GetSection(sectionName)
		if err != nil {
			section, _ = parser.NewSection(sectionName)
		}
		section.NewKey(key, value)
	}
	return parser
}

func Test_4_conf_and_force_conf(t *testing.T) {
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking forced config generation and merge with standard config...  (test_4_conf_and_force_conf)\n\n")
	curDir, _ := os.Getwd()
	force_config_file_path := "force_config.cfg"
	// remove force config if exists
	if _, err := os.Stat(filepath.Join(curDir, force_config_file_path)); err == nil {
		_ = os.Remove(filepath.Join(curDir, force_config_file_path))
	}
	// create forced config file
	force_map := map[string]string{}
	parser := ini.Empty()
	for key := range testing_common_map {
		force_map[key] = key
	}
	parser = LoadMapIntoINI(COMMON_SECTION, force_map, parser)
	parser.SaveTo(force_config_file_path)

	generate_force_conf := NewConfig("test from forced config", default_file_cfg, force_config_file_path, true)
	// check if all value have been overloaded
	selfConfItems := reflect.ValueOf(generate_force_conf).Elem()
	selfConfTypes := reflect.TypeOf(generate_force_conf).Elem()
	// loop over config object properties and update it with new value (=> "COMMON" section)
	for i := 0; i < selfConfItems.NumField(); i++ {
		if !PyHelpers.InStringSlice(selfConfTypes.Field(i).Name, notUpdatableConfigProperties) {
			if selfConfItems.Field(i).String() != selfConfTypes.Field(i).Name {
				t.Errorf("%s doesn't match to %s force config has not been overloaded...", selfConfTypes.Field(i).Name, selfConfItems.Field(i).String())
			}
		}
	}
	// remove force config if exists
	if _, err := os.Stat(filepath.Join(curDir, force_config_file_path)); err == nil {
		_ = os.Remove(filepath.Join(curDir, force_config_file_path))
	}
}

func Test_5_conf_and_force_conf_plus_other_sections(t *testing.T) {
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking forced config generation and merge with standard config plus other section...  (test_5_conf_and_force_conf_plus_other_sections)\n\n")
	curDir, _ := os.Getwd()
	force_config_file_path := "force_config.cfg"
	// remove force config if exists
	if _, err := os.Stat(filepath.Join(curDir, force_config_file_path)); err == nil {
		_ = os.Remove(filepath.Join(curDir, force_config_file_path))
	}
	// create forced config file
	force_map := map[string]string{}
	parser := ini.Empty()
	section, _ := parser.NewSection("new section")
	section.NewKey("option", "option_val")
	// generate forced and set names in value, e.g: NAME=NAME, MAIN_WAIT_BEAT=MAIN_WAIT_BEAT...
	for key := range testing_common_map {
		force_map[key] = key
	}

	parser = LoadMapIntoINI(COMMON_SECTION, force_map, parser)
	parser.SaveTo(force_config_file_path)

	generate_force_conf := NewConfig("test from forced config plus other sections", default_file_cfg, force_config_file_path, true)
	// check if all value have been overloaded
	selfConfItems := reflect.ValueOf(generate_force_conf).Elem()
	selfConfTypes := reflect.TypeOf(generate_force_conf).Elem()
	// loop over config object properties and update it with new value (=> "COMMON" section)
	for i := 0; i < selfConfItems.NumField(); i++ {
		if !PyHelpers.InStringSlice(selfConfTypes.Field(i).Name, notUpdatableConfigProperties) {
			if selfConfItems.Field(i).String() != selfConfTypes.Field(i).Name {
				t.Errorf("%s doesn't match to %s force config has not been overloaded...", selfConfTypes.Field(i).Name, selfConfItems.Field(i).String())
			}
		}
	}
	if generate_force_conf.Parser.Section("new section").Key("option").Value() != "option_val" {
		t.Errorf("section 'new section' hasn't been loaded correctly by parser...")
	}

	// remove force config if exists
	if _, err := os.Stat(filepath.Join(curDir, force_config_file_path)); err == nil {
		_ = os.Remove(filepath.Join(curDir, force_config_file_path))
	}
	// delete template config.cfg if exists
	if _, err := os.Stat(filepath.Join(curDir, default_file_cfg)); err == nil {
		_ = os.Remove(filepath.Join(curDir, default_file_cfg))
	}
}

func addEnv(key, value string) {
	_ = os.Setenv(key, value)
}

func rmEnv(key string) {
	_ = os.Unsetenv(key)
}

func Test_6_only_env_var(t *testing.T) {
	for key, val := range testing_common_map {
		addEnv(key, val)
	}
	defer func() {
		for key := range testing_common_map {
			rmEnv(key)
		}
	}()
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking the loading of env vars only ...  (test_6_only_env_var)\n\n")
	env_conf := NewConfig("test only env vars", "", "", true)
	checkConf(t, env_conf)
}

func Test_7_env_var_plus_conf(t *testing.T) {
	fmt.Printf("\n\n#####################################################################\n\n")
	fmt.Printf("Checking the loading of env vars plus standard config ...  (test_7_env_var_plus_conf)\n\n")
	curDir, _ := os.Getwd()
	std_config_file_path := filepath.Join(curDir, "std_config.cfg")
	// remove std config if exists
	if _, err := os.Stat(std_config_file_path); err == nil {
		_ = os.Remove(std_config_file_path)
	}
	// create random list of key to modify :
	std_map := map[string]string{}

	parser := ini.Empty()
	section, _ := parser.NewSection("new section")
	section.NewKey("option", "option_val")
	// generate forced and set names in value, e.g: NAME=NAME, MAIN_WAIT_BEAT=MAIN_WAIT_BEAT...
	for key := range testing_common_map {
		std_map[key] = key
	}

	parser = LoadMapIntoINI(COMMON_SECTION, std_map, parser)
	parser.SaveTo(std_config_file_path)

	env_plus_std_conf := NewConfig("test env vars plus std config", std_config_file_path, "", true)
	// check if all value have been overloaded

	selfConfItems := reflect.ValueOf(env_plus_std_conf).Elem()
	selfConfTypes := reflect.TypeOf(env_plus_std_conf).Elem()
	// loop over config object properties and update it with new value (=> "COMMON" section)
	for i := 0; i < selfConfItems.NumField(); i++ {
		if !PyHelpers.InStringSlice(selfConfTypes.Field(i).Name, notUpdatableConfigProperties) {
			if selfConfItems.Field(i).String() != selfConfTypes.Field(i).Name {
				t.Errorf("%s doesn't match to %s standard config has not been overloaded...", selfConfTypes.Field(i).Name, selfConfItems.Field(i).String())
			}
		}
	}

	// remove std config if exists
	if _, err := os.Stat(std_config_file_path); err == nil {
		_ = os.Remove(std_config_file_path)
	}
}

func randomChoices(data map[string]string) []string {
	slicee := []string{}
	for key := range data {
		slicee = append(slicee, key)
	}
	nbValPicked := rand.Intn(len(slicee))
	retSlice := []string{}
	for range nbValPicked {
		idx := rand.Intn(len(slicee))
		retSlice = append(retSlice, slicee[idx])
	}
	return retSlice
}

// GetFieldByName returns the value of a field by its name
func GetFieldByName(obj interface{}, fieldName string) interface{} {
	// Get the reflect.Value of the object
	v := reflect.ValueOf(obj)

	// Ensure the object is a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return nil
	}

	// Get the reflect.Type of the object
	t := v.Type()

	// Get the field by name
	field, ok := t.FieldByName(fieldName)
	if !ok {
		fmt.Printf("Field '%s' not found\n", fieldName)
		return nil
	}

	// Get the value of the field
	fieldValue := v.FieldByIndex(field.Index)
	return fieldValue.Interface()
}

func Test_8_partial_env_var_plus_partial_conf_plus_partial_force_conf(t *testing.T) {
	for key, val := range testing_common_map {
		addEnv(key, val)
	}
	defer func() {
		for key := range testing_common_map {
			rmEnv(key)
		}
	}()
	fmt.Printf("\n\n#####################################################################\n\n")
	curDir, _ := os.Getwd()
	original_testing_common_map := map[string]string{}
	for key, val := range testing_common_map {
		original_testing_common_map[key] = val
	}
	for x := range 10 {
		fmt.Printf("\nRun %d : checking the partial loading of env vars, standard config and forced config...  (test_8_partial_env_var_plus_partial_conf_plus_partial_force_conf)\n\n", x+1)
		// re-init map with original values
		original_testing_common_map := map[string]string{}
		for key, val := range testing_common_map {
			original_testing_common_map[key] = val
		}
		std_config_file := "std_config.cfg"
		force_config_file := "force_config.cfg"
		std_config_file_path := filepath.Join(curDir, std_config_file)
		force_config_file_path := filepath.Join(curDir, force_config_file)
		// remove force and std config if exists
		if _, err := os.Stat(std_config_file_path); err == nil {
			_ = os.Remove(std_config_file_path)
		}
		if _, err := os.Stat(force_config_file_path); err == nil {
			_ = os.Remove(force_config_file_path)
		}
		// create parser
		parser_std_conf_test := ini.Empty()
		parser_forced_conf_test := ini.Empty()

		// generate standard and set names in value, e.g: NAME=NAME, MAIN_WAIT_BEAT=MAIN_WAIT_BEAT...
		std_map := map[string]string{}
		force_map := map[string]string{}

		// create standard config file
		random_std := randomChoices(original_testing_common_map)
		for _, key := range random_std {
			std_map[key] = fmt.Sprintf("std_%s", key)
		}
		parser_std_conf_test = LoadMapIntoINI(COMMON_SECTION, std_map, parser_std_conf_test)
		parser_std_conf_test.SaveTo(std_config_file)

		// create forced config file
		random_forced := randomChoices(original_testing_common_map)
		for _, key := range random_forced {
			force_map[key] = fmt.Sprintf("force_%s", key)
		}
		parser_forced_conf_test = LoadMapIntoINI(COMMON_SECTION, force_map, parser_forced_conf_test)
		parser_forced_conf_test.SaveTo(force_config_file)

		env_std_forced_conf_plus := NewConfig("test only env vars plus partial conf +++", std_config_file_path, force_config_file_path, true)

		// what the config actually does
		for key, val := range force_map {
			std_map[key] = val
		}
		for key, val := range std_map {
			original_testing_common_map[key] = val
		}

		// check if all value have been overloaded
		for key, val := range original_testing_common_map {
			if !PyHelpers.InStringSlice(key, notUpdatableConfigProperties) {
				if GetFieldByName(env_std_forced_conf_plus, key) != val {
					t.Errorf("%s doesn't match to %s standard config has not been overloaded correctly...", GetFieldByName(env_std_forced_conf_plus, key), val)
				}
			}
		}

		for key, val := range original_testing_common_map {
			if env_std_forced_conf_plus.Parser.Section(COMMON_SECTION).Key(key).Value() != val {
				t.Errorf("%s doesn't match to %s standard config has not been overloaded correctly...", env_std_forced_conf_plus.Parser.Section(COMMON_SECTION).Key(key).Value(), val)
			}
		}

		// remove force and std config if exists
		if _, err := os.Stat(std_config_file_path); err == nil {
			_ = os.Remove(std_config_file_path)
		}
		if _, err := os.Stat(force_config_file_path); err == nil {
			_ = os.Remove(force_config_file_path)
		}
	}
}
