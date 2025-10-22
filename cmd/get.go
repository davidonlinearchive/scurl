/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func SendGetRequest(url string) {
	c := &http.Client{Timeout: 30 * time.Second} /*  Timeout client request after 30 seconds */

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)

}

/*  getCmd represents the GET command */
var getCmd = &cobra.Command{
	Use:   "GET",
	Short: "Perform a HTTP get request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0] /* args[0] != os.Args[0] */
		SendGetRequest(url)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
