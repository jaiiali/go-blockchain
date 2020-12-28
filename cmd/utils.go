package main

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

const exitCode = 1

type errorCLI struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Error      string `json:"error,omitempty"`
}

func printCli(cmd *cobra.Command, msg interface{}) {
	result, err := json.MarshalIndent(msg, "", "\t")
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(exitCode)
	}

	cmd.Println(string(result))
}
