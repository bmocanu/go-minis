package accuweather_client

type AccuweatherConfig struct {
	ApiKey      string
	LocationKey string
}

type Accuweather5DaysForecast struct {
	Headline       AccuweatherHeadline
	DailyForecasts []AccuweatherDailyForecast
}

type AccuweatherHeadline struct {
	Text string
}

type AccuweatherDailyForecast struct {
	Date        string
	EpochDate   int
	Temperature AccuweatherTemperature
	Day         AccuweatherDayDetail
	Night       AccuweatherDayDetail
}

type AccuweatherTemperature struct {
	Minimum AccuweatherTemperatureDetail
	Maximum AccuweatherTemperatureDetail
}

type AccuweatherTemperatureDetail struct {
	Value    float32
	Unit     string
	UnitType int
}

type AccuweatherDayDetail struct {
	Icon       int
	IconPhrase string
}
