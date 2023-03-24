package apis

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetWeather(city string) any {

	API_KEY := os.Getenv("WEATHER_API")
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
	fmt.Println(string(body))

	return string(body)
}
