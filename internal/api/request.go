package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quasarium/internal/models"
)

func CheckUpdate(platform, deviceID, version string) models.FirmwareResponse {
	url := fmt.Sprintf("https://quasar.yandex.net/check_updates?device_id=%s&platform=%s&version=%s", deviceID, platform, version)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return models.FirmwareResponse{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result models.FirmwareResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return models.FirmwareResponse{}
	}
	return result
}
