package wabot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Rhymen/go-whatsapp"
)

var users = BotUserList{}

// BotUserList saves the BotUser-Array - easy to save&load
type BotUserList struct {
	BotUsers []*BotUser
}

// BotUser contains the contact and his personal settings
type BotUser struct {
	Contact  whatsapp.Contact
	Nickname string
	Settings interface{}
}

var standardSettings interface{}

// addUser adds a new member to the group and prepares a Settings struct for him
func addUser(user whatsapp.Contact) {
	users.BotUsers = append(users.BotUsers, &BotUser{Contact: user, Nickname: "", Settings: standardSettings})
}

// CreateNewSettingsOption adds the interface to the general BotUser struct
func CreateNewSettingsOption(settings interface{}) {
	standardSettings = settings
}

// GetUserNickname returns a users nickname
func GetUserNickname(jid string) string {
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return u.Nickname
		}
	}
	return ""
}

// SetUserNickname sets a users nickname
func SetUserNickname(jid string, nickname string) {
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			u.Nickname = nickname
			return
		}
	}

	// User isn't registered yet. Do it now!
	AddUserByJid(jid)
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			u.Nickname = nickname
		}
	}
}

// IsUserRegistered checks if the given jid exists in the array
func IsUserRegistered(jid string) bool {
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return true
		}
	}
	return false
}

// AddUserByJid - AddUser alternative
func AddUserByJid(jid string) {
	if !IsUserRegistered(jid) {
		for _, c := range contacList {
			if c.Jid == jid {
				users.BotUsers = append(users.BotUsers, &BotUser{Contact: c, Nickname: "", Settings: standardSettings})
				break
			}
		}
	}
}

// ChangeUserSettings a users settings
func ChangeUserSettings(jid string, settings interface{}) {
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			u.Settings = settings
		}
	}

	// User isn't registered yet. Do it now!
	AddUserByJid(jid)
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			u.Settings = settings
		}
	}
}

// GetUserSettings returns the settings of a specific user
func GetUserSettings(jid string) interface{} {
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return u.Settings
		}
	}

	// User isn't registered yet. Do it now!
	AddUserByJid(jid)
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return u.Settings
		}
	}
	return nil
}

// GetUserIndex returns a Users index within the Array
func GetUserIndex(message whatsapp.TextMessage) int {
	for i, u := range users.BotUsers {
		if u.Contact.Jid == MessageToJid(message) {
			return i
		}
	}

	return -1
}

// SaveUsersToDisk saves the BotUser-Slice
func SaveUsersToDisk() bool {
	if len(users.BotUsers) > 0 {
		usersJSON, _ := json.Marshal(users)

		usersJSON = encryptData(usersJSON)

		ioutil.WriteFile(usersFile, usersJSON, 0600)
		return true
	}
	return false
}

// GetSaveData returns the stored savedata
// -> bool is false if there was no data
func GetSaveData() (BotUserList, bool) {
	savedUsers := BotUserList{}

	savedData, err := ioutil.ReadFile(usersFile)
	if err == nil {
		savedData = decryptData(savedData)
		err = json.Unmarshal(savedData, &savedUsers)
		if err == nil {
			return savedUsers, true
		}
	}
	return BotUserList{}, false
}

// UseSaveData loads the given data
func UseSaveData(savedata BotUserList) {
	users.BotUsers = savedata.BotUsers
}
