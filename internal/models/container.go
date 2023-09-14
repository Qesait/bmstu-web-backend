package models

var DRY_CUBE_20 = Dimentions{
	Length: 6058,
	Width:  2438,
	Height: 2591,
}

var HIGH_CUBE_20 = Dimentions{
	Length: 6058,
	Width:  2438,
	Height: 2896,
}

type Dimentions struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
	Length uint `json:"length"`
}

type Cargo struct {
	Name   string `json:"name"`
	Weight uint   `json:"weight"`
}

type Container struct {
	SerialNumber string     `json:"serial_number"`
	Type         string     `json:"type"`
	Dimentions   Dimentions `json:"dimentions"`
	ImageURL     string     `json:"image_url"`
	Cargo        Cargo      `json:"cargo"`
}
