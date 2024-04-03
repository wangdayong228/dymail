/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/wangdayong228/test-go-email/core"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-go-email",
	Short: "send email",
	Long:  `send email`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		err := core.SendToMail(host, port, to, user, password, subject, body)
		if err != nil {
			fmt.Println("Send mail error!")
			fmt.Println(err)
		} else {
			fmt.Println("Send mail success!")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	user     string
	password string
	host     string
	port     int
	to       string
	subject  string
	body     string
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-go-email.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&user, "user", "u", "", "file merkle root")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "file merkle root")
	rootCmd.Flags().StringVarP(&host, "host", "", "", "file merkle root")
	rootCmd.Flags().IntVarP(&port, "port", "", 465, "file merkle root")
	rootCmd.Flags().StringVarP(&to, "to", "t", "", "file merkle root")
	rootCmd.Flags().StringVarP(&subject, "subject", "s", "", "file merkle root")
	rootCmd.Flags().StringVarP(&body, "body", "b", "", "file merkle root")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("password")
	rootCmd.MarkFlagRequired("host")
	rootCmd.MarkFlagRequired("port")
	rootCmd.MarkFlagRequired("to")
	rootCmd.MarkFlagRequired("subject")
	rootCmd.MarkFlagRequired("body")
}
