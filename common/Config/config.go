package Config

// FIXME pour lancer le serveur de config il faut savoir si il est sur la mêmem machine que le process
//
// IP du serveur de config doit être mis dans les variables d'environmenet au préalable => raise error on client side (TODO)
// Le serveur de config doit ajouter l'IP public dans les variable d'environement TODO  (ou fait manuellement)
// décider si IP ou nom de machine ou ....
// cas 1 : le serveur de config n'est pas démarré et est sur la même machine// 	le serveur de config démarre et remplace les IP qui commence par 127.0.0...., avec l'IP public de la machine (TODO here)// 	COMMON_FILE_PATH ? comment le remplir et quoi envoyer comme path ? Dictionnaire quoi ? changer le launch config server pour ajouter la deuxième config TODO // 	le client attend le démarrage et load -> already implemented not tested

//	-> ajouter une fonction client qui remplace automatiquement leur IP publique avec IP local (TODO)
//	aussi vérifier q

// cas 2 : le serveur de config n'est pas démarré et pas sur la même machine
// 	le client (TODO here), raise error and exit

// cas 3 : le serveur de config n'est pas démarré et pas atteignable (sur la même machine ou pas)
// 	le client (TODO here), raise error and exit

// faire les tests d'INSERT, UPDATE, DELETE without config server TODO
// faire les tests d'INSERT, UPDATE, DELETE with config server TODO
// implementer les stocky socket pour le config Handlers... TODO

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	pb "govenv/api/proto/configMsg"
	"govenv/pkg/common/Helpers"
	"govenv/pkg/common/PyHelpers"

	"gopkg.in/ini.v1"
)

// GO specific ExitFunc used for Testing (mock os.Exit :()
var ExitFunc = os.Exit

// parser main section
var COMMON_SECTION string = "COMMON"

// # This variables should also be set in OS environ variables
// # LG_PORT = 9020  = IANA This port is listed as "Python logging default TCP port"
// # CF_PORT = 14142  = IANA This port is listed as "Windows Messenger"
// # NT_PORT = 10080 = IANA This port is listed as "Amanda Backup"
// # TB_PORT = 31337 = IANA This port is listed as "Back Orifice" :)
var autoConf = map[string]string{
	"MAIN_WAIT_BEAT": "0.01",
	"LG_IP":          "127.0.0.1",
	"LG_PORT":        "9020",
	"CF_IP":          "127.0.0.1",
	"CF_PORT":        "1026",
	"NT_IP":          "127.0.0.1",
	"NT_PORT":        "10080",
	"TB_IP":          "127.0.0.1",
	"TB_PORT":        "31337",
	"SC_IP":          "127.0.0.1",
	"SC_PORT":        "5001",
	"MT_IP":          "127.0.0.1",
	"MT_PORT":        "5000",
	"AE_IP":          "127.0.0.1",
	"AE_INTPORT":     "9002",
	"AE_EXTPORT":     "9003",
	"JP_IP":          "127.0.0.1",
	"JP_PORT":        "8888",
}

// Write config template if needed
var configTemplate string = `
[COMMON]

# Config name
NAME = common
MAIN_WAIT_BEAT = 0.01

# Log server
LG_IP = 127.0.0.1
LG_PORT = 9020

# Config server
CF_IP = 127.0.0.1
CF_PORT = 1026
CF_REFRESH = 300 

# Telegram Notifs
NT_IP = 127.0.0.1
NT_PORT = 10080

# Telebot Remote Control
TB_TOKEN = !!! FILL ME !!!
TB_CHATID = !!! FILL ME !!!
TB_IP = 127.0.0.1
TB_PORT = 31337

# Scheduler 
SC_IP = 127.0.0.1
SC_PORT = 5001

# Monitoring 
MT_IP = 127.0.0.1
MT_PORT = 5000

# Matrix Q
AE_IP = 127.0.0.1 
AE_INTPORT = 9002
AE_EXTPORT = 9003

# Database
DB_IP = 127.0.0.1 
DB_PORT = 5123
DB_ENDPOINT = http://db:5123
DB_SSLCERT = false
DB_BACKEND = Database/backend.db

# datas on file system
FS_TEMP = ./fs_temp
FS_DATA = ./fs_data

# Jupyter Lab Server
JP_IP = 127.0.0.1
JP_PORT = 8888

# Reset True, remove backendDb/history
RESET = false
`

