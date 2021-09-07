package model

import (
	"sync"
)

type InsertData struct {
	CountryName string `json:"countryName"`
	DeviceName  string `json:"deviceName"`
	WebRequest  int64  `json:"webRequest"`
	TimeSpent   int64  `json:"timeSpent"`
}

type FetchData struct {
	CountryName string `json:"countryName"`
	DeviceName  string `json:"deviceName"`
}

type Root struct {
	Mu               sync.Mutex
	ChildCountryData *map[string]*Country
	WebRequest       int64
	TimeSpent        int64
}

type Country struct {
	ChildDeviceData *map[string]*Device
	WebRequest      int64
	TimeSpent       int64
	CountryName     string
}

type Device struct {
	DeviceName string
	WebRequest int64
	TimeSpent  int64
}
