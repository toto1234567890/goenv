package Notifie

import (
	"fmt"

	notifMsg "govenv/api/capnp/notifieMsg"
	"govenv/pkg/common/Config"

	capnplib "capnproto.org/go/capnp/v3"
)

type NotifMessage struct {
	Message    string
	Attachment string
	Tags       []string
}

type NotifNcapHandler struct {
	Name         string
	config       *Config.Config
	notifMessage *notifMsg.NotifieMsg
	memSeg       *capnplib.Segment
	msgSerDeSer  *capnplib.Message
}

func NewNotifHandler(name string, parentClassConfig *Config.Config) *NotifNcapHandler {
	capnplibMsg, memSeg, err := capnplib.NewMessage(capnplib.SingleSegment(nil))
	if err != nil {
		panic(fmt.Sprintf("Error while trying to initialize Notif Handler :'%v'\n", err))
	}

	notifObj, err := notifMsg.NewRootNotifieMsg(memSeg)
	if err != nil {
		panic(fmt.Sprintf("Error while trying to initialize Notif Handler :'%v'\n", err))
	}
	return &NotifNcapHandler{Name: name, config: parentClassConfig, memSeg: memSeg, notifMessage: &notifObj, msgSerDeSer: capnplibMsg}
}

var capnpList capnplib.TextList

func (notifNcapHandler *NotifNcapHandler) NotifNcapSerialize(notifMessage *NotifMessage) []byte {
	notifNcapHandler.notifMessage.SetMessage_(notifMessage.Message)
	notifNcapHandler.notifMessage.SetAttachment(notifMessage.Attachment)
	for i, val := range notifMessage.Tags {
		capnpList.Set(i, val)
	}
	notifNcapHandler.notifMessage.SetTags(capnpList)
	byteMsg, _ := notifNcapHandler.msgSerDeSer.MarshalPacked()
	return byteMsg
}

func (notifNcapHandler *NotifNcapHandler) NotifNcapDeSerialize(data []byte) *NotifMessage {
	capnpMessage, _ := capnplib.UnmarshalPacked(data)
	goObj, _ := notifMsg.ReadRootNotifieMsg(capnpMessage)
	notifMessage := &NotifMessage{}
	notifMessage.Message, _ = goObj.Message_()
	notifMessage.Attachment, _ = goObj.Attachment()
	tagList, _ := goObj.Tags()
	for i := 0; i < tagList.Len(); i++ {
		notifMessage.Tags[i], _ = tagList.At(i)
	}
	return notifMessage
}
