package Config

import (
	"errors"
	"fmt"

	pb "govenv/api/proto/configMsg"
	"govenv/pkg/common/PyHelpers"

	"google.golang.org/protobuf/proto"
)

var serializeErr error

/////////////////////////////////
// main serializer

type ConfigProtoHandler struct {
	Name                  string
	parentClassName       string
	parentClassConfig     *Config
	loggerRefreshLogLevel func(map[string][]string)
	notifRefreshSender    func(map[string]map[string]string) map[string][]string
	loggerLog             func(string, string)
}

func NewConfigHandler(name string, parentClassName string, parentClassConfig *Config) *ConfigProtoHandler {
	if name == "" {
		name = "ConfigProtoHandler"
	}
	return &ConfigProtoHandler{Name: name, parentClassName: parentClassName, parentClassConfig: parentClassConfig}
}

// set callback function for parent logger, to change notif error log level in logger
func (ConfigProtoHandler *ConfigProtoHandler) SetLoggerCallBack(loggerRefreshLogLevel func(map[string][]string)) {
	ConfigProtoHandler.loggerRefreshLogLevel = loggerRefreshLogLevel
}

// set callback function for notif client, to change notifs senders in notifie
func (ConfigProtoHandler *ConfigProtoHandler) SetNotifCallBack(notifRefreshSender func(map[string]map[string]string) map[string][]string) {
	ConfigProtoHandler.notifRefreshSender = notifRefreshSender
}

// mem config file dumps
func (ConfigProtoHandler *ConfigProtoHandler) ProtoFileDumps(memConfig map[string]map[string]string) []byte {
	SectionMap := map[string]*pb.KeysValues{}
	for sectKey, keyValMap := range memConfig {
		// Do not save empty conffig
		if len(keyValMap) > 0 {
			SectionMap[sectKey] = &pb.KeysValues{KeyValue: keyValMap}
		}
	}
	// should have no error !
	dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{SectionsKeysValues: SectionMap})
	if serializeErr != nil {
		goto serializeErr
	}
	return dataSer
serializeErr:
	// unexpected error...
	fmt.Printf("Error while trying to serialize 'proto.protoFileDumps' : unexcepted error -> %s\n", serializeErr)
	return nil
}

// mem config file loads
func (ConfigProtoHandler *ConfigProtoHandler) ProtoFileLoads(dataSer []byte) map[string]map[string]string {
	// try deserialisation
	data := &pb.ConfigMsg{}
	deserializeErr := proto.Unmarshal(dataSer, data)
	if deserializeErr != nil {
		fmt.Printf("Error while trying to deserialize 'proto.protoFileLoads' : %v\n", deserializeErr)
		return nil
	}
	memConfig := make(map[string]map[string]string)
	for sectKey, KeysValues := range data.SectionsKeysValues {
		// should not be empty, nobody should send empty message
		memConfig[sectKey] = KeysValues.GetKeyValue()
	}
	return memConfig
}

// main handler
/////////////////////////////////
// client side

func (ConfigProtoHandler *ConfigProtoHandler) ClientSerialize(ClientCmd pb.ConfigMsg_ConfigClientCmd) []byte {
	switch ClientCmd {

	// send an update of the mem_config to the config server
	case pb.ConfigMsg_update_mem_config:
		SectionMap := map[string]*pb.KeysValues{}
		for sectKey, keyValMap := range ConfigProtoHandler.parentClassConfig.MemConfig {
			// Do not send empty message
			if len(keyValMap) > 0 {
				for key, val := range keyValMap {
					if val == "" {
						delete(keyValMap, key)
					}
				}
				SectionMap[sectKey] = &pb.KeysValues{KeyValue: keyValMap}
			}
		}
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_update_mem_config, SectionsKeysValues: SectionMap})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// send an update of config parser to the config server
	case pb.ConfigMsg_update_config_object:
		SectionMap := map[string]*pb.KeysValues{}
		KeyValMap := map[string]string{}
		for _, section := range ConfigProtoHandler.parentClassConfig.Parser.Sections() {
			// Do not send empty message
			if len(section.Keys()) > 0 {
				for _, key := range section.Keys() {
					if key.Value() != "" {
						KeyValMap[key.Name()] = key.Value()
					}
				}
			}
			SectionMap[section.Name()] = &pb.KeysValues{KeyValue: KeyValMap}
			KeyValMap = map[string]string{}
		}
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_update_config_object, SectionsKeysValues: SectionMap})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// ask mem config to the config server
	case pb.ConfigMsg_get_mem_config:
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_get_mem_config})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// ask common config to the config server
	case pb.ConfigMsg_get_config_object:
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_get_config_object})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// ask for new listener to the config server
	case pb.ConfigMsg_add_config_listener:
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_add_config_listener})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// ask a dump of memory config to the config server (on service close mainly)
	case pb.ConfigMsg_dump_mem_config:
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_dump_mem_config})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// ask notifLogLevel config to the config server
	case pb.ConfigMsg_get_notif_loglevel:
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_get_notif_loglevel})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

		// GO specific as config server writed in GO !
		// only used on config server side, send an update of notifLogLevel
		// for config server, data located in volatile process env vars and notifLogLevel map variable in logger class
		// for client, data only located in notifLogLevel map variable in logger class
		// FIXME maybe only on server side ??
	case pb.ConfigMsg_update_notif_loglevel:
		SectionMap := map[string]*pb.KeysValues{}
		notifLogLevelChan := make(chan map[string]map[string]string)
		// Launch the goroutine to get notification profiles, save in process os env...
		go func() {
			// notifLogLevelChan <- ConfigProtoHandler.loggerRefreshLogLevel(nil)
		}()
		// Wait for return
		notifLogLevel := <-notifLogLevelChan
		for sectKey, keyValMap := range notifLogLevel {
			// Do not send empty message
			if len(keyValMap) > 0 {
				for key, val := range keyValMap {
					if val == "" {
						delete(keyValMap, key)
					}
				}
				SectionMap[sectKey] = &pb.KeysValues{KeyValue: keyValMap}
			}
		}
		dataSer, serializeErr := proto.Marshal(&pb.ConfigMsg{ReqClient: pb.ConfigMsg_update_mem_config, SectionsKeysValues: SectionMap})
		if serializeErr != nil {
			goto serializeErr
		}
		return dataSer

	// unknown message
	default:
		fmt.Printf("Error while trying to serialize 'proto.ConfigMsg' : unknown ClientCmd -> %s\n", ClientCmd)
		panic("Unable to serialize configMsg this shouldn't occurs !!!")
	}
