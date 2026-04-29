package main

import (
	"log"
	"os"

	"github.com/henrygd/beszel/beszel/internal/hub"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "beszel",
	Short: "Beszel - lightweight server monitoring hub",
	Long: `Beszel is a lightweight server monitoring platform that includes
a hub for managing and viewing system metrics from connected agents.`,
}

var hubCmd = &cobra.Command{
	Use:   "hub",
	Short: "Start the Beszel hub server",
	Long:  `Start the Beszel hub server which provides the web UI and API for monitoring connected agents.`,
	Run: func(cmd *cobra.Command, args []string) {
		h := hub.NewHub()
		if err := h.Start(); err != nil {
			log.Fatal("Failed to start hub: ", err)
		}
	},
}

func init() {
	// Hub flags
	hubCmd.Flags().StringP("addr", "a", "0.0.0.0:8090", "Address to listen on (host:port)")
	hubCmd.Flags().StringP("data-dir", "d", "./beszel_data", "Directory to store data")

	rootCmd.AddCommand(hubCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
