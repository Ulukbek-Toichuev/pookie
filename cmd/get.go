/*
Pookie — simple http client.
Copyright © 2025 Toichuev Ulukbek t.ulukbek01@gmail.com

Licensed under the MIT License.
*/

package cmd

import (
	"github.com/Ulukbek-Toichuev/pookie/internal"
	"github.com/spf13/cobra"
)

func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "http GET method",
		Args:    cobra.ExactArgs(1),
		RunE:    internal.Get,
		Example: "pookie get http://localhost:88081/api/v1/users -H \"Authorization: Bearer Token\" -T 10",
	}

	cmd.Flags().IntP("timeout", "T", 5, "http request timeout, duration must be in sec")
	cmd.Flags().StringToStringP("headers", "H", map[string]string{}, "http headers")
	cmd.Flags().StringToStringP("params", "P", map[string]string{}, "http params")

	return cmd
}
