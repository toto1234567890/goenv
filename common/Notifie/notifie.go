package Notifie

import (
	"govenv/pkg/common/Config"
	"govenv/pkg/common/PyHelpers"
)

// ce que doit faire le Notifie :
//
// soit instancié dans le logger (withNotifServer = false)
// 		- chargement des notifs serveur selon la config (notif par défaut ajouté si rien dans la config ??)
// 		- le logger place les notifs Message dans la NotifChan
// 		- processMessage (goroutine) récupére les messages et les envoi (selon le tag)
// soit instancié dans le notif_server
// 		- instanciation avec le notif_server, même fonctionnement que l'instanciation avec le logger
//		- dans ce cas le logger n'utilise pas cette class...
//      - le logger utilise sa propre NotifChan et son 'serializer socket' pour envoyer les messages au notif_server...
//

/////////////////////////////////
// Notifie

// when called by logger and if notif_server, this class is not called...
type Notifie struct {
	Name           string
	config         *Config.Config
	TagToSenderMap map[string]NotifSenderInterface
	NotifChan      chan *NotifMessage
}

// This class is called by logger only with local notifie or notif_server
func NewNotifie(config *Config.Config, parentName string) *Notifie {
	curNotifie := &Notifie{Name: parentName, config: config, NotifChan: make(chan *NotifMessage)}
	// set callback to init config
	curNotifie.config.Handler.SetNotifCallBack(curNotifie.loadNotifSender)
	// init senders and send back notifLogLevel to logger
	config.GetNotifConfig()
	go curNotifie.processMessage()
	return curNotifie
}

func (notifie *Notifie) loadNotifSender(notifiersConf map[string]map[string]string) map[string][]string {
	errorList := []string{}
	returnLogLevelByTag := map[string][]string{}
	// temp dict used to update missing setionKeyValues in "LevelByTag", if some missing...
	// this should not occurs, raise a warning if it happens...
	// LevelByTagMap := make(map[string]map[string]string)

	confName := "TELEGRAM"
	if confTele, ok := notifiersConf[confName]; ok {
		telegram, telegramErrorList := NewTelegramSender(confTele, confName)
		if telegramErrorList != "" {
			errorList = append(errorList, telegramErrorList)
		} else {
			for _, logLevel := range PyHelpers.GetSplitedParam(telegram.logLevel) {
				if curConf, ok := returnLogLevelByTag[logLevel]; ok {
					curConf = append(curConf, telegram.tag)
				} else {
					returnLogLevelByTag[logLevel] = []string{telegram.tag}
				}
			}
			var telegramInterface NotifSenderInterface = telegram
			notifie.TagToSenderMap[telegram.tag] = telegramInterface
		}
	}
	confName = "DISCORD"
	if confDisco, ok := notifie.config.MemConfig[confName]; ok {
		discord, discordErrorList := NewDiscordSender(confDisco, confName)
		if discordErrorList != "" {
			errorList = append(errorList, discordErrorList)
		} else {
			for _, logLevel := range PyHelpers.GetSplitedParam(discord.logLevel) {
				if curConf, ok := returnLogLevelByTag[logLevel]; ok {
					curConf = append(curConf, discord.tag)
				} else {
					returnLogLevelByTag[logLevel] = []string{discord.tag}
				}
			}
			var discordInterface NotifSenderInterface = discord
			notifie.TagToSenderMap[discord.tag] = discordInterface
		}
	}
	confName = "MATRIX"
	if confMatrix, ok := notifie.config.MemConfig[confName]; ok {
		matrix, matrixErrorList := NewMatrixSender(confMatrix, confName)
		if matrixErrorList != "" {
			errorList = append(errorList, matrixErrorList)
		} else {
			for _, logLevel := range PyHelpers.GetSplitedParam(matrix.logLevel) {
				if curConf, ok := returnLogLevelByTag[logLevel]; ok {
					curConf = append(curConf, matrix.tag)
				} else {
					returnLogLevelByTag[logLevel] = []string{matrix.tag}
				}
			}
			var matrixInterface NotifSenderInterface = matrix
			notifie.TagToSenderMap[matrix.tag] = matrixInterface
		}
	}
	confName = "GMAIL"
	if confGmail, ok := notifie.config.MemConfig[confName]; ok {
		gmail, gmailErrorList := NewGmailSender(confGmail, confName)
		if gmailErrorList != "" {
			errorList = append(errorList, gmailErrorList)
		} else {
			for _, logLevel := range PyHelpers.GetSplitedParam(gmail.logLevel) {
				if curConf, ok := returnLogLevelByTag[logLevel]; ok {
					curConf = append(curConf, gmail.tag)
				} else {
					returnLogLevelByTag[logLevel] = []string{gmail.tag}
				}
			}
			var gmailInterface NotifSenderInterface = gmail
			notifie.TagToSenderMap[gmail.tag] = gmailInterface
		}
	}
	// if no config sender found load telegram default one
	//if len(notifie.TagToSenderMap) == 0 {
	//}
	return returnLogLevelByTag
}

func (notifie *Notifie) processMessage() {
	var recvNotifMessage *NotifMessage
	recvNotifMessage = <-notifie.NotifChan
	for _, tag := range recvNotifMessage.Tags {
		// attachement and notif.Name (FIXME notif.Name  will be changed probably) only used with email sender...
		go notifie.TagToSenderMap[tag].SendMessage(recvNotifMessage.Message, recvNotifMessage.Attachment, notifie.Name)
	}
}

// Notifie
/////////////////////////////////
