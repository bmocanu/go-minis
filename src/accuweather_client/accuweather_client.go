package accuweather_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// URL of Accuweather API for 5 days forecast
// Param 1 = location key
// Param 2 = API key
// Param 3 = boolean, true=metric or false=imperial
const accuweatherAPI5DaysForecast = "http://dataservice.accuweather.com/forecasts/v1/daily/5day/%s?apikey=%s&metric=%s"

// Get5DaysForecast calls the same-name API from Accuweather, retrieving the weather forecast
// for the next 5 days
func Get5DaysForecast(config AccuweatherConfig, metric bool) (Accuweather5DaysForecast, error) {
	var result Accuweather5DaysForecast
	response, err := http.Get(fmt.Sprintf(accuweatherAPI5DaysForecast, config.LocationKey, config.ApiKey, strconv.FormatBool(metric)))
	if err != nil {
		fmt.Println("Cannot call the Accuweather 5Days API", err)
		return result, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Success Accuweather API call: " + string(data))
	json.Unmarshal([]byte(data), &result)
	return result, nil
}
