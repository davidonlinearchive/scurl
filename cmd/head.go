/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func sendHeadRequest(url string) {
	c := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	resp, err := c.Do(req
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, v := range values {
			fmt.Printf("%s: %s\n", key, v)
		}
	}
}

// headCmd represents the head command
var headCmd = &cobra.Command{
	Use:   "HEAD",
	Short: "Perform a HTTP head request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("head called")
		url := args[0]
		sendHeadRequest(url)
	},
}

func init() {
	rootCmd.AddCommand(headCmd)
}
