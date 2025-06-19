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
	Text    string `json:"text"`
	Alt     int    `json:"alt"`
	Class   string `json:"class"`
	Tooltip string `json:"tooltip"`
}

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	location := os.Getenv("WEATHER_API_LOCATION")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=%s&key=%s&lang=ja", location, apiKey)

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

	var weatherData APIResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	output := Output{
		Text:    fmt.Sprintf("%3.1f°C (%s)", weatherData.Current.Temperature, weatherData.Current.Condition.Text),
		Alt:     weatherData.Current.Condition.Code,
		Class:   "class-" + strconv.Itoa(weatherData.Current.Condition.Code),
		Tooltip: getTooltip(weatherData.Current.Condition.Text),
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

func getTooltip(value string) string {
	a, ok := kanjiFuriganaMap[value]

	if !ok {
		return "Not found for value" + value
	}

	return fmt.Sprintf("%s\n\n\n---------------\n(%s)", a.Kana, a.English)
}

type CodeTranslation struct {
	Kana    string `json:"kana"`
	English string `json:"english"`
}

var kanjiFuriganaMap = map[string]CodeTranslation{
	"晴れ":         {Kana: "はれ", English: "Sunny"},
	"快晴":         {Kana: "かいせい", English: "Clear"},
	"所により曇り":     {Kana: "ところによりくもり", English: "Partly cloudy"},
	"曇り":         {Kana: "くもり", English: "Cloudy"},
	"本曇り":        {Kana: "ほんくもり", English: "Overcast"},
	"もや":         {Kana: "もや", English: "Mist"},
	"近くで所により雨":   {Kana: "ちかくでところによりあめ", English: "Patchy rain possible"},
	"近くで所により雪":   {Kana: "ちかくでところによりゆき", English: "Patchy snow possible"},
	"近くで所によりみぞれ": {Kana: "ちかくでところによりみぞれ", English: "Patchy sleet possible"},
	"近くで所により着氷性の霧雨": {Kana: "ちかくでところによりちゃくひょうせいのきりさめ", English: "Patchy freezing drizzle possible"},
	"近くで雷の発生":       {Kana: "ちかくでかみなりのはっせい", English: "Thundery outbreaks possible"},
	"吹雪":            {Kana: "ふぶき", English: "Blowing snow"},
	"猛吹雪":           {Kana: "もうふぶき", English: "Blizzard"},
	"霧":             {Kana: "きり", English: "Fog"},
	"着氷性の霧":         {Kana: "ちゃくひょうせいのきり", English: "Freezing fog"},
	"所により霧雨":        {Kana: "ところによりきりさめ", English: "Patchy fog"},
	"霧雨":            {Kana: "きりさめ", English: "Freezing drizzle"},
	"着氷性の霧雨":        {Kana: "ちゃくひょうせいのきりさめ", English: "Heavy freezing drizzle"},
	"強い着氷性の霧雨":      {Kana: "つよいちゃくひょうせいのきりさめ", English: "Heavy freezing drizzle"},
	"所により弱い雨":       {Kana: "ところによりよわいあめ", English: "Light freezing rain"},
	"弱い雨":           {Kana: "よわいあめ", English: "Light rain"},
	"時々穏やかな雨":       {Kana: "ときどきおだやかなあめ", English: "Moderate rain at times"},
	"穏やかな雨":         {Kana: "おだやかなあめ", English: "Moderate rain"},
	"時々大雨":          {Kana: "ときどきおおあめ", English: "Heavy rain at times"},
	"大雨":            {Kana: "おおあめ", English: "Heavy rain"},
	"着氷性の弱い雨":       {Kana: "ちゃくひょうせいのよわいあめ", English: "Light freezing rain"},
	"着氷性の穏やかな雨または大雨": {Kana: "ちゃくひょうせいのおだやかなあめまたはおおあめ", English: "Moderate or heavy freezing rain"},
	"軽いみぞれ":        {Kana: "かるいみぞれ", English: "Light sleet"},
	"穏やかなまたは強いみぞれ": {Kana: "おだやかなまたはつよいみぞれ", English: "Moderate or heavy sleet"},
	"所により小雪":       {Kana: "ところによりこゆき", English: "Patchy light snow"},
	"小雪":           {Kana: "こゆき", English: "Light snow"},
	"所により穏やかな雪":    {Kana: "ところによりおだやかなゆき", English: "Patchy moderate snow"},
	"穏やかな雪":        {Kana: "おだやかなゆき", English: "Moderate snow"},
	"所により大雪":       {Kana: "ところによりおおゆき", English: "Patchy heavy snow"},
	"大雪":           {Kana: "おおゆき", English: "Heavy snow"},
	"凍雨":           {Kana: "とうう", English: "Freezing rain"},
	"軽いにわか雨":       {Kana: "かるいにわかあめ", English: "Light rain"},
	"穏やかなまたは強いにわか雨": {Kana: "おだやかなまたはつよいにわかあめ", English: "Moderate or heavy rain"},
	"急な豪雨":           {Kana: "きゅうなごうう", English: "Heavy rain"},
	"急な軽いみぞれ":        {Kana: "きゅうなかるいみぞれ", English: "Light sleet"},
	"穏やかなまたは強い急なみぞれ": {Kana: "おだやかなまたはつよいきゅうなみぞれ", English: "Moderate or heavy sleet"},
	"急な軽い雪":          {Kana: "きゅうなかるいゆき", English: "Light snow"},
	"穏やかなまたは強い急な雪":   {Kana: "おだやかなまたはつよいきゅうなゆき", English: "Moderate or heavy snow"},
	"軽い急な凍雨":         {Kana: "かるいきゅうなとうう", English: "Light freezing rain"},
	"穏やかなまたは強い急な凍雨":  {Kana: "おだやかなまたはつよいきゅうなとうう", English: "Moderate or heavy freezing rain"},
	"所により雷を伴う弱い雨":    {Kana: "ところによりかみなりをともなうよわいあめ", English: "Patchy light rain with thunder"},
	"雷を伴う穏やかなまたは強い雨": {Kana: "かみなりをともなうおだやかなまたはつよいあめ", English: "Moderate or heavy rain with thunder"},
	"雷を伴う軽い雪":        {Kana: "かみなりをともなうかるいゆき", English: "Light snow with thunder"},
	"雷を伴う穏やかなまたは強い雪": {Kana: "かみなりをともなうおだやかなまたはつよいゆき", English: "Moderate or heavy snow with thunder"},
}
