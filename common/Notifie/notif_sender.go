package Notifie

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
)

/////////////////////////////////
// Main interface

type NotifSenderInterface interface {
	SendMessage(msg, to, subject string) error
}

// Main interface
/////////////////////////////////
// Notif sender

type TelegramSender struct {
	tag string
	// only used for telegram
	telegramUrl string // -> tgram://{TOKEN}/{CHAT_ID}
	token       string
	chatId      string
	// Level
	logLevel string
}

func NewTelegramSender(telegramConf map[string]string, confName string) (*TelegramSender, string) {
	curError := ""
	telegranSender := &TelegramSender{}
	if tag, ok := telegramConf["TAG"]; ok {
		telegranSender.tag = tag
	} else {
		curError += fmt.Sprintf("missing 'TAG' option for config '%s'\n", confName)
	}
	if token, ok := telegramConf["TOKEN"]; ok {
		telegranSender.token = token
	} else {
		curError += fmt.Sprintf("missing 'TOKEN' option for config '%s'\n", confName)
	}
	if chatId, ok := telegramConf["CHATID"]; ok {
		telegranSender.chatId = chatId
	} else {
		curError += fmt.Sprintf("missing 'CHATID' option for config '%s'\n", confName)
	}
	if logLevel, ok := telegramConf["LOGLEVEL"]; ok {
		telegranSender.logLevel = logLevel
	} else {
		curError += fmt.Sprintf("missing 'LOGLEVEL' option for config '%s'\n", confName)
	}
	if curError == "" {
		telegranSender.telegramUrl = fmt.Sprintf("tgram://%s/%s", telegranSender.token, telegranSender.chatId)
		return telegranSender, ""
	}
	return nil, curError
}

func (telegramSender *TelegramSender) SendMessage(msg, notUsed, notUsedAlso string) error {
	jsonByteMessage, err := json.Marshal(map[string]string{"chat_id": telegramSender.chatId, "text": msg})
	if err != nil {
		return fmt.Errorf("failed to marshall message (telegram): %v", err)
	}
	httpsResp, err := http.Post(telegramSender.telegramUrl, "application/json", bytes.NewBuffer(jsonByteMessage))
	if err != nil {
		return fmt.Errorf("failed to post http request (telegram): %v", err)
	}
	defer httpsResp.Body.Close()
	if httpsResp.StatusCode != http.StatusOK {
		return fmt.Errorf(fmt.Sprintf("telegram, unexpected http status (telegram): %s", httpsResp.StatusCode))
	}
	return nil
}

/////////////////////////////////

type DiscordSender struct {
	tag string
	// only used for discord
	discordUrl string // discord://WebhookID/WebhookToken/ direct url from the discord Room
	// Level
	logLevel string
}

func NewDiscordSender(discordConf map[string]string, confName string) (*DiscordSender, string) {
	curError := ""
	discordSender := &DiscordSender{}
	if tag, ok := discordConf["TAG"]; ok {
		discordSender.tag = tag
	} else {
		curError += fmt.Sprintf("missing 'TAG' option for config '%s'\n", confName)
	}
	if discordUrl, ok := discordConf["URL"]; ok {
		discordSender.discordUrl = discordUrl
	} else {
		curError += fmt.Sprintf("missing 'URL' option for config '%s'\n", confName)
	}
	if logLevel, ok := discordConf["LOGLEVEL"]; ok {
		discordSender.logLevel = logLevel
	} else {
		curError += fmt.Sprintf("missing 'LOGLEVEL' option for config '%s'\n", confName)
	}
	if curError == "" {
		return discordSender, ""
	}
	return nil, curError
}

func (discordSender *DiscordSender) SendMessage(msg, notUsed, notUsedAlso string) error {
	jsonByteMessage, err := json.Marshal(map[string]string{"content": msg})
	if err != nil {
		return fmt.Errorf("failed to marshall message (discord): %v", err)
	}
	httpsResp, err := http.Post(discordSender.discordUrl, "application/json", bytes.NewBuffer(jsonByteMessage))
	if err != nil {
		return fmt.Errorf("failed to post http request (discord): %v", err)
	}
	defer httpsResp.Body.Close()
	if httpsResp.StatusCode != http.StatusOK {
		return fmt.Errorf(fmt.Sprintf("unexpected http status (discord): %s", httpsResp.StatusCode))
	}
	return nil
}

/////////////////////////////////

type MatrixSender struct {
	tag string
	// only used for discord
	matrixUrl string
	// Level
	logLevel string
}

