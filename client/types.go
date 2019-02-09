package client

type Location struct {
	ContactBookEnabled bool        `json:"contactBookEnabled,omitempty"`
	Name               string      `json:"name,omitempty"`
	TemperatureScale   string      `json:"temperatureScale,omitempty"`
	ZipCode            string      `json:"zipCode,omitempty"`
	Latitude           string      `json:"latitude,omitempty"`
	Longitude          string      `json:"longitude,omitempty"`
	TimeZone           string      `json:"timeZone,omitempty"`
	CurrentMode        CurrentMode `json:"currentMode,omitempty"`
	Hubs               []Hub       `json:"hubs,omitempty"`
}

type CurrentMode struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Hub struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
