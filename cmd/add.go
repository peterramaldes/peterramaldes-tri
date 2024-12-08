package cmd

import (
	"log/slog"

	"github.com/peterramaldes/tri/todo"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	items := []todo.Item{}
	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}

	err := todo.SaveItems(dataFile, items)
	if err != nil {
		slog.Error(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
