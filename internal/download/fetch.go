package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"quasarium/internal/models"
)

func SaveFirmware(data models.FirmwareResponse) {
	dir := filepath.Join("firmwares", data.Version)
	os.MkdirAll(dir, 0755)

	zipPath := filepath.Join(dir, data.Version + ".zip")
	out, err := os.Create(zipPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	resp, err := http.Get(data.DownloadUrl)
	if err != nil {
		fmt.Println("Download failed:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("[+] Firmware saved to:", zipPath)

	// Optional: write metadata
	meta := filepath.Join(dir, "README.txt")
	os.WriteFile(meta, []byte(fmt.Sprintf(
		"Platform version: %s\nCRC32: %d\nUpdateID: %s\nURL: %s\n",
		data.Version, data.Crc32, data.UpdateId, data.DownloadUrl,
	)), 0644)
}
