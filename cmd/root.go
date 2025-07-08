package cmd

import (
	"fmt"
	"os"

	"quasarium/internal/api"
	"quasarium/internal/download"

	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "quasarium",
	Short: "Quasarium — утилита для скачивания прошивок для устройств от компании "Я"",
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("quasarium-version")
		if showVersion {
			fmt.Println("Quasarium version:", Version)
			return
		}

		deviceID, _ := cmd.Flags().GetString("device-id")
		platform, _ := cmd.Flags().GetString("platform")
		version, _ := cmd.Flags().GetString("version")

		fmt.Println("[FETCH] Проверка обновлений...")

		result, err := api.CheckForUpdate(deviceID, platform, version)
		if err != nil {
			fmt.Println("[ERR] Ошибка запроса:", err)
			os.Exit(1)
		}

		if !result.HasUpdate {
			fmt.Println("[OK] Обновлений нет")
			return
		}

		fmt.Printf("[UPD] Найдена версия: %s\n", result.Version)
		fmt.Println("[=] Скачивание...")

		path, err := download.DownloadFirmware(result.DownloadURL, result.Version)
		if err != nil {
			fmt.Println("[ERR] Ошибка загрузки:", err)
			os.Exit(1)
		}

		fmt.Println("[ok] Скачано в:", path)
	},
}

func Execute() {
	rootCmd.Flags().String("device-id", "", "ID устройства (обязателен)")
	rootCmd.Flags().String("platform", "", "Платформа (например, saturn)")
	rootCmd.Flags().String("version", "", "Текущая версия (можно пустую)")
	rootCmd.Flags().Bool("quasarium-version", false, "Показать версию quasarium")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Ошибка запуска:", err)
		os.Exit(1)
	}
}
