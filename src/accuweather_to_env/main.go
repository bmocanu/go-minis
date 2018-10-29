package main

import (
	_ "accuweather_client"
	"errors"
	"fmt"
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

const weatherForecast = "☀ Sunny, small clouds / 13°C / 24°C"

func main() {
	apiKey, locationKey, metric, err := readCommandLineArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	var awConfig AccuweatherConfig
	awConfig.ApiKey = apiKey
	awConfig.LocationKey = locationKey

	aw5DaysForecast, err := Get5DaysForecast(awConfig, metric)
	if err != nil {
		fmt.Println("Error occurred while calling AccuWeather for the 5 days forecast", err)
		return
	}

	fmt.Println(fmt.Sprintf("Forecast for tomorrow: %s / %s",
		aw5DaysForecast.DailyForecasts[0].Temperature.Minimum.Value,
		aw5DaysForecast.DailyForecasts[0].Temperature.Maximum.Value))

	// fmt.Println(weatherForecast)
	// os.Setenv("WEATHER_DAY1", weatherForecast)
}

func readCommandLineArgs() (string, string, bool, error) {
	if len(os.Args) != 4 {
		return "", "", false, errors.New("invalid nr of arguments. Call with: <apiKey> <locationKey> <true for metric, false for imperial>")
	}
	var apiKey = os.Args[1]
	var locationKey = os.Args[2]

	var metric, err = strconv.ParseBool(os.Args[3])
	if err != nil {
		return "", "", false, err
	}

	return apiKey, locationKey, metric, nil
}
