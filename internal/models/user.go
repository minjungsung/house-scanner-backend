package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

type User struct {
	Name           string `gorm:"column:name"`
	Phone          string `gorm:"column:phone"`
	Email          string `gorm:"column:email;uniqueIndex"`
	HashedPassword string `gorm:"column:hashed_password"`
	Address        string `gorm:"column:address"`
	AddressDetail  string `gorm:"column:address_detail"`
	Birthday       string `gorm:"column:birthday"`
	Message        string `gorm:"column:message"`
}