// required config used to check if minimal config is filled
type RequiredCommonConf struct {
	NAME           string
	MAIN_WAIT_BEAT string

	LG_IP   string
	LG_PORT string

	CF_IP      string
	CF_PORT    string
	CF_REFRESH string

	NT_IP   string
	NT_PORT string

	TB_TOKEN  string
	TB_CHATID string
	TB_IP     string
	TB_PORT   string

	SC_IP   string
	SC_PORT string

	MT_IP   string
	MT_PORT string

	AE_IP      string
	AE_INTPORT string
	AE_EXTPORT string

	DB_IP       string
	DB_PORT     string
	DB_ENDPOINT string
	DB_SSLCERT  string
	DB_BACKEND  string

	FS_TEMP string
	FS_DATA string

	JP_IP   string
	JP_PORT string

	RESET string
}

// main Config class
type Config struct {
	NAME           string
	MAIN_WAIT_BEAT string

	LG_IP   string
	LG_PORT string

	CF_IP      string
	CF_PORT    string
	CF_REFRESH string

	NT_IP   string
	NT_PORT string

	TB_TOKEN  string
	TB_CHATID string
	TB_IP     string
	TB_PORT   string

	SC_IP   string
	SC_PORT string

	MT_IP   string
	MT_PORT string

	AE_IP      string
	AE_INTPORT string
	AE_EXTPORT string

	DB_IP       string
	DB_PORT     string
	DB_ENDPOINT string
	DB_SSLCERT  string
	DB_BACKEND  string

	FS_TEMP string
	FS_DATA string

	JP_IP   string
	JP_PORT string

	RESET string

	// not updated
	COMMON_FILE_PATH string
	Name             string
	loggerLog        func(string, string)
	// updated
	MemConfig     map[string]map[string]string
	MemConfigPath string
	Parser        *ini.File
	// config server
	withConfigServer bool
	Handler          *ConfigProtoHandler
	// trigger parent function on received config_server update
	ParentOnMemConfUpdate func(map[string]map[string]string)
}

var confLoadedFrom = []string{}

func configLoadedFrom(conf string, get bool) string {
	if get {
		confsListStr := confLoadedFrom[0]
		if len(confLoadedFrom) > 1 {
			confsListStr = fmt.Sprintf("%s", confLoadedFrom[1])
			for _, confs := range confLoadedFrom[2:] {
				confsListStr += fmt.Sprintf(" <- %s", confs)
			}
		}
		return confsListStr
	} else {
		confLoadedFrom = append(confLoadedFrom, conf)
		return ""
	}
}

