package models

import (
	"fmt"

	"github.com/Bernar11296/rest/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (c *Contact) Validate() (map[string]interface{}, bool) {
	if c.Name == "" {
		return utils.Message(false, "Contact name should be on the payload"), false
	}
	if c.Phone == "" {
		return utils.Message(false, "Phone number should be on the payload"), false
	}
	if c.UserId <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}
	return utils.Message(true, "succes"), true
}

func (c *Contact) Create() map[string]interface{} {
	if resp, ok := c.Validate(); !ok {
		return resp
	}
	GetDB().Create(c)
	resp := utils.Message(true, "succes")
	resp["contact"] = c
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}

	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
