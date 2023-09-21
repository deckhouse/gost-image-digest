/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gost-image-digest/imagedigest"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var fixGostDigest bool

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate GOST Image Digest",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []crane.Option{}
		if insecure {
			opts = append(opts, crane.Insecure)
		}
		err := imagedigest.ValidateGostImageDigest(args[0], opts...)
		if err != nil {
			log.Error().Err(err).Msg("ValidateGostImageDigest")
			if fixGostDigest {
				fmt.Println("Fix GOST Image Digest")
				err := imagedigest.AddGostImageDigest(args[0], opts...)
				if err != nil {
					log.Fatal().Err(err).Msg("AddGostImageDigest")
				}
				fmt.Println("Added successfully")
				return
			}
			return
		}
		fmt.Println("Validate successfully")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().BoolVarP(&fixGostDigest, "fix", "", false, "Fix Gost Image Digest")
}
