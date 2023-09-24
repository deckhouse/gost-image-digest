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

var fixGostDigest bool

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validating the image digest in the image metadata calculated according to the GOST standard Streebog (GOST R 34.11-2012)",
	Long:  `Validating the image digest in the image metadata calculated according to the GOST standard Streebog (GOST R 34.11-2012)`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []crane.Option{}
		if insecure {
			opts = append(opts, crane.Insecure)
		}
		err := imagedigest.ValidateGostImageDigest(args[0], opts...)
		if err != nil {
			log.Error().Err(err).Msg("ValidateGostImageDigest")
			if fixGostDigest {
				log.Info().Msg("Fix GOST Image Digest")
				err := imagedigest.AddGostImageDigest(args[0], opts...)
				if err != nil {
					log.Fatal().Err(err).Msg("AddGostImageDigest")
				}
				log.Info().Msg("Added successfully")
				return
			}
			return
		}
		log.Info().Msg("Validate successfully")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().BoolVarP(&fixGostDigest, "fix", "", false, "Fix Gost Image Digest if it is incorrect")
}
