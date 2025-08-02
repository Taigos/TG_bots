package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func GetMoscowWeather(apiKey string) (string, error) {
	baseURL := "http://api.openweathermap.org/data/2.5/weather"

	params := url.Values{}
	params.Add("q", "Moscow")
	params.Add("appid", apiKey)
	params.Add("units", "metric")
	params.Add("lang", "ru")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return "", fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	if len(weather.Weather) == 0 {
		return "", fmt.Errorf("данные о погоде не найдены")
	}

	message := fmt.Sprintf("🌤 Погода в Москве:\n"+
		"🌡 Температура: %.1f°C\n"+
		"💧 Влажность: %d%%\n"+
		"💨 Ветер: %.1f м/с\n"+
		"📝 Описание: %s",
		weather.Main.Temp,
		weather.Main.Humidity,
		weather.Wind.Speed,
		weather.Weather[0].Description)

	return message, nil
}
