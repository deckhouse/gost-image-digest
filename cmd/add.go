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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Calculate and add GOST Image Digest",
	Long:  `Calculate and add GOST Image Digest`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := []crane.Option{}
		if insecure {
			opts = append(opts, crane.Insecure)
		}
		err := imagedigest.AddGostImageDigest(args[0], opts...)
		if err != nil {
			log.Fatal().Err(err).Msg("AddGostImageDigest")
		}
		fmt.Println("Added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
