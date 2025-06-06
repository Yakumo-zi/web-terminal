package domain

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	Id          uuid.UUID    `json:"id"`
	Type        string       `json:"type"`
	Name        string       `json:"name"`
	Port        int          `json:"port"`
	Ip          string       `json:"ip"`
	Groups      []AssetGroup `json:"groups"`
	Credentials []Credential `json:"credentials"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type AssetGroup struct {
	Id        uuid.UUID             `json:"id"`
	Name      string                `json:"name"`
	Members   []Asset               `json:"members"`
	Attr      []AssetGroupAttribute `json:"attrs"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type AssetGroupAttribute struct {
	Id        int       `json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Credential struct {
	Id        uuid.UUID `json:"id"`
	Asset     Asset     `json:"asset"`
	Type      string    `json:"type"`
	Secret    string    `json:"secret"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	Id         uuid.UUID  `json:"id"`
	Asset      Asset      `json:"asset"`
	Credential Credential `json:"credential"`
	Status     string     `json:"status"`
	Type       string     `json:"type"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	StopedAt   time.Time  `json:"stoped_at"`
}
