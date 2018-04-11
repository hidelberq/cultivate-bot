package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Text string `json:"text"`
}

// https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html
type YahooWeather struct {
	Feature []struct {
		Property struct {
			WeatherList struct {
				Weather []struct {
					Rainfall float32 `json:"Rainfall"`
					Type     string  `json:"Type"`
				} `json:"Weather"`
			} `json:"WeatherList"`
		} `json:"Property"`
	} `json:"Feature"`
}

func Handler() (*Response, error) {
	fmt.Println("Hello")
	lon := os.Getenv("LOCATION_LON")
	lan := os.Getenv("LOCATION_LAT")
	appId := os.Getenv("APP_ID")

	u := "https://map.yahooapis.jp/weather/V1/place" +
		"?coordinates=" + lon + "," + lan +
		"&appid=" + appId +
		"&output=json"
	resp, err := http.Get(u)
	if err != nil {
		return &Response{
			Text: "failed to get yahooapis",
		}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Text: "failed to read to yahooapis body",
		}, err
	}
	fmt.Println(string(body))

	var y YahooWeather
	if err != json.Unmarshal(body, &y) {
		return &Response{
			Text: "failed to decode to yahooapis body",
		}, err
	}

	fmt.Println(y)
	weathers := y.Feature[0].Property.WeatherList.Weather
	future := false
	for _, w := range weathers {
		if w.Type == "forecast" && w.Rainfall > 0 {
			future = true
		}
	}

	var nowRainfall float32 = 0.0
	for _, w := range weathers {
		if w.Type == "observation" && w.Rainfall > 0 {
			nowRainfall = w.Rainfall
			break
		}
	}

	return sendMessage(future, nowRainfall)
}

func sendMessage(future bool, nowRainfall float32) (*Response, error) {
	text := "現在雨は降っていません。"
	emoji := ":sunny:"
	if nowRainfall > 0 {
		text = "雨が降っています。"
		if nowRainfall >= 80 {
			text += "記録的豪雨です。"
			emoji = ":thunder_cloud_and_rain:"
		} else if nowRainfall >= 50 {
			text += "滝のような雨です。"
			emoji = ":thunder_cloud_and_rain:"
		} else if nowRainfall >= 30 {
			text += "バケツをひっくり返したような雨です。"
			emoji = ":thunder_cloud_and_rain:"
		} else if nowRainfall >= 20 {
			text += "どしゃぶりです。"
			emoji = ":rain_cloud:"
		} else if nowRainfall >= 10 {
			text += "ざーざー振りです。"
			emoji = ":rain_cloud:"
		} else if nowRainfall >= 4 {
			text += "降っています。。"
			emoji = ":rain_cloud:"
		} else if nowRainfall >= 1 {
			text += "傘があったほうが良いです。"
			emoji = ":umbrella:"
		} else if nowRainfall > 0.2 {
			text += "霧雨です。。"
			emoji = ":cloud:"
		}
	} else if future {
		text += "現在雨は降っていませんが、雨が降りそうな気配を感じます。"
		emoji = ":partly_sunny_rain:"
	}

	return &Response{
		Text: emoji + " " + text,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
