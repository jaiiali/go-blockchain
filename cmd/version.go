package main

import (
	"github.com/spf13/cobra"
)

var (
	versionTag, versionCommit, versionDate string
)

func VersionCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Printf("version %s build from %s on %s\n", versionTag, versionCommit, versionDate)
		},
	}
}
