package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/spf13/cobra"
)

func SendPostRequest(url string, data string, HeaderMap map[string]string) error {
	c := &http.Client{Timeout: 13 * time.Second}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "scurl/0.1")
	for key, value := range HeaderMap {
		req.Header.Set(key, value)
	}

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	fmt.Println("HTTP STATUS: ", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	switch {
	// Pretty-prints JSON if possible, else print raw body
	case strings.Contains(contentType, "application/json"):
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, body, "", "\t"); err == nil {
			fmt.Println(prettyJSON.String())
		} else {
			fmt.Println(string(body))
		}

	// Pretty-prints XML if possible, else print raw body
	case strings.Contains(contentType, "application/xml"), strings.Contains(contentType, "text/xml"):
		prettyXml := xmlfmt.FormatXML(string(body), "", "\t", false)
		if prettyXml != "" {
			fmt.Println(prettyXml)
		} else {
			fmt.Println(string(body))
		}

	default:
		fmt.Println(string(body))
	}

	return nil
}

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Perform a HTTP POST request",
	Long: `Send a HTTP POST request to the specified URL.

Examples:
  scurl post https://example.com/post -d '{"key":"value"}' -H "Content-Type: application/json"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		data, _ := cmd.Flags().GetString("data")
		headers, _ := cmd.Flags().GetStringArray("header")

		headerMap := make(map[string]string)
		for _, h := range headers {
			key, value, ok := strings.Cut(h, ":")
			if ok {
				headerMap[strings.TrimSpace(key)] = strings.TrimSpace(value)
				continue
			}
		}

		return SendPostRequest(url, data, headerMap)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringP("data", "d", "", "HTTP POST data")
	postCmd.Flags().StringArrayP("header", "H", []string{}, "set request header")
}
