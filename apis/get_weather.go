package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Omega014/weather-tool/viewmodel"
)

func GetWeather(city string) *viewmodel.Weather {

	API_KEY := os.Getenv("WEATHER_API_KEY")
	BASE_URL := "http://api.openweathermap.org/data/2.5/weather"

	res, err := http.Get(BASE_URL + "?q=" + city + ",jp&appid=" + API_KEY)
	if err != nil {
		fmt.Printf("Can not get: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Can not read body: %v", err)
	}

	jsonBytes := ([]byte)(body)
	data := new(viewmodel.Weather)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Printf("Can not unmarshal: %v", err)
	}

	// fmt.Println(string(body))

	return data
}
