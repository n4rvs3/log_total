package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "A command can be generated to tally the number of requests by IP.",
	Run: func(cmd *cobra.Command, args []string) {
		var command string = `ionice -c2 -n7 nice -n19 grep ` + date + " " + path + ` | cut -d " " -f 1 | sort | uniq -c`
		clipboard.WriteAll(command)
		fmt.Println(command + "\n\nCopy Complete!")
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
