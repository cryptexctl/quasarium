package models

type FirmwareResponse struct {
	HasUpdate   bool   `json:"hasUpdate"`
	Version     string `json:"version"`
	DownloadURL string `json:"downloadUrl"`
	CRC32       uint32 `json:"crc32"`
	Critical    bool   `json:"critical"`
	FallbackURL string `json:"fallbackUrl"`
	FallbackIP  string `json:"fallbackIp"`
	UpdateID    string `json:"updateId"`
}
