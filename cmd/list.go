package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"text/tabwriter"

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

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, v := range items {
		pos := fmt.Sprintf("%d.", v.Position)
		fmt.Fprintln(w, pos, v.PrettyP(), v.Text)
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
