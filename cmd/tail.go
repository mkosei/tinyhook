package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Show incoming webhooks in real time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("❌ `tail` must be run together with server")
		fmt.Println("")
		fmt.Println("👉 Use:")
		fmt.Println("   tinyhook serve --tail")
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
}
