package api

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"quasarium/internal/models"
)

func getQuasarURL() string {
	encoded := "cXVhc2FyLnlhbmRleC5uZXQ=" // base64("quasar.yandex.net")
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decoded)
}

func CheckForUpdate(deviceID, platform, version string) (*models.FirmwareResponse, error) {
	// формируем URL запроса (не выводим домен!)
	path := fmt.Sprintf("/check_updates?device_id=%s&platform=%s&version=%s", deviceID, platform, version)

	host := getQuasarURL()

	// резолв через 8.8.8.8
	ip, err := resolveViaGoogle(host)
	if err != nil {
		return nil, fmt.Errorf("dns resolve error: %w", err)
	}

	client := stealthClient(ip, host)

	// собираем полную URL
	url := "https://" + host + path

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data models.FirmwareResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// 👇 Сюда magic
func stealthClient(ip, sni string) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, _ string) (net.Conn, error) {
				dialer := &net.Dialer{}
				return dialer.DialContext(ctx, network, ip+":443")
			},
			TLSClientConfig: &tls.Config{
				ServerName: sni, // это SNI для успешного TLS
			},
		},
		Timeout: 10 * time.Second,
	}
}

func resolveViaGoogle(host string) (string, error) {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(_ context.Context, _, _ string) (net.Conn, error) {
			d := net.Dialer{}
			return d.Dial("udp", "8.8.8.8:53")
		},
	}
	ips, err := resolver.LookupHost(context.Background(), host)
	if err != nil {
		return "", err
	}
	return ips[0], nil
}
