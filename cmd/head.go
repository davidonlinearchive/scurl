package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func SendHeadRequest(url string) error {
	c := &http.Client{Timeout: 13 * time.Second}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "scurl/0.1")

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
			fmt.Printf("\033[1m%s\033[0m: %s\n", key, v) // \033[1m starts bold, \033[0m resets formatting
		}
	}

	return nil
}

var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Perform a HTTP head request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		return SendHeadRequest(url)
	},
}

func init() {
	rootCmd.AddCommand(headCmd)
}