func NewMatrixSender(matrixConf map[string]string, confName string) (*MatrixSender, string) {
	curError := ""
	matrixSender := &MatrixSender{}
	if tag, ok := matrixConf["TAG"]; ok {
		matrixSender.tag = tag
	} else {
		curError += fmt.Sprintf("missing 'TAG' option for config '%s'\n", confName)
	}
	if matrixUrl, ok := matrixConf["URL"]; ok {
		matrixSender.matrixUrl = matrixUrl
	} else {
		curError += fmt.Sprintf("missing 'URL' option for config '%s'\n", confName)
	}
	if logLevel, ok := matrixConf["LOGLEVEL"]; ok {
		matrixSender.logLevel = logLevel
	} else {
		curError += fmt.Sprintf("missing 'LOGLEVEL' option for config '%s'\n", confName)
	}
	if curError == "" {
		return matrixSender, ""
	}
	return nil, curError
}

func (matrixSender *MatrixSender) SendMessage(msg, notUsed, notUsedAlso string) error {
	jsonByteMessage, err := json.Marshal(map[string]string{"content": msg})
	if err != nil {
		return fmt.Errorf("failed to marshall message (matrix): %v", err)
	}
	httpsResp, err := http.Post(matrixSender.matrixUrl, "application/json", bytes.NewBuffer(jsonByteMessage))
	if err != nil {
		return fmt.Errorf("failed to post http request (matrix): %v", err)
	}
	defer httpsResp.Body.Close()
	if httpsResp.StatusCode != http.StatusOK {
		return fmt.Errorf(fmt.Sprintf("unexpected http status (matrix): %s", httpsResp.StatusCode))
	}
	return nil
}

/////////////////////////////////

type GmailSender struct {
	tag string
	// used for email
	from string // one sender-> one receiver
	to   string // one sender-> one receiver
	// config
	smtp   string
	port   int32
	passwd string
	// Level
	logLevel string
}

func NewGmailSender(gmailConf map[string]string, confName string) (*GmailSender, string) {
	curError := ""
	gmailSender := &GmailSender{}
	if tag, ok := gmailConf["TAG"]; ok {
		gmailSender.tag = tag
	} else {
		curError += fmt.Sprintf("missing 'TAG' option for config '%s'\n", confName)
	}
	if gmailFrom, ok := gmailConf["FROM"]; ok {
		gmailSender.from = gmailFrom
	} else {
		curError += fmt.Sprintf("missing 'FROM' option for config '%s'\n", confName)
	}
	if gmailTo, ok := gmailConf["TO"]; ok {
		gmailSender.logLevel = gmailTo
	} else {
		curError += fmt.Sprintf("missing 'TO' option for config '%s'\n", confName)
	}
	if gmailPasswd, ok := gmailConf["PASSWD"]; ok {
		gmailSender.passwd = gmailPasswd
	} else {
		curError += fmt.Sprintf("missing 'PASSWD' option for config '%s'\n", confName)
	}
	if logLevel, ok := gmailConf["LOGLEVEL"]; ok {
		gmailSender.logLevel = logLevel
	} else {
		curError += fmt.Sprintf("missing 'LOGLEVEL' option for config '%s'\n", confName)
	}
	if curError == "" {
		gmailSender.smtp = "smtp.gmail.com"
		gmailSender.port = 587
		return gmailSender, ""
	}
	return nil, curError
}

func (gmailSender *GmailSender) SendMessage(subject, attachment, body string) error {
	// prepare attachment
	attachmentBytes, err := os.ReadFile(attachment)
	if err != nil {
		return fmt.Errorf("failed to read attachment (gmail): %v", err)
	}
	attachmentName := filepath.Base(attachment)
	encodedAttachment := base64.StdEncoding.EncodeToString(attachmentBytes)
	// create email buffer
	var email bytes.Buffer
	writer := multipart.NewWriter(&email)
	// Add email headers
	email.WriteString(fmt.Sprintf("From: %s\r\n", gmailSender.from))
	email.WriteString(fmt.Sprintf("To: %s\r\n", gmailSender.to))
	email.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	email.WriteString("MIME-Version: 1.0\r\n")
	email.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", writer.Boundary()))
	email.WriteString("\r\n")
	// Add email body
	bodyPart, err := writer.CreatePart(
		map[string][]string{
			"Content-Type": {"text/plain; charset=\"utf-8\""},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create body part (gmail): %v", err)
	}
	bodyPart.Write([]byte(body))
	// Add attachment
	attachmentPart, err := writer.CreatePart(
		map[string][]string{
			"Content-Type":              {fmt.Sprintf("application/octet-stream; name=\"%s\"", attachmentName)},
			"Content-Transfer-Encoding": {"base64"},
			"Content-Disposition":       {fmt.Sprintf("attachment; filename=\"%s\"", attachmentName)},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create attachment part (gmail): %v", err)
	}
	attachmentPart.Write([]byte(encodedAttachment))
	// Close multipart writer
	writer.Close()
	// Create authentication
	auth := smtp.PlainAuth("", gmailSender.from, gmailSender.passwd, gmailSender.smtp)
	// Send email
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", gmailSender.smtp, gmailSender.port),
		auth,
		gmailSender.from,
		[]string{gmailSender.to},
		email.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("error while trying to send email (gmail): %v", err)
	}
	return nil
}

/////////////////////////////////
