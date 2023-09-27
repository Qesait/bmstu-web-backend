package models

type dimentions struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
	Length uint `json:"length"`
}

type ContainerType struct {
	Name       string     `json:"name"`
	Dimentions dimentions `json:"dimentions"`
	MaxGross   uint       `json:"max_gross"`
	Tare       uint       `json:"tare"`
}

type Cargo struct {
	Name   string `json:"name"`
	Weight uint   `json:"weight"`
}

type Container struct {
	// owner code - 3 uppercase Latin letters
	// equipment category - 1 uppercase Latin letter (U, J, Z, R)
	// serial number - 6 digits
	// check digit
	Id              string        `json:"id"`
	Type            ContainerType `json:"type"`
	ImageURL        string        `json:"image_url"`
	Cargo           Cargo         `json:"cargo"`
	CurrentLocation string
}


var DRY_FREIGHT_20 = ContainerType{
	Name: "Стандартный 20-ти футовый контейнер",
	Dimentions: dimentions{
		Length: 6058,
		Width:  2438,
		Height: 2591,
	},
	Tare:     2230,
	MaxGross: 21770,
}

var HIGH_CUBE_20 = ContainerType{
	Name: "20 футовый контейнер увеличенной высоты",
	Dimentions: dimentions{
		Length: 6058,
		Width:  2438,
		Height: 2896,
	},
	Tare:     2350,
	MaxGross: 21650,
}

var DRY_FREIGHT_40 = ContainerType{
	Name: "Стандартный 40-ка футовый контейнер",
	Dimentions: dimentions{
		Length: 12192,
		Width:  2438,
		Height: 2591,
	},
	Tare:     3780,
	MaxGross: 26700,
}

var HIGH_CUBE_40 = ContainerType{
	Name: "40 футовый контейнер увеличенной высоты",
	Dimentions: dimentions{
		Length: 12192,
		Width:  2438,
		Height: 2896,
	},
	Tare:     3780,
	MaxGross: 26700,
}