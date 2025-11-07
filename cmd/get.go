package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func SendGetRequest(url string) error {
	c := &http.Client{Timeout: 13 * time.Second} //  Timeout client requests after 13 seconds

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

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
	Use:   "get",
	Short: "Perform a HTTP get request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		SendGetRequest(url)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
