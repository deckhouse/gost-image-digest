/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	debug    bool
	insecure bool
)

var rootCmd = &cobra.Command{
	Use:   "imagedigest",
	Short: "",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "ignore TLS verify")
}
