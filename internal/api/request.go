package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"quasarium/internal/models"
)

func CheckForUpdate(deviceID, platform, version string) (*models.FirmwareResponse, error) {
	url := fmt.Sprintf("https://quasar.yandex.net/check_updates?device_id=%s&platform=%s&version=%s", deviceID, platform, version)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.FirmwareResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
