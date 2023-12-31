/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/deckhouse/gost-image-digest/pkg/imagedigest"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Calculating and adding the image digest to the image metadata according to the GOST standard Streebog (GOST R 34.11-2012)",
	Long:  `Calculating and adding the image digest to the image metadata according to the GOST standard Streebog (GOST R 34.11-2012)`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []crane.Option{}
		if insecure {
			opts = append(opts, crane.Insecure)
		}
		err := imagedigest.AddGostImageDigest(args[0], opts...)
		if err != nil {
			log.Fatal().Err(err).Msg("AddGostImageDigest")
		}
		log.Info().Msg("Added successfully")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
