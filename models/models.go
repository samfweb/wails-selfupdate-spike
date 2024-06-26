package models

import (
	"time"

	"gorm.io/gorm"
)

type Connection struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	Name           string         `json:"name"`
	Protocol       string         `json:"protocol"`
	Host           string         `json:"host"`
	Port           int            `json:"port"`
	Username       JsonNullString `json:"username"`
	Password       JsonNullString `json:"password"`
	IsProtoEnabled bool           `json:"isProtoEnabled"`
	ProtoRegDir    JsonNullString `json:"protoRegDir"`
	IsCertsEnabled bool           `json:"isCertsEnabled"`
	CertCa         JsonNullString `json:"certCa"`
	CertClient     JsonNullString `json:"certClient"`
	CertClientKey  JsonNullString `json:"certClientKey"`
	Subscriptions  []Subscription `json:"subscriptions"`
}

type Subscription struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	ConnectionID    uint           `json:"connectionId"`
	QoS             uint           `json:"qos"`
	Topic           string         `json:"topic"`
	ProtoDescriptor string         `json:"protoDescriptor"`
}
