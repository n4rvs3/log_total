/*
Copyright © 2023 n4rvs3
*/
package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

var date string
var path string

var NowDate string = time.Now().Format("2/Jan/2006:15")
var DefaultPath string = "/var/log/httpd/access_log"

var rootCmd = &cobra.Command{
	Use:   "log_total",
	Short: "Aggregate commands can be generated.",
	Long: `Aggregate commands can be generated.
	
log_total is a log aggregation command tool that supports Apache and Nginx combined format logs.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&date, "date", "d", NowDate, "You can set the date and time in dd/mm/yy format.")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", DefaultPath, "You can set the access log path.")
}
