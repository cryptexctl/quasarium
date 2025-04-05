package models

type FirmwareResponse struct {
	HasUpdate  bool   `json:"hasUpdate"`
	Version    string `json:"version"`
	DownloadUrl string `json:"downloadUrl"`
	Crc32      uint32 `json:"crc32"`
	UpdateId   string `json:"updateId"`
}
