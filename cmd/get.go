/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "GET",
	Short: "Perform a HTTP get request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get command called") // remove line later; shows GET command called
		url := args[0]                    // args[0] not os.Args[0]
		SendGetRequest(url)
	},
}

func SendGetRequest(url string) {
	c := &http.Client{Timeout: 30 * time.Second} /*  Timeout client request after 30 seconds */

	res, err := c.Get(url)
	if err != nil {
		log.Fatalln("Failed to get domain ", err)
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
