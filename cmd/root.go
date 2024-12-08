package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long: `Tri will help you get more done in less 
time. It's designed to be as simple as possible to 
help you accomplish your goals.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var dataFile string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("unable to detect the home directory. Please set data file using --datafile")
	}

	datafilePath := home + string(os.PathSeparator) + ".tridos.json"
	rootCmd.PersistentFlags().StringVarP(&dataFile, "datafile", "d", datafilePath, "data file to stores todos")
}
