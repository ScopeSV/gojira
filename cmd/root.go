package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "yo",
}

func Execute() {
	rootCmd.PersistentFlags().StringP("name", "n", "stranger", "Name of student")

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
