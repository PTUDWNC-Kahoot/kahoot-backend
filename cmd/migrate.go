package cmd

import (
	"github.com/spf13/cobra"

	"examples/kahootee/internal/app"
)

func migrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			app.StartMigrate()
		},
	}
}
