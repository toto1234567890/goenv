package Helpers

import (
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strings"

	"govenv/pkg/common/PyHelpers"
)

var MainAppList = []string{"common", "trading", "scrapyt", "analyst", "backtest", "MT5", "CCXT", "INVESTOPEDIA"}

// ####################################################################
// ###################### Default argParser ###########################
func DefaultArguments(argv []string, specificArgs []string) (bool, map[string]string, []string) {
	args := argv[1:]
	defaultArgs := []string{"--name", "--host", "--port", "--conf", "--log_level"}
	cmdArgs := map[string]string{}
	for i := 0; i < len(args)-1; i += 2 {
		if PyHelpers.InStringSlice(strings.ToLower(args[i]), defaultArgs) {
			arg := strings.Replace(args[i], "--", "", 1)
			cmdArgs[arg] = args[i+1]
		} else if len(specificArgs) > 0 {
			if PyHelpers.InStringSlice(strings.ToLower(args[i]), defaultArgs) {
				arg := strings.Replace(args[i], "--", "", 1)
				cmdArgs[arg] = args[i+1]
			} else {
				return false, cmdArgs, defaultArgs
			}
		} else {
			return false, cmdArgs, defaultArgs
		}
	}
	return true, cmdArgs, defaultArgs
}

// ###################### Default argParser ###########################
// ####################################################################
// ################## Get unused port from OS #########################
func GetUnusedPort() int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

// ################## Get unused port from OS #########################
// ####################################################################
// ################ Get value or return default  ######################
func GetOrDefault(valueMap map[string]string, valueString string, defaultValue string, key string) string {
	if key != "" {
		if v, found := valueMap[key]; found {
			return v
		}
		return defaultValue
	} else {
		if valueString != "" {
			return valueString
		}
		return defaultValue
	}
}

// ################ Get value or return default  ######################
// ####################################################################
// ################### Search for config files  #######################
func LoadConfigFiles(root string, dirFilters []string, extFilters string) map[string]string {
	if root == "" {
		root, _ = os.Getwd()
	}
	if len(dirFilters) == 0 {
		dirFilters = MainAppList
	}
	if extFilters == "" {
		extFilters = ".cfg"
	}
	config := make(map[string]string)
	currentFilePath := filepath.Join(root, "current.cfg")
	if _, err := os.Stat(currentFilePath); err == nil {
		config["current"] = currentFilePath
	}
	_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		dirFolder := filepath.Base(filepath.Dir(path))
		if !info.IsDir() && PyHelpers.InStringSlice(dirFolder, dirFilters) {
			if info.Name() == dirFolder+extFilters {
				config[dirFolder] = path
			}
		}
		return nil // (required by filepath.Walk)
	})
	return config
}

// ################### Search for config files  #######################
// ####################################################################
// ################ Search for main parent config  ####################
func DefaultConfigFile(confFilePath string) string {
	root, err := os.Getwd()
	if err != nil {
		return ""
	}
	// if not config file in root dir
	if !strings.HasPrefix(confFilePath, root) {
		return ""
	}
	dirList := strings.Split(confFilePath, string(os.PathSeparator))
	nbAppRelated := 0
	dir := ""
	for _, mainApp := range MainAppList {
		for _, dir := range dirList {
			if mainApp == dir {
				nbAppRelated += 1
				break // ok at least one root folder is named like a MainAppList
			}
		}
	}
	if nbAppRelated == 0 {
		return ""
	} // not a main app in MainAppList
	// find the most coherent config file (from leaf to root)
	parentDir := filepath.Dir(confFilePath)
	for {
		if strings.HasSuffix(filepath.Base(parentDir), dir) {
			break
		}
		parentDir = filepath.Dir(parentDir)
	}
	return fmt.Sprint(filepath.Join(parentDir, fmt.Sprintf("%s.cfg", dir)))
}

//################ Search for main parent config  ####################
//####################################################################
