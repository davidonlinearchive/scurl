package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func SendGetRequest(url string) error {
	c := &http.Client{Timeout: 13 * time.Second} //  Timeout client requests after 13 seconds

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "scurl/0.1")

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return nil
}

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Perform a HTTP get request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		return SendGetRequest(url)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
