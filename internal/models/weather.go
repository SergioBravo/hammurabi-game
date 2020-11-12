package models

// WeatherAPIResponse ...
type WeatherAPIResponse struct {
	Base       string `json:"base"`
	Visibility int    `json:"visibility"`
	Dt         int    `json:"dt"`
	Timezone   int    `json:"timezone"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`

	Coord   Coord   `json:"coord"`
	Weather Weather `json:"weather"`
	Main    Main    `json:"main"`
	Wind    Wind    `json:"wind"`
	Clouds  Clouds  `json:"clouds"`
	Sys     System  `json:"sys"`
}

// Coord ...
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Clouds ...
type Clouds struct {
	All int `json:"all"`
}

// Main ...
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Weather ...
type Weather []struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Wind ...
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `jsons:"deg"`
}

// System ...
type System struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
