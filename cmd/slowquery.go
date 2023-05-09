package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// slowqueryCmd represents the slowquery command
var slowqueryCmd = &cobra.Command{
	Use:   "slowquery",
	Short: "Generate command to aggregate mysql slow queries",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "/var/log/httpd/access_log" {
			command := "ionice -c2 -n7 nice -n19 tail -n1000 /var/log/mysql/mysql-slow.log | awk '/^# Time/{ d=$3; t=$4; } (/^# Query/ && d){ c=$2; q=$3; print d,t,c,q; }' | cut -d : -f 1,2 | sort -nk2 | uniq -c"
			clipboard.WriteAll(command)
			fmt.Println(command + "\n\nCopy Complete!")
		} else {
			command := "ionice -c2 -n7 nice -n19 tail -n1000 " + path + " | awk '/^# Time/{ d=$3; t=$4; } (/^# Query/ && d){ c=$2; q=$3; print d,t,c,q; }' | cut -d : -f 1,2 | sort -nk2 | uniq -c"
			clipboard.WriteAll(command)
			fmt.Println(command + "\n\nCopy Complete!")
		}
	},
}

func init() {
	rootCmd.AddCommand(slowqueryCmd)
}