func NewConfig(name, configFilePath, forceConfigFilePath string, ignoreConfigServer bool) *Config {
	confLoadedFrom = []string{"Not loaded", "OS environment variables"}
	// return new config with minimal required COMMON config section
	config := &Config{}
	config.Name = "Config"
	if name != "" {
		config.Name = name
	}

	// check if config file path is provided, if not, try to load config from os env var
	if configFilePath == "" {
		// no file path, try to load from env vars
		fmt.Printf("Trying to load config from os environment variables.. . .\n")

		// init empty *ini.File
		config.Parser = ini.Empty()

		// load config, forced config used as overload of env vars or for secrets also
		config.loadConfigObj(COMMON_SECTION, forceConfigFilePath)

		// no exit ok !
		if forceConfigFilePath != "" {
			configLoadedFrom(forceConfigFilePath, false)
		}
		confsLoaded := configLoadedFrom("get config file(s) loaded", true)
		config.Parser.Section(COMMON_SECTION).Key("COMMON_FILE_PATH").SetValue(confsLoaded)
		fmt.Printf(fmt.Sprintf("Configuration has been loaded succesfully from : '%s'\n", confsLoaded))
	} else {

		// check if config file exists, try to load it, otherwise use the hard coded template file
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			// config file not exist at provided path
			fmt.Printf("Config file '%s' does not exist, skip loading\n", configFilePath)

			// check if directory path exists to save the template file (configTemplate)
			configDir := filepath.Dir(configFilePath)
			if _, err := os.Stat(filepath.Dir(configFilePath)); os.IsNotExist(err) {
				fmt.Printf("Unable to save template config parent directory '%s' doesn't exists or unreachable : '%v'\n", configDir, err)
				ExitFunc(1)
			}

			// All ok, add config file path in common section, and save it
			configTemplate += fmt.Sprintf("\n# Config file path\nCOMMON_FILE_PATH = %s\n", configFilePath)
			config.Parser, _ = ini.Load(bytes.NewBufferString(configTemplate))

			err = config.Parser.SaveTo(configFilePath)
			if err != nil {
				fmt.Printf("Unexpected error while trying to save config file template : '%v' in directory : '%s'\n", err, configDir)
				ExitFunc(1)
			} else {
				fmt.Printf("Config template file '%s', saved in directory : '%s', please check...\n", filepath.Base(configFilePath), configDir)
			}
			// no error, just need to check and fill in the template config file just created...
			ExitFunc(0)
		} else {
			// try to load the config file provided
			config.Parser, err = ini.Load(configFilePath)
			if err != nil {
				fmt.Printf("Error while trying to load config file : '%s' -> '%v'\n", configFilePath, err)
				fmt.Printf("Exit program...\n")
				ExitFunc(1)
			}
			configLoadedFrom(configFilePath, false)
		}

		// load and check config
		config.loadConfigObj(COMMON_SECTION, forceConfigFilePath)
		if forceConfigFilePath != "" {
			configLoadedFrom(forceConfigFilePath, false)
		}
		confsLoaded := configLoadedFrom("get config file(s) loaded", true)
		fmt.Printf("Configuration has been loaded succesfully from : '%s'\n", confsLoaded)
	}

	// start listener (and previously config_sever if not running)
	if !ignoreConfigServer {
		config.withConfigServer = true
		// load Handler
		config.Handler = NewConfigHandler("configProto", config.Name, config)
		// start_config_server if not already running
		config.startConfigServer(config.COMMON_FILE_PATH, 1)
		// FIXME self.COMMON_FILE_PATH should be overloaded by server conf with for example "from config server..."

		// FIXME check if parser and mem_config UPDATE IS NEEDED ?!
		//// start listener, and continiously update config on server sending
		//go config.configListener()

		//// get config updated with config parser with complete infos form server
		//config.getConfig()
		//// check if update is necessary

		//// if yes, update parser
		//// send update of self current parser object to the config server (without COMMON_SECTION)
		//preloadedConfig := make(map[string]map[string]string)
		//for _, section := range config.Parser.Sections() {
		//	if section.Name() != COMMON_SECTION {
		//		preloadedConfig[section.Name()] = section.KeysHash()
		//	}
		//}
		//config.Update(preloadedConfig, true)

		//// get complete memConfig
		//config.getMemConfig()

		//// check if update is necessary
		//// load personal mem_config if exists... if yes, send an update
		//config.LoadLocalMemConfig(false)
		//if len(config.MemConfig) > 0 {
		//	config.UpdateMemConfig(config.MemConfig)
		//}

		// ready to start !
	} else {
		config.withConfigServer = false
		curDir, _ := os.Getwd()
		config.MemConfigPath = filepath.Join(curDir, fmt.Sprintf("%s_mem_config.binu", config.Name))
		config.LoadLocalMemConfig(false)
	}

	return config
}

/////////////////////////////////
// loading step (ordered -> env vars <- config file <- force_config)

