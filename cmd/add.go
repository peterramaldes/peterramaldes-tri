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
		i := todo.Item{Text: x}
		i.SetPriority(priority)
		items = append(items, i)
	}

	err := todo.SaveItems(dataFile, items)
	if err != nil {
		slog.Error(err.Error())
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
}
