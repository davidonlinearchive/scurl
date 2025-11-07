package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func SendHeadRequest(url string) error {
	c := &http.Client{Timeout: 13 * time.Second}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("%v %v\n", resp.Proto, resp.Status)
	// resp.Header is a map[string][]string where each key is a header name
	// and each value is a list of header values. This loop prints them all.
	for key, values := range resp.Header {
		for _, v := range values {
			// \033[1m starts bold, \033[0m resets formatting
			fmt.Printf("\033[1m%s\033[0m: %s\n", key, v)
		}
	}

	return nil
}

var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Perform a HTTP head request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		SendHeadRequest(url)
	},
}

func init() {
	rootCmd.AddCommand(headCmd)
}