func (config *Config) loadOsEnvVars(objConf *RequiredCommonConf) {
	// load OS env variables
	NAME, exist := os.LookupEnv("NAME")
	if exist {
		objConf.NAME = NAME
	}
	MAIN_WAIT_BEAT, exist := os.LookupEnv("MAIN_WAIT_BEAT")
	if exist {
		objConf.MAIN_WAIT_BEAT = MAIN_WAIT_BEAT
	}
	LG_IP, exist := os.LookupEnv("LG_IP")
	if exist {
		objConf.LG_IP = LG_IP
	}
	LG_PORT, exist := os.LookupEnv("LG_PORT")
	if exist {
		objConf.LG_PORT = LG_PORT
	}
	CF_IP, exist := os.LookupEnv("CF_IP")
	if exist {
		objConf.CF_IP = CF_IP
	}
	CF_PORT, exist := os.LookupEnv("CF_PORT")
	if exist {
		objConf.CF_PORT = CF_PORT
	}
	CF_REFRESH, exist := os.LookupEnv("CF_REFRESH")
	if exist {
		objConf.CF_REFRESH = CF_REFRESH
	}
	NT_IP, exist := os.LookupEnv("NT_IP")
	if exist {
		objConf.NT_IP = NT_IP
	}
	NT_PORT, exist := os.LookupEnv("NT_PORT")
	if exist {
		objConf.NT_PORT = NT_PORT
	}
	TB_IP, exist := os.LookupEnv("TB_IP")
	if exist {
		objConf.TB_IP = TB_IP
	}
	TB_PORT, exist := os.LookupEnv("TB_PORT")
	if exist {
		objConf.TB_PORT = TB_PORT
	}
	TB_TOKEN, exist := os.LookupEnv("TB_TOKEN")
	if exist {
		objConf.TB_TOKEN = TB_TOKEN
	}
	TB_CHATID, exist := os.LookupEnv("TB_CHATID")
	if exist {
		objConf.TB_CHATID = TB_CHATID
	}
	SC_IP, exist := os.LookupEnv("SC_IP")
	if exist {
		objConf.SC_IP = SC_IP
	}
	SC_PORT, exist := os.LookupEnv("SC_PORT")
	if exist {
		objConf.SC_PORT = SC_PORT
	}
	MT_IP, exist := os.LookupEnv("MT_IP")
	if exist {
		objConf.MT_IP = MT_IP
	}
	MT_PORT, exist := os.LookupEnv("MT_PORT")
	if exist {
		objConf.MT_PORT = MT_PORT
	}
	AE_IP, exist := os.LookupEnv("AE_IP")
	if exist {
		objConf.AE_IP = AE_IP
	}
	AE_INTPORT, exist := os.LookupEnv("AE_INTPORT")
	if exist {
		objConf.AE_INTPORT = AE_INTPORT
	}
	AE_EXTPORT, exist := os.LookupEnv("AE_EXTPORT")
	if exist {
		objConf.AE_EXTPORT = AE_EXTPORT
	}
	DB_IP, exist := os.LookupEnv("DB_IP")
	if exist {
		objConf.DB_IP = DB_IP
	}
	DB_PORT, exist := os.LookupEnv("DB_PORT")
	if exist {
		objConf.DB_PORT = DB_PORT
	}
	DB_ENDPOINT, exist := os.LookupEnv("DB_ENDPOINT")
	if exist {
		objConf.DB_ENDPOINT = DB_ENDPOINT
	}
	DB_SSLCERT, exist := os.LookupEnv("DB_SSLCERT")
	if exist {
		objConf.DB_SSLCERT = DB_SSLCERT
	}
	DB_BACKEND, exist := os.LookupEnv("DB_BACKEND")
	if exist {
		objConf.DB_BACKEND = DB_BACKEND
	}
	FS_DATA, exist := os.LookupEnv("FS_DATA")
	if exist {
		objConf.FS_DATA = FS_DATA
	}
	FS_TEMP, exist := os.LookupEnv("FS_TEMP")
	if exist {
		objConf.FS_TEMP = FS_TEMP
	}
	JP_IP, exist := os.LookupEnv("JP_IP")
	if exist {
		objConf.JP_IP = JP_IP
	}
	JP_PORT, exist := os.LookupEnv("JP_PORT")
	if exist {
		objConf.JP_PORT = JP_PORT
	}
	RESET, exist := os.LookupEnv("RESET")
	if exist {
		objConf.RESET = RESET
	}
}

func (config *Config) loadForcedConfig(forceConfigFilePath string) *ini.File {
	inMemoryParser, err := ini.Load(forceConfigFilePath)
	if err != nil {
		fmt.Printf("\nUnable to load provided 'forced config' into *ini.File object : '%v'\n", err)
		fmt.Printf("Exit program...\n")
		ExitFunc(1)
	}
	return inMemoryParser
}

func (config *Config) setAutoConfig(conf *RequiredCommonConf) {
	confValues := reflect.ValueOf(conf).Elem()
	for key, value := range autoConf {
		if val := confValues.FieldByName(key).String(); val == "" {
			confValues.FieldByName(key).SetString(value)
		}
	}
}

