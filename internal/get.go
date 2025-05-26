/*
Pookie — simple http client.
Copyright © 2025 Toichuev Ulukbek t.ulukbek01@gmail.com

Licensed under the MIT License.
*/

package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ulukbek-Toichuev/pookie/pkg"
	"github.com/spf13/cobra"
)

func Get(cmd *cobra.Command, args []string) error {
	url := args[0]
	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return err
	}

	headers, err := cmd.Flags().GetStringToString("headers")
	if err != nil {
		return err
	}

	httpClient := pkg.NewHttpWrapper(timeout)
	resp, _, err := httpClient.HttpGetRaw(url, headers, map[string]string{})
	if err != nil {
		return err
	}
	printResult(resp)
	return nil
}

func printResult(resp *http.Response) {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("Url: %s\n", resp.Request.URL))
	result.WriteString(fmt.Sprintf("Method: %s\n", resp.Request.Method))
	result.WriteString("--------------------------------\n")

	result.WriteString(fmt.Sprintf("Status: %s Code: %d\n", resp.Status, resp.StatusCode))

	fmt.Println(result.String())
}
