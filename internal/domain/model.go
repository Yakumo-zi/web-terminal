package domain

import "time"

type Asset struct {
	Id          string       `json:"id"`
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
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Members   []Asset           `json:"members"`
	Attr      map[string]string `json:"attr"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Credential struct {
	Id        string    `json:"id"`
	Asset     Asset     `json:"asset"`
	Type      string    `json:"type"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	Id         string            `json:"id"`
	Asset      Asset             `json:"asset"`
	Credential Credential        `json:"credential"`
	Status     string            `json:"status"`
	Type       string            `json:"type"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	StopAt     time.Time         `json:"stop_at"`
	SourceInfo map[string]string `json:"source_info"`
}
