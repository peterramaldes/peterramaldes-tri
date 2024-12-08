package cmd

import (
	"fmt"
	"log/slog"

	"github.com/peterramaldes/tri/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		msg := fmt.Sprintf("could not read file '%s': %s", dataFile, err)
		slog.Error(msg)
	}

	fmt.Println(items)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
