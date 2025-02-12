package PyHelpers

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unicode"
)

// ####################################################################
// ################# pass function  (= python pass) ###################
func Pass() { /* No-op, similar to Python's `pass*/ }

// ################# pass function  (= python pass) ###################
// ####################################################################
// ############# check if val in slice (= python list) ################
func InStringSlice(val string, slice []string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func InIntSlice(val int, slice []int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// ############# check if val in slice (= python list) ################
// ####################################################################
// ####### check os error, in go os generally return err also #########
func raiseOsError(osExit bool) {
	if osExit {
		os.Exit(1)
	}
}

// ####### check os error, in go os generally return err also #########
// ####################################################################
// ####################### update object ##############################
func _isZeroValue(v reflect.Value) bool {
	zero := reflect.Zero(v.Type()).Interface()
	return reflect.DeepEqual(v.Interface(), zero)
}

func UpdateStruct(target interface{}, update interface{}, noEmptyValues bool) {
	targetValue := reflect.ValueOf(target).Elem()
	updateValue := reflect.ValueOf(update)

	for i := 0; i < updateValue.NumField(); i++ {
		updateField := updateValue.Field(i)
		updateFieldName := reflect.TypeOf(update).Field(i).Name

		targetField := targetValue.FieldByName(updateFieldName)
		if targetField.IsValid() && targetField.CanSet() {
			if noEmptyValues {
				if !_isZeroValue(updateField) {
					targetField.Set(updateField)
				}
			} else {
				targetField.Set(updateField)
			}
		}

	}
}

func UpdateStructFromMapStringString(target interface{}, update map[string]string, noEmptyValues bool) {
	targetValue := reflect.ValueOf(target).Elem()
	for key, val := range update {
		targetField := targetValue.FieldByName(key)
		if targetField.IsValid() && targetField.CanSet() {
			if noEmptyValues {
				if val != "" {
					targetField.SetString(val)
				}
			} else {
				targetField.SetString(val)
			}
		}
	}
}

// ####################### update object ##############################
// ####################################################################
// ################# Trm func with or without space ###################
func Trim(input string, spaceRemove bool) string {
	var ret []rune
	if spaceRemove {
		for _, r := range input {
			if unicode.IsGraphic(r) && !unicode.IsSpace(r) {
				ret = append(ret, r)
			}
		}
	} else {
		for _, r := range input {
			if unicode.IsGraphic(r) {
				ret = append(ret, r)
			}
		}
	}
	return string(ret)
}

// ################# Trm func with or without space ###################
// ####################################################################
// ############## Capitalize function like python does ################
func Capitalize(aString string) string {
	return string(unicode.ToUpper(rune(aString[0]))) + aString[1:]
}

// ############## Capitalize function like python does ################
// ####################################################################
// ############## Substring no [:] python equivalent... ###############
func SubstringRunes(s string, startIndex int, count int) string {
	runes := []rune(s)
	length := len(runes)
	maxCount := length - startIndex
	if count > maxCount {
		count = maxCount
	}
	return string(runes[startIndex:count])
}

// ############## Substring no [:] python equivalent... ###############
// ####################################################################
// ############### Check if path exists (dir or file) #################
func FileExists(aPath string) bool {
	if _, err := os.Stat(aPath); err == nil {
		return true
	} else {
		return false
	}
}

// ############### Check if path exists (dir or file) #################
// ####################################################################
// ################# Safe and easy creation of files ##################
// createOrTruncateFile writes the given data to the specified file path.
// It creates the file if it doesn't exist or truncates it if it does.
func CreateFile(replaceFile bool, filePath string, data []byte, keepsItOpen bool) (*os.File, error) {
	// Open the file with create and write-only permissions.
	// O_CREATE: Create the file if it doesn't exist.
	// O_WRONLY: Open the file for write-only.
	// O_APPEND: Append data to the file if it exists.
	// O_TRUNC: Truncate the file if it exists.
	err := MakeDirectoryIfNotExists(filepath.Dir(filePath))
	if err != nil {
		return nil, err
	}

	var file *os.File
	if replaceFile {
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	} else {
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	}

	if err != nil {
		return nil, err
	}
	if keepsItOpen {
		return file, nil
	}
	defer file.Close()

	// Write the data to the file.
	_, err = file.Write(data)
	if err != nil {
		return nil, nil
	}

	return nil, nil
}

// ################# Safe and easy creation of files ##################
// ####################################################################
// ############### Safe and easy creation of folders ##################
func MakeDirectoryIfNotExists(dirPath string) error {
	// Use os.MkdirAll to create the directory and its parents if they don't exist.
	// os.MkdirAll does nothing if the directory already exists.
	err := os.MkdirAll(dirPath, 0o755)
	if err != nil {
		return err
	}
	return nil
}

// ############### Safe and easy creation of folders ##################
// ####################################################################
// ###################### Get splitted params #########################
func GetSplitedParam(line string) []string {
	// return one line mulitple configuration parameters into list (from string)
	return strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(line), " ", ""), " ", "|"), ";", "|"), ",", "|"), "|")
}

// ###################### Get splitted params #########################
