package download

import (
	//	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

func DownloadFirmware(url, version string) (string, error) {
	dir := filepath.Join("firmwares", version)
	_ = os.MkdirAll(dir, 0755)

	dest := filepath.Join(dir, version+".zip")

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Скачиваем...",
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return "", err
	}

	return dest, nil
}
