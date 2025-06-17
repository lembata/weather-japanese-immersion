package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type APIResponse struct {
	Current APICurrent `json:"current"`
}

type APICurrent struct {
	Temperature float64      `json:"temp_c"`
	Condition   APICondition `json:"condition"`
}

type APICondition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}

type Output struct {
	Text  string `json:"text"`
	Alt   int    `json:"alt"`
	Class string `json:"class"`
}

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=Sofia&key=%s&lang=ja", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// fmt.Println("Raw response:")
	// fmt.Println(string(body))

	// Parse the response into our struct
	var weatherData APIResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// fmt.Printf("\nParsed weather data:\n")
	// fmt.Printf("Temperature: %.1f°C\n", weatherData.Current.Temperature)
	// fmt.Printf("Condition: %s\n", weatherData.Current.Condition.Text)

	output := Output{
		Text:  fmt.Sprintf("%3.1f°C (%s)", weatherData.Current.Temperature, weatherData.Current.Condition.Text),
		Alt:   weatherData.Current.Condition.Code,
		Class: "class-" + strconv.Itoa(weatherData.Current.Condition.Code),
	}

	s, err := json.Marshal(output)

	if err != nil {
		output.Text = "Error"
		output.Alt = 0
		output.Class = "error"

		s, _ = json.Marshal(output)
	}

	fmt.Println(string(s))
}

type CodeTranslation struct {
	Kana    string `json:"kana"`
	English string `json:"english"`
}

var asd = map[int]CodeTranslation{
	1000: {English: "Sunny", Kana: "はれ"},
	1003: {English: "Partly cloudy", Kana: "ところによりくもり"},
	1006: {English: "Cloudy", Kana: "くもり"},
	1009: {English: "Overcast", Kana: "ほんくもり"},
	1030: {English: "Mist", Kana: "もや"},
}

/*
code,day,night,icon
1000,Sunny,Clear,113
1003,"Partly cloudy","Partly cloudy",116
1006,Cloudy,Cloudy,119
1009,Overcast,Overcast,122
1030,Mist,Mist,143
1063,"Patchy rain possible","Patchy rain possible",176
1066,"Patchy snow possible","Patchy snow possible",179
1069,"Patchy sleet possible","Patchy sleet possible",182
1072,"Patchy freezing drizzle possible","Patchy freezing drizzle possible",185
1087,"Thundery outbreaks possible","Thundery outbreaks possible",200
1114,"Blowing snow","Blowing snow",227
1117,Blizzard,Blizzard,230
1135,Fog,Fog,248
1147,"Freezing fog","Freezing fog",260
1150,"Patchy light drizzle","Patchy light drizzle",263
1153,"Light drizzle","Light drizzle",266
1168,"Freezing drizzle","Freezing drizzle",281
1171,"Heavy freezing drizzle","Heavy freezing drizzle",284
1180,"Patchy light rain","Patchy light rain",293
1183,"Light rain","Light rain",296
1186,"Moderate rain at times","Moderate rain at times",299
1189,"Moderate rain","Moderate rain",302
1192,"Heavy rain at times","Heavy rain at times",305
1195,"Heavy rain","Heavy rain",308
1198,"Light freezing rain","Light freezing rain",311
1201,"Moderate or heavy freezing rain","Moderate or heavy freezing rain",314
1204,"Light sleet","Light sleet",317
1207,"Moderate or heavy sleet","Moderate or heavy sleet",320
1210,"Patchy light snow","Patchy light snow",323
1213,"Light snow","Light snow",326
1216,"Patchy moderate snow","Patchy moderate snow",329
1219,"Moderate snow","Moderate snow",332
1222,"Patchy heavy snow","Patchy heavy snow",335
1225,"Heavy snow","Heavy snow",338
1237,"Ice pellets","Ice pellets",350
1240,"Light rain shower","Light rain shower",353
1243,"Moderate or heavy rain shower","Moderate or heavy rain shower",356
1246,"Torrential rain shower","Torrential rain shower",359
1249,"Light sleet showers","Light sleet showers",362
1252,"Moderate or heavy sleet showers","Moderate or heavy sleet showers",365
1255,"Light snow showers","Light snow showers",368
1258,"Moderate or heavy snow showers","Moderate or heavy snow showers",371
1261,"Light showers of ice pellets","Light showers of ice pellets",374
1264,"Moderate or heavy showers of ice pellets","Moderate or heavy showers of ice pellets",377
1273,"Patchy light rain with thunder","Patchy light rain with thunder",386
1276,"Moderate or heavy rain with thunder","Moderate or heavy rain with thunder",389
1279,"Patchy light snow with thunder","Patchy light snow with thunder",392
1282,"Moderate or heavy snow with thunder","Moderate or heavy snow with thunder",395
*/
