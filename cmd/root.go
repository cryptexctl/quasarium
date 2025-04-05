package cmd

import (
	"fmt"
	"os"

	"quasarium/internal/api"
	"quasarium/internal/download"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quasarium",
	Short: "Quasarium — прошивочный качатель для Яндекса",
	Run: func(cmd *cobra.Command, args []string) {
		deviceID, _ := cmd.Flags().GetString("device-id")
		platform, _ := cmd.Flags().GetString("platform")
		version, _ := cmd.Flags().GetString("version")

		fmt.Println("Проверка обновлений...")

		result, err := api.CheckForUpdate(deviceID, platform, version)
		if err != nil {
			fmt.Println("Ошибка запроса:", err)
			os.Exit(1)
		}

		if !result.HasUpdate {
			fmt.Println("Обновлений нет")
			return
		}

		fmt.Printf("Найдена версия: %s\n", result.Version)
		fmt.Println("Скачивание...")

		path, err := download.DownloadFirmware(result.DownloadURL, result.Version)
		if err != nil {
			fmt.Println("Ошибка загрузки:", err)
			os.Exit(1)
		}

		fmt.Println("Скачано в:", path)
	},
}

func Execute() {
	rootCmd.Flags().String("device-id", "", "ID устройства (обязателен)")
	rootCmd.Flags().String("platform", "", "Платформа (например, saturn)")
	rootCmd.Flags().String("version", "", "Текущая версия (можно пустую)")
	rootCmd.MarkFlagRequired("device-id")
	rootCmd.MarkFlagRequired("platform")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Ошибка запуска:", err)
		os.Exit(1)
	}
}
