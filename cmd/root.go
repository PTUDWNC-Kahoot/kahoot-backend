package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "kahootee",
		Short: "This is a short description",
	}

	rootCmd.AddCommand(serveCommand(), migrateCommand())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
