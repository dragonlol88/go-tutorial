package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	mathCommand = &cobra.Command{
		Use: "math",
		Short :`math calculation cmd`,
		Long  : `math calculation cmd`,
		Run   : func(cmd *cobra.Command, args []string) {
			fmt.Println("hello math")
		},
	}
)

func init() {
	rootCmd.AddCommand(mathCommand)
}