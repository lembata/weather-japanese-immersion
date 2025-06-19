Tool using https://www.weatherapi.com/ intended to be used with waybar.
It uses 2 environment variables `WEATHER_API_KEY` and `WEATHER_API_LOCATION` to output json in this format:

```json
{"text":"19.4°C (晴れ)","alt":1000,"class":"class-1000","tooltip":"はれ \n\n\n------------(Sunny)"
```

Sample of waybar config:
```json
"custom/weather": {
    "return-type": "json",
    "format" : "{text}",
    "tooltip": true,
    "interval": 900,
    "exec": "WEATHER_API_KEY=[Your api key] WEATHER_API_LOCATION=[Your Location] weather"
}
```
