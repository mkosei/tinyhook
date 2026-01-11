package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"webhook-receiver/internal/store"
)

func StartTail(store *store.MemoryStore) {
	ch := store.Subscribe()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\n👋 stopped tailing")
		store.Unsubscribe(ch)
		os.Exit(0)
	}()

	for event := range ch {
		fmt.Printf(
			"[%s] %-8s\n",
			event.CreatedAt.Format("15:04:05"),
			event.Provider,
		)
	}
}
