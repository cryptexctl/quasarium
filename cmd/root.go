package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"quasarium/internal/api"
	"quasarium/internal/download"
)

var (
	deviceID string
	platform string
	version  string
)

var rootCmd = &cobra.Command{
	Use:   "quasarium",
	Short: "Firmware fetcher from Yandex Quasar",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[*] Fetching update...")
		data := api.CheckUpdate(platform, deviceID, version)
		if data.HasUpdate {
			download.SaveFirmware(data)
		} else {
			fmt.Println("[!] No update available.")
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&deviceID, "device-id", "443078968408042905d0", "Device ID")
	rootCmd.PersistentFlags().StringVar(&platform, "platform", "", "Platform name (e.g. saturn)")
	rootCmd.PersistentFlags().StringVar(&version, "version", "", "Optional: specify version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
