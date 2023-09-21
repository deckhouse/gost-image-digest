/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"gost-image-digest/imagedigest"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate a gost image digest",
	Long:  `Calculate a gost image digest`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []crane.Option{}
		if insecure {
			opts = append(opts, crane.Insecure)
		}
		gostImageDigest, err := imagedigest.CalculateGostImageDigest(args[0], opts...)
		if err != nil {
			log.Fatal().Err(err).Msg("CalculateGostImageDigest")
		}
		fmt.Printf("GOST Image Digest: %s\n", hex.EncodeToString(gostImageDigest))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)
}
