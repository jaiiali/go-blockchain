package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}

	cmd.AddCommand(
		VersionCMD(),
	)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error while Execute(), trace: %v", err)
	}
}
