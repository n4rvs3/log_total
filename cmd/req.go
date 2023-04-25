/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "A command can be generated to tally the number of requests.",
	Long: `A command can be generated to tally the number of requests.
	
It can generate commands from seconds to days.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grep " + date + " /var/log/httpd/access_log | awk '{print $4}' | cut -b 2-12 | sort | uniq -c")
	},
}

func init() {
	rootCmd.AddCommand(reqCmd)
}
