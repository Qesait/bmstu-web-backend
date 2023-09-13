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
	Name    string `json:"name"`
	Weight  uint   `json:"weight"`
}

type Container struct {
	SerialNumber uint       `json:"serial_number"`
	Type         string     `json:"type"`
	ImageURL     string     `json:"image_url"`
	Dimentions   Dimentions `json:"dimentions"`
	Cargo        Cargo      `json:"cargo"`
}