serializeErr:
	// unexpected error...
	fmt.Printf("Error while trying to serialize 'proto.ConfigMsg_ClientCmd' : unexcepted error -> %s\n", serializeErr)
	return nil // or os exit ?
}

func (ConfigProtoHandler *ConfigProtoHandler) ClientDeSerialize(dataSer []byte) {
	// must not panic ! no error should occurs here !!
	// config server should send well formatted datas
	data := &pb.ConfigMsg{}
	_ = proto.Unmarshal(dataSer, data)

	switch data.RespServer {

	// receive mem config object
	case pb.ConfigMsg_propagate_mem_config:
		if ConfigProtoHandler.parentClassConfig.ParentOnMemConfUpdate == nil {
			for sectKey, KeysValues := range data.SectionsKeysValues {
				// should not be empty, nobody should send empty message
				ConfigProtoHandler.parentClassConfig.MemConfig[sectKey] = KeysValues.GetKeyValue()
			}
		} else {
			memConfig := make(map[string]map[string]string)
			for sectKey, KeysValues := range data.SectionsKeysValues {
				// should not be empty, nobody should send empty message
				memConfig[sectKey] = KeysValues.GetKeyValue()
			}
			// config update function can be then updated with self.updateMemConfig
			ConfigProtoHandler.parentClassConfig.ParentOnMemConfUpdate(memConfig)
		}

	// receive confirmation that server has updated mem config
	case pb.ConfigMsg_mem_config_update_done:
		// ok nothing to do, just waiting for response...

	// receive an update of parser config from the config server
	case pb.ConfigMsg_propagate_config:
		for sectKey, KeysValues := range data.SectionsKeysValues {
			// should not be empty, nobody should send empty message
			if sectKey == COMMON_SECTION {
				ConfigProtoHandler.parentClassConfig.UpdateSelf(KeysValues.GetKeyValue())
			}
			ConfigProtoHandler.parentClassConfig.Parser.Section(sectKey).MapTo(KeysValues.GetKeyValue())
		}

	// only used on client side while receiving an update of notif_loglevel
	// for client, data only located in notifLogLevel map variable in logger class
	case pb.ConfigMsg_propagate_notif_loglevel:
		notifLogLevel := make(map[string]map[string]string)
		for sectKey, KeysValues := range data.SectionsKeysValues {
			// should not be empty, nobody should send empty message
			notifLogLevel[sectKey] = KeysValues.GetKeyValue()
		}
		if ConfigProtoHandler.notifRefreshSender != nil {
			// if parent local notifie
			// async update local notif senders first
			notifLogLevelChan := make(chan map[string][]string)
			go func() {
				notifLogLevelChan <- ConfigProtoHandler.notifRefreshSender(notifLogLevel)
			}()
			// then update logger
			ConfigProtoHandler.loggerRefreshLogLevel(<-notifLogLevelChan)
		} else {
			// parent directly connected to notif_server (only update notif log levels)
			returnedLogLevel := map[string][]string{}
			for _, keyVal := range notifLogLevel {
				for _, logLevel := range PyHelpers.GetSplitedParam(keyVal["LOGLEVEL"]) {
					if val, ok := returnedLogLevel[logLevel]; ok {
						val = append(val, keyVal["TAG"])
					} else {
						returnedLogLevel[logLevel] = []string{keyVal["TAG"]}
					}
				}
			}
			// then update logger
			ConfigProtoHandler.loggerRefreshLogLevel(returnedLogLevel)
		}

	// recieved init mem config object
	case pb.ConfigMsg_send_mem_config_init:
		if ConfigProtoHandler.parentClassConfig.ParentOnMemConfUpdate == nil {
			for sectKey, KeysValues := range data.SectionsKeysValues {
				// should not be empty, nobody should send empty message
				ConfigProtoHandler.parentClassConfig.MemConfig[sectKey] = KeysValues.GetKeyValue()
			}
		} else {
			memConfig := make(map[string]map[string]string)
			for sectKey, KeysValues := range data.SectionsKeysValues {
				// should not be empty, nobody should send empty message
				memConfig[sectKey] = KeysValues.GetKeyValue()
			}
			// config update function can be then updated with self.updateMemConfig
			ConfigProtoHandler.parentClassConfig.ParentOnMemConfUpdate(memConfig)
		}

	// receive init parser config
	case pb.ConfigMsg_send_config_init:
		for sectKey, KeysValues := range data.SectionsKeysValues {
			// should not be empty, nobody should send empty message
			if sectKey == COMMON_SECTION {
				ConfigProtoHandler.parentClassConfig.UpdateSelf(KeysValues.GetKeyValue())
			}
			ConfigProtoHandler.parentClassConfig.Parser.Section(sectKey).MapTo(KeysValues.GetKeyValue())
		}

	// receive init notif log level
	case pb.ConfigMsg_send_notif_loglevel_init:
		notifLogLevel := make(map[string]map[string]string)
		for sectKey, KeysValues := range data.SectionsKeysValues {
			// should not be empty, nobody should send empty message
			notifLogLevel[sectKey] = KeysValues.GetKeyValue()
		}
		if ConfigProtoHandler.notifRefreshSender != nil {
			// if parent local notifie
			// async update local notif senders first
			notifLogLevelChan := make(chan map[string][]string)
			go func() {
				notifLogLevelChan <- ConfigProtoHandler.notifRefreshSender(notifLogLevel)
			}()
			// then update logger
			ConfigProtoHandler.loggerRefreshLogLevel(<-notifLogLevelChan)
		} else {
			// parent directly connected to notif_server (only update notif log levels)
			returnedLogLevel := map[string][]string{}
			for _, keyVal := range notifLogLevel {
				for _, logLevel := range PyHelpers.GetSplitedParam(keyVal["LOGLEVEL"]) {
					if val, ok := returnedLogLevel[logLevel]; ok {
						val = append(val, keyVal["TAG"])
					} else {
						returnedLogLevel[logLevel] = []string{keyVal["TAG"]}
					}
				}
			}
			// then update logger
			ConfigProtoHandler.loggerRefreshLogLevel(returnedLogLevel)
		}

	// receive confirmation that server has updated parser config
	case pb.ConfigMsg_config_update_done:
	// ok nothing to do, just waiting for response...

	// received an update error from the server : mem_config update failed
	case pb.ConfigMsg_mem_config_update_failed:
		// This should not occurs oftenly...
		// #FIXME this is not correct
		deserializeErr := errors.New(fmt.Sprintf("config server failure, unable to dumps mem config : RespServer '%s'", pb.ConfigMsg_mem_config_update_failed.String()))
		fmt.Printf("Error while trying to deserialize 'proto.ConfigMsg_RespServer' : unexcepted error -> %s\n", deserializeErr)

	// received an update error from the server : config update failed
	case pb.ConfigMsg_config_update_failed:
		// This should not occurs oftenly...
		// #FIXME this is not correct
		deserializeErr := errors.New(fmt.Sprintf("config server failure, unable to save parser config : RespServer '%s'", pb.ConfigMsg_config_update_failed.String()))
		fmt.Printf("Error while trying to deserialize 'proto.ConfigMsg_RespServer' : unexcepted error -> %s\n", deserializeErr)

	default:
		// unknown message, should not happen !!
		fmt.Printf("Unknown RespServer : '%s', this should not happens !!\n", data.RespServer.String())
		panic(fmt.Sprintf("Unknown RespServer : '%s', this should not happens !!\n", data.RespServer.String()))
	}
}

// client side
// ///////////////////////////////
// set logger.Log func while instanciating MyLogger

func (ConfigProtoHandler *ConfigProtoHandler) SetLoggerLog(LoggerLogCallBack func(string, string)) {
	ConfigProtoHandler.loggerLog = LoggerLogCallBack
}

// set logger.Log func while instanciating MyLogger
/////////////////////////////////
