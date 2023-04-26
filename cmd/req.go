package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "A command can be generated to tally the number of requests.",
	Long: `A command can be generated to tally the number of requests.
	
It can generate commands from seconds to days.`,
	Run: func(cmd *cobra.Command, args []string) {
		var command string = "ionice -c2 -n7 nice -n19 egrep " + date + " " + path + " | awk '{print $4}' | cut -b 2-18 | sort | uniq -c"
		clipboard.WriteAll(command)
		fmt.Println(command + "\n\nCopy Complete!")
	},
}

func init() {
	rootCmd.AddCommand(reqCmd)
}
