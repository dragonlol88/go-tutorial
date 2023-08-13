package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)




var (
	// rootCmd represents the base command when called without any sub-commands
	rootCmd = &cobra.Command{
		Use:   "tutorial",
		Short: ``,
		Long:  ``,
	}
)


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitByError(err)
	}
}

// panicRed raises error with text.
func exitByError(err error) {
	fmt.Println(color.RedString("[err] %s", err.Error()))
	os.Exit(1)
}


func initConfig() {

}

func init() {
	cobra.OnInitialize(initConfig)
}