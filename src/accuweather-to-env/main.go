package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

/*
Sun with rays:

Black sun with rays	☀	U+2600	&#9728;	Clear weather
Cloud	☁	U+2601	&#9729;	Cloud, cloudy weather
Umbrella	☂	U+2602	&#9730;	Umbrella, rainy weather

⛈ &#x26C8;
🌧	Cloud With Rain	&#x1F327;
🌩	Cloud With Lightning	&#x1F329;
🌫	Fog	&#x1F32B
❄	Snowflake	&#10052;
*/

// Accuweather API for 5 days forecast
// Param 1 = location key
// Param 2 = API key
// Param 3 = boolean, true=metric or false=imperial
const accuweatherAPI5DaysForecast = "http://dataservice.accuweather.com/forecasts/v1/daily/5day/%s?apikey=%s&metric=%s"

const weatherForecast = "☀ Sunny, small clouds / 13°C / 24°C"

func main() {
	apiKey, locationKey, metric, err := readCommandLineArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(weatherForecast)
	// os.Setenv("WEATHER_DAY1", weatherForecast)
	retrieveWeatherData(apiKey, locationKey, metric)
}

func readCommandLineArgs() (*string, *string, *bool, *string) {
	if len(os.Args) != 4 {
		return nil, nil, nil, "Invalid nr of arguments. Call with: <apiKey> <locationKey> <true for metric, false for imperial>"
	}
	var apiKey = os.Args[1]
	var locationKey = os.Args[2]
	var metric, _ = strconv.ParseBool(os.Args[3])

	return apiKey, locationKey, metric, nil
}

func retrieveWeatherData(apiKey string, locationKey string, metric bool) {
	response, err := http.Get(fmt.Sprintf(accuweatherAPI5DaysForecast, locationKey, apiKey, strconv.FormatBool(metric)))
	if err != nil {
		fmt.Println("Cannot call the Accuweather 5Days API", err)
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Success Accuweather API call: " + string(data))
}
