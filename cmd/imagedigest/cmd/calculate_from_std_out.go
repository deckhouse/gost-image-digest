/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/deckhouse/gost-image-digest/pkg/calculator"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateFromStdInCmd = &cobra.Command{
	Use:   "calculate-from-file",
	Short: "Calculating the file digest according to the GOST standard Streebog (GOST R 34.11-2012). For stdin use '-'",
	Long:  `Calculating the file digest according to the GOST standard Streebog (GOST R 34.11-2012)`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		reader := os.Stdin
		if filename != "-" {
			file, err := os.Open(filename)
			if err != nil {
				log.Err(err)
				os.Exit(1)
			}
			defer file.Close()
			reader = file
		}
		sum, err := calculator.Calculate(reader)
		if err != nil {
			log.Err(err)
			os.Exit(2)
		}
		fmt.Println(hex.EncodeToString(sum))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(calculateFromStdInCmd)
}