func (config *Config) loadConfigObj(section, forceConfigFilePath string) {
	// load OS env vars
	objConf := RequiredCommonConf{}
	config.loadOsEnvVars(&objConf)

	// overload objConf if value in config file exists
	if config.Parser.HasSection(section) {
		if val := config.Parser.Section(section).Key("NAME").String(); val != "" {
			objConf.NAME = val
		}
		if val := config.Parser.Section(section).Key("MAIN_WAIT_BEAT").String(); val != "" {
			objConf.MAIN_WAIT_BEAT = val
		}
		if val := config.Parser.Section(section).Key("LG_IP").String(); val != "" {
			objConf.LG_IP = val
		}
		if val := config.Parser.Section(section).Key("LG_PORT").String(); val != "" {
			objConf.LG_PORT = val
		}
		if val := config.Parser.Section(section).Key("CF_IP").String(); val != "" {
			objConf.CF_IP = val
		}
		if val := config.Parser.Section(section).Key("CF_PORT").String(); val != "" {
			objConf.CF_PORT = val
		}
		if val := config.Parser.Section(section).Key("CF_REFRESH").String(); val != "" {
			objConf.CF_REFRESH = val
		}
		if val := config.Parser.Section(section).Key("NT_IP").String(); val != "" {
			objConf.NT_IP = val
		}
		if val := config.Parser.Section(section).Key("NT_PORT").String(); val != "" {
			objConf.NT_PORT = val
		}
		if val := config.Parser.Section(section).Key("TB_IP").String(); val != "" {
			objConf.TB_IP = val
		}
		if val := config.Parser.Section(section).Key("TB_PORT").String(); val != "" {
			objConf.TB_PORT = val
		}
		if val := config.Parser.Section(section).Key("TB_TOKEN").String(); val != "" {
			objConf.TB_TOKEN = val
		}
		if val := config.Parser.Section(section).Key("TB_CHATID").String(); val != "" {
			objConf.TB_CHATID = val
		}
		if val := config.Parser.Section(section).Key("SC_IP").String(); val != "" {
			objConf.SC_IP = val
		}
		if val := config.Parser.Section(section).Key("SC_PORT").String(); val != "" {
			objConf.SC_PORT = val
		}
		if val := config.Parser.Section(section).Key("MT_IP").String(); val != "" {
			objConf.MT_IP = val
		}
		if val := config.Parser.Section(section).Key("MT_PORT").String(); val != "" {
			objConf.MT_PORT = val
		}
		if val := config.Parser.Section(section).Key("AE_IP").String(); val != "" {
			objConf.AE_IP = val
		}
		if val := config.Parser.Section(section).Key("AE_INTPORT").String(); val != "" {
			objConf.AE_INTPORT = val
		}
		if val := config.Parser.Section(section).Key("AE_EXTPORT").String(); val != "" {
			objConf.AE_EXTPORT = val
		}
		if val := config.Parser.Section(section).Key("DB_IP").String(); val != "" {
			objConf.DB_IP = val
		}
		if val := config.Parser.Section(section).Key("DB_PORT").String(); val != "" {
			objConf.DB_PORT = val
		}
		if val := config.Parser.Section(section).Key("DB_ENDPOINT").String(); val != "" {
			objConf.DB_ENDPOINT = val
		}
		if val := config.Parser.Section(section).Key("DB_SSLCERT").String(); val != "" {
			objConf.DB_SSLCERT = val
		}
		if val := config.Parser.Section(section).Key("DB_BACKEND").String(); val != "" {
			objConf.DB_BACKEND = val
		}
		if val := config.Parser.Section(section).Key("FS_DATA").String(); val != "" {
			objConf.FS_DATA = val
		}
		if val := config.Parser.Section(section).Key("FS_TEMP").String(); val != "" {
			objConf.FS_TEMP = val
		}
		if val := config.Parser.Section(section).Key("JP_IP").String(); val != "" {
			objConf.JP_IP = val
		}
		if val := config.Parser.Section(section).Key("JP_PORT").String(); val != "" {
			objConf.JP_PORT = val
		}
		if val := config.Parser.Section(section).Key("RESET").String(); val != "" {
			objConf.RESET = val
		}
	}

	// overload config with forcedConfig converted as memory *ini.File
	if forceConfigFilePath != "" {
		ConfigFromForced := config.loadForcedConfig(forceConfigFilePath)
		if val := ConfigFromForced.Section(section).Key("NAME").String(); val != "" {
			objConf.NAME = val
		}
		if val := ConfigFromForced.Section(section).Key("MAIN_WAIT_BEAT").String(); val != "" {
			objConf.MAIN_WAIT_BEAT = val
		}
		if val := ConfigFromForced.Section(section).Key("LG_IP").String(); val != "" {
			objConf.LG_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("LG_PORT").String(); val != "" {
			objConf.LG_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("CF_IP").String(); val != "" {
			objConf.CF_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("CF_PORT").String(); val != "" {
			objConf.CF_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("CF_REFRESH").String(); val != "" {
			objConf.CF_REFRESH = val
		}
		if val := ConfigFromForced.Section(section).Key("NT_IP").String(); val != "" {
			objConf.NT_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("NT_PORT").String(); val != "" {
			objConf.NT_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("TB_IP").String(); val != "" {
			objConf.TB_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("TB_PORT").String(); val != "" {
			objConf.TB_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("TB_TOKEN").String(); val != "" {
			objConf.TB_TOKEN = val
		}
		if val := ConfigFromForced.Section(section).Key("TB_CHATID").String(); val != "" {
			objConf.TB_CHATID = val
		}
		if val := ConfigFromForced.Section(section).Key("SC_IP").String(); val != "" {
			objConf.SC_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("SC_PORT").String(); val != "" {
			objConf.SC_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("MT_IP").String(); val != "" {
			objConf.MT_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("MT_PORT").String(); val != "" {
			objConf.MT_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("AE_IP").String(); val != "" {
			objConf.AE_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("AE_INTPORT").String(); val != "" {
			objConf.AE_INTPORT = val
		}
		if val := ConfigFromForced.Section(section).Key("AE_EXTPORT").String(); val != "" {
			objConf.AE_EXTPORT = val
		}
		if val := ConfigFromForced.Section(section).Key("DB_IP").String(); val != "" {
			objConf.DB_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("DB_PORT").String(); val != "" {
			objConf.DB_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("DB_ENDPOINT").String(); val != "" {
			objConf.DB_ENDPOINT = val
		}
		if val := ConfigFromForced.Section(section).Key("DB_SSLCERT").String(); val != "" {
			objConf.DB_SSLCERT = val
		}
		if val := ConfigFromForced.Section(section).Key("DB_BACKEND").String(); val != "" {
			objConf.DB_BACKEND = val
		}
		if val := ConfigFromForced.Section(section).Key("FS_DATA").String(); val != "" {
			objConf.FS_DATA = val
		}
		if val := ConfigFromForced.Section(section).Key("FS_TEMP").String(); val != "" {
			objConf.FS_TEMP = val
		}
		if val := ConfigFromForced.Section(section).Key("JP_IP").String(); val != "" {
			objConf.JP_IP = val
		}
		if val := ConfigFromForced.Section(section).Key("JP_PORT").String(); val != "" {
			objConf.JP_PORT = val
		}
		if val := ConfigFromForced.Section(section).Key("RESET").String(); val != "" {
			objConf.RESET = val
		}
		// finally merge force_config parsers, with main parser, if config forced has other sections than COMMON_SECTION
		for _, section := range ConfigFromForced.Sections() {
			if section.Name() == COMMON_SECTION {
				continue
			} else {
				if !config.Parser.HasSection(section.Name()) {
					config.Parser.NewSection(section.Name())
				}
				for _, key := range section.Keys() {
					config.Parser.Section(section.Name()).Key(key.Name()).SetValue(key.Value())
				}
			}
		}
	}

	// fill in some of the main required vars, that can be set automatically if not exists
	config.setAutoConfig(&objConf)

	// check if all fields have been filled, if yes load it into configparser
	config.checkConfig(&objConf)

	// add properties to current config object properties and parser
	MapCommon := config.Parser.Section(COMMON_SECTION)
	MapCommon.ReflectFrom(&objConf)
	PyHelpers.UpdateStruct(config, objConf, true)
}

func (config *Config) checkConfig(objConf *RequiredCommonConf) {
	var missingValue []string
	objConfItems := reflect.ValueOf(objConf).Elem()
	objConfType := reflect.TypeOf(objConf).Elem()
	// loop over objConf and search for unfilled key
	for i := 0; i < objConfItems.NumField(); i++ {
		if objConfItems.Field(i).Interface() == "" {
			missingValue = append(missingValue, objConfType.Field(i).Name)
		}
	}
	// exit if one key is missing
	if len(missingValue) > 0 {
		fmt.Printf("Missing required values for field(s) : '%s'\n", strings.Join(missingValue, ", "))
		ExitFunc(1)
	}
}

// loading step (ordered -> env vars <- config file <- force_config)
/////////////////////////////////
// set parent update function

func (config *Config) OnMemConfUpdate(onMemConfUpdateFn func(map[string]map[string]string)) {
	config.ParentOnMemConfUpdate = onMemConfUpdateFn
}

// set parent update function
/////////////////////////////////
// main config methods

// only with config server while starting
func (config *Config) getConfig() {
	name := fmt.Sprintf("%s config.get_server_config", config.Name)
	configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
	configSock.SayHelloId()
	configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_get_config_object))
	// sync wait to get config
	config.Handler.ClientDeSerialize(configSock.ReceiveData())
}

// only with config server while starting
func (config *Config) getMemConfig() {
	name := fmt.Sprintf("%s config.get_mem_config", config.Name)
	configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
	configSock.SayHelloId()
	configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_get_mem_config))
	// sync wait to get mem config
	config.Handler.ClientDeSerialize(configSock.ReceiveData())
}

// with or without config server, anytime
func (config *Config) Update(SectionsKeysValues map[string]map[string]string, serverOnly bool) {
	if !serverOnly {
		// SectionsKeysValues param only for self update (with or without config server)
		for section, keyVal := range SectionsKeysValues {
			if !config.Parser.HasSection(section) {
				config.Parser.NewSection(section)
			}
			for key, val := range keyVal {
				if section == COMMON_SECTION {
					// no empty values or new key in "common section or config struct"
					PyHelpers.UpdateStructFromMapStringString(config, keyVal, true)
				}
				config.Parser.Section(section).Key(key).SetValue(val)
			}
		}
	}
	if config.withConfigServer {
		// to prevent config synchronisation issues, complete config objet is brodcasted to clients
		name := fmt.Sprintf("%s config.update", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_update_config_object))
		// wait for update done
		configSock.ReceiveData()
	}
}

// with or without config server, anytime
func (config *Config) UpdateMemConfig(SectionsKeysValues map[string]map[string]string) {
	// SectionsKeysValues param only used for self update (with or without config server)
	for section, keyVal := range SectionsKeysValues {
		if _, ok := config.MemConfig[section]; !ok {
			config.MemConfig[section] = nil
		}
		for key, val := range keyVal {
			config.MemConfig[section][key] = val
		}
	}
	if config.withConfigServer {
		// to prevent config synchronisation issues, complete mem config is brodcasted to clients
		name := fmt.Sprintf("%s config.update_mem_config", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_update_mem_config))
		configSock.ReceiveData()
	}
}

// with or without config server, anytime
func (config *Config) DeleteMemConfig(sectionKeyList map[string][]string) {
	for section, keyList := range sectionKeyList {
		if keyVal, ok := config.MemConfig[section]; ok {
			for _, key := range keyList {
				delete(keyVal, key)
			}
			if len(config.MemConfig[section]) == 0 {
				// remove section entry if no child remain
				delete(config.MemConfig, section)
			}
		}
	}
	if config.withConfigServer {
		// to prevent config synchronisation issues, complete mem config is brodcasted to clients
		name := fmt.Sprintf("%s config.update_mem_config", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_update_mem_config))
		configSock.ReceiveData()
	}
}

// with or without config server, anytime
func (config *Config) DeleteParserConfig(sectionKeyList map[string][]string) {
	for section, keyList := range sectionKeyList {
		iniSection, notFound := config.Parser.GetSection(section)
		if notFound == nil {
			// COMMON section can only be updated... no deletion
			if section != COMMON_SECTION {
				for _, key := range keyList {
					if iniSection.HasKey(key) {
						iniSection.DeleteKey(key)
					}
				}
				if len(iniSection.Keys()) == 0 {
					// remove section entry if no child remain
					config.Parser.DeleteSection(section)
				}
			}
		}
	}
	if config.withConfigServer {
		// to prevent config synchronisation issues, complete config objet is brodcasted to clients
		name := fmt.Sprintf("%s config.update", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_update_config_object))
		// wait for update done
		configSock.ReceiveData()
	}
}

// with or without config server, anytime
func (config *Config) DumpMemConfig(localDump bool, serverDump bool) {
	if !config.withConfigServer && serverDump {
		errMsg := "invalid 'config.dump_mem_config' call, the current config object is not linked to the 'config_server', check if the 'config_server' is started..."
		if config.loggerLog != nil {
			config.loggerLog(fmt.Sprintf("%s : %s", config.Name, errMsg), "ERROR")
		} else {
			fmt.Printf("%s\n", errMsg)
		}
	} else if config.withConfigServer && serverDump {
		// for *client with config server* ask for MemConfig dump...
		name := fmt.Sprintf("%s config.dump_mem_config", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_dump_mem_config))
		// wait for update done
		configSock.ReceiveData()
	}
	if localDump {
		// create or replace file and write 'binu' file, for *config server* or *client without config server*
		PyHelpers.CreateFile(true, config.MemConfigPath, config.Handler.ProtoFileDumps(config.MemConfig), false)
	}
}

// with or without config server, client are serveur loading
func (config *Config) LoadLocalMemConfig(loadRequired bool) {
	// create or replace file and write 'binu' file, for *config server* or *client*
	binContent, err := os.ReadFile(config.MemConfigPath)
	if err != nil {
		if loadRequired {
			errMsg := fmt.Sprintf("error while trying to load local mem_config file '%s' : '%v', loads aborted...", config.MemConfigPath, err)
			if config.loggerLog != nil {
				config.loggerLog(fmt.Sprintf("%s : %s", config.Name, errMsg), "ERROR")
			} else {
				fmt.Printf("%s\n", errMsg)
			}
			ExitFunc(1)
		}
		return
	}
	loadedMemConfig := config.Handler.ProtoFileLoads(binContent)
	if loadedMemConfig != nil && len(loadedMemConfig) > 0 {
		config.UpdateMemConfig(loadedMemConfig)
	}
}

//
// !!!! have to be used before or after UpdateFunc if it's set

// with config server only, update self.properties used by 'config_handler'
// used if UpdateFunc is set or unset
func (config *Config) UpdateSelf(KeysValues map[string]string) {
	// while receiving a parser "propagate_config" and deserializing 'ConfigMsg' message
	// update config properties from COMMON map
	// no empty values or new key in "common section or config class"
	PyHelpers.UpdateStructFromMapStringString(config, KeysValues, true)
}

// !!!! have to be used before or after UpdateFunc if it's set

// only with config server and UpdateFunc set
func (config *Config) UpdateSelfMemConfig(SectionsKeysValues map[string]map[string]string) {
	// used with UpdateFunc from parent
	// while receiving a "propagate_mem_config" and deserializing 'ConfigMsg' message
	config.MemConfig = SectionsKeysValues
}

// !!!! have to be used before or after UpdateFunc if it's set
//

// main config methods
/////////////////////////////////
// for local notifier, or notif server, forward message for notification and NotifLogLevel

func (config *Config) GetNotifConfig() {
	if config.withConfigServer {
		name := fmt.Sprintf("%s config.get_notif_loglevel", config.Name)
		configSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
		configSock.SayHelloId()
		configSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_get_config_object))
		// sync wait for notif loglevel conf
		config.Handler.ClientDeSerialize(configSock.ReceiveData())
	}
}

// for local notifier, or notif server, with forward message for notification and NotifLogLevel
/////////////////////////////////
// config auto update received from server

// only with config server, anytime
func (config *Config) configListener() {
	name := fmt.Sprintf("%s config.config_listener", config.Name)
	listenerSock := Helpers.MySocket(name, config.CF_IP, config.CF_PORT, 1)
	listenerSock.SayHelloId()
	listenerSock.SendData(config.Handler.ClientSerialize(pb.ConfigMsg_add_config_listener))
	for {
		// updateFunc triggered via ConfigProtoHandler.parentClassConfig.ParentOnMemConfUpdate
		config.Handler.ClientDeSerialize(listenerSock.ReceiveData())
	}
}

// config auto update received from server
/////////////////////////////////
// start config server FIXME -> if located on the same machine

func (config *Config) startConfigServer(pathConf string, wait int64) {
	if !Helpers.IsServiceListen(config.CF_IP, config.CF_PORT, 1, false) {
		err := Helpers.LaunchConfigServer(config.CF_PORT, pathConf, "", "", "", "")
		if err != nil {
			fmt.Printf("Error while trying to start 'config_server' : '%v'", err)
			panic(fmt.Sprintf("Error while trying to start 'config_server' : '%v'", err))
		} else {
			fmt.Printf("Main Config server is starting.. .  .")
		}
		for {
			if Helpers.IsServiceListen(config.CF_IP, config.CF_PORT, 1, false) {
				break
			}
			fmt.Printf("Waiting %d second(s) while config_server is starting.. .  .", wait)
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
	}
}

// start config server FIXME -> if located on the same machine
// ///////////////////////////////
// set loggerLog callback after loggerLog instanciation
func (config *Config) SetLoggerLog(LoggerLogCallBack func(string, string)) {
	config.loggerLog = LoggerLogCallBack
}

// set loggerLog callback after loggerLog instanciation
// ///////////////////////////////
