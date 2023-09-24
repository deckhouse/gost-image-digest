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
	Use:   "image-gost-digest",
	Short: "Add a gost image digest",
	Long:  `Add a gost image digest to an image in registry`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "ignore TLS verify")
}
