package cmd

import (
	"github.com/spf13/cobra"

	"webhook-receiver/internal/cli"
	"webhook-receiver/internal/server"
	"webhook-receiver/internal/store"
)

var (
	addr     string
	withTail bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start webhook receiver server",
	Run: func(cmd *cobra.Command, args []string) {
		// ① 共有する MemoryStore を作る
		st := store.NewMemoryStore()

		// ② tail を同一プロセスで起動
		if withTail {
			go cli.StartTail(st)
		}

		// ③ server に store を渡す
		server.Start(addr, st)
	},
}

func init() {
	serveCmd.Flags().StringVarP(
		&addr,
		"addr",
		"a",
		":8080",
		"listen address",
	)

	serveCmd.Flags().BoolVar(
		&withTail,
		"tail",
		false,
		"show incoming webhooks in real time",
	)

	rootCmd.AddCommand(serveCmd)
}
