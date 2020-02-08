package main

import (
	"fmt"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

type messageHandler struct{}

// Mainly caused by another instance of Whatsapp Web being opened
func (messageHandler) HandleError(err error) {
	// Reconnect after a given amount of time
	println("Another instance of Whatsapp Web has been opened. Waiting to try again...")
	time.Sleep(errorTimeout)
	sess, conn = HandleLogin()
}

func (messageHandler) HandleTextMessage(message whatsapp.TextMessage) {

	if message.Info.Timestamp > startTime && jidToName(message.Info.RemoteJid) == conn.Info.Pushname {
		go HandleBotMsg(message, conn)

		println(fmt.Sprintf("%s: %s", jidToName(MessageToJid(message)), message.Text))
	}
}

func (messageHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleJSONMessage(message string) {
	//	fmt.Println(message)
}

func (messageHandler) HandleContactMessage(message whatsapp.ContactMessage) {
	//fmt.Println(message)
}

/*	authorID := "-"

	if len(message.Info.Source.GetPushName()) > 0 {
		authorID = message.Info.Source.GetPushName()
	} else if message.Info.Source.Participant != nil {
		authorID = *message.Info.Source.Participant
	} else {
		authorID = message.Info.RemoteJid // Personennamen
	}

	println(fmt.Sprintf("%s: %s", jidToName(authorID), message.Text))
*/
