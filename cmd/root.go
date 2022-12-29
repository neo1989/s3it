package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "s3it",
	Short: "自用图床",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello there")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
