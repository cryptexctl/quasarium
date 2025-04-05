package download

import (
	//	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"net/url"

	"github.com/schollz/progressbar/v3"
)

func safeURL(raw string) string {
	// Парсим строку в объект URL и собираем обратно, чтобы не светилась полностью
	u, _ := url.Parse(raw)
	return u.Scheme + "://" + u.Host + u.Path // без query
}

func DownloadFirmware(url, version string) (string, error) {
	dir := filepath.Join("firmwares", version)
	_ = os.MkdirAll(dir, 0755)

	dest := filepath.Join(dir, version+".zip")

	resp, err := http.Get(safeURL(url))
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
