package wabot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Rhymen/go-whatsapp"
)

var users = BotUserList{}

// BotUserList saves the BotUser-Array - easy to save&load
type BotUserList struct {
	BotUsers []BotUser
}

// BotUser contains the contact and his personal settings
type BotUser struct {
	Contact  whatsapp.Contact
	Settings map[string]interface{}
}

// AddUser adds a new member to the group and prepares a Settings struct for him
func AddUser(user whatsapp.Contact) {
	users.BotUsers = append(users.BotUsers, BotUser{Contact: user})
}

// CreateNewSettingsOption adds the interface to the general BotUser struct
func CreateNewSettingsOption(identifier string, options interface{}) {
	for i := 0; i < len(users.BotUsers); i++ {
		users.BotUsers[i].Settings[identifier] = options
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

// AddUserByJid - AddUser alternative  // NOT WORKING RN
func AddUserByJid(jid string) {
	if !IsUserRegistered(jid) {
		for _, c := range contacList {
			if c.Jid == jid {
				users.BotUsers = append(users.BotUsers, BotUser{Contact: c})
				break
			}
		}
	}
}

// DoesUserExist checks if the given jid exists in the user Array
func DoesUserExist(jid string) bool {
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return true
		}
	}
	return false
}

// GetAllUserSettings returns the settings of a specific user
func GetAllUserSettings(jid string) map[string]interface{} {
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

// GetUserSettings returns the settings of a specific user
func GetUserSettings(settingsIdentifier string, jid string) interface{} {
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return u.Settings[settingsIdentifier]
		}
	}

	// User isn't registered yet. Do it now!
	AddUserByJid(jid)
	// Return the settings
	for _, u := range users.BotUsers {
		if u.Contact.Jid == jid {
			return u.Settings[settingsIdentifier]
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
func SaveUsersToDisk(path string) {
	if len(users.BotUsers) > 0 {
		usersJSON, _ := json.Marshal(users)

		usersJSON = EncryptData(usersJSON)

		ioutil.WriteFile(path, usersJSON, 0600)
		println("--- Users saved ---")
	} else {
		println("---Save failed, no entries---")
	}
}

// LoadUsersFromDisk loads the BotUser-Slice
// - returns false if no data could be loaded
func LoadUsersFromDisk(path string) bool {
	savedUsers := BotUserList{}

	savedData, err := ioutil.ReadFile(path)
	if err == nil {
		savedData = DecryptData(savedData)
		err = json.Unmarshal(savedData, &savedUsers)
		if err == nil {
			users = savedUsers
			return true
		}
	}

	return false
}
