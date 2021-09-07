package controller

import (
	"assignment.com/collection-tree/model"
)

type DataCollectionController struct {
	Data model.InsertData
}

var Model model.Root

func (value DataCollectionController) RunTask() {
	Model.Mu.Lock()
	value.AddRoot()
	value.AddCountry()
	value.AddDevice()
	Model.Mu.Unlock()
}

func (value DataCollectionController) AddRoot() {
	Model.WebRequest += value.Data.WebRequest
	Model.TimeSpent += value.Data.TimeSpent
}

func (value DataCollectionController) AddCountry() {
	if Model.ChildCountryData == nil {
		temp := make(map[string]*model.Country)
		Model.ChildCountryData = &temp
	}
	if !value.CheckCountry() {
		temp := make(map[string]*model.Device)
		(*Model.ChildCountryData)[value.Data.CountryName] = &model.Country{
			ChildDeviceData: &temp,
			WebRequest:      value.Data.WebRequest,
			TimeSpent:       value.Data.TimeSpent,
			CountryName:     value.Data.CountryName,
		}
	} else {
		(*Model.ChildCountryData)[value.Data.CountryName].WebRequest += value.Data.WebRequest
		(*Model.ChildCountryData)[value.Data.CountryName].TimeSpent += value.Data.TimeSpent
	}
}

func (value DataCollectionController) AddDevice() {
	country := (*Model.ChildCountryData)[value.Data.CountryName]
	if country.ChildDeviceData == nil {
		temp := make(map[string]*model.Device)
		country.ChildDeviceData = &temp
	}
	if !value.CheckDevice() {
		(*country.ChildDeviceData)[value.Data.DeviceName] = &model.Device{
			WebRequest: value.Data.WebRequest,
			TimeSpent:  value.Data.TimeSpent,
			DeviceName: value.Data.DeviceName,
		}
	} else {
		(*country.ChildDeviceData)[value.Data.DeviceName].WebRequest += value.Data.WebRequest
		(*country.ChildDeviceData)[value.Data.DeviceName].TimeSpent += value.Data.TimeSpent
	}
}

func (value DataCollectionController) CheckCountry() bool {
	if Model.ChildCountryData == nil {
		return false
	}
	countryData := (*Model.ChildCountryData)[value.Data.CountryName]
	if countryData == nil {
		return false
	}
	if countryData.CountryName == value.Data.CountryName {
		return true
	}
	return false
}

func (value DataCollectionController) CheckDevice() bool {
	isCountry := value.CheckCountry()
	if !isCountry {
		return false
	}
	deviceData := (*((*Model.ChildCountryData)[value.Data.CountryName]).ChildDeviceData)
	if deviceData == nil {
		return false
	}
	if deviceData[value.Data.DeviceName] == nil {
		return false
	}
	if deviceData[value.Data.DeviceName].DeviceName == value.Data.DeviceName {
		return true
	}
	return false
}

func GetData(country string, device string) (webRequest int64, timeSpent int64) {
	if country == "" && device == "" {
		return Model.WebRequest, Model.TimeSpent
	} else if country != "" && device == "" {
		if *Model.ChildCountryData == nil {
			return -1, -1
		}
		data := (*Model.ChildCountryData)[country]
		return data.WebRequest, data.TimeSpent
	} else if country != "" && device != "" {
		data := (*((*Model.ChildCountryData)[country]).ChildDeviceData)
		if data == nil {
			return -1, -1
		}
		if data[device] == nil {
			return -1, -1
		}
		return data[device].WebRequest, data[device].TimeSpent
	} else {
		if *Model.ChildCountryData == nil {
			return -1, -1
		}
		var webRequest int64
		var timeSpent int64
		for _, value := range *Model.ChildCountryData {
			deviceData := value.ChildDeviceData
			for key1, value1 := range *deviceData {
				if key1 == device {
					webRequest += value1.WebRequest
					timeSpent += value1.TimeSpent
				}
			}
		}
		return webRequest, timeSpent
	}
}
