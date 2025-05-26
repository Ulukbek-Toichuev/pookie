/*
Pookie — simple http client.
Copyright © 2025 Toichuev Ulukbek t.ulukbek01@gmail.com

Licensed under the MIT License.
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pookie",
	Short: "simple http client.",
	Long:  "simple http client.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(GetCmd())
}
