package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go-tutorial/tutorial/sort"

)


var (
	sortCommand = &cobra.Command{
		Use: "sort",
		Short :`Calculation cmd of two number`,
		Long  : `Calculation that performs arithmetic operations of two numbers.`,
		Run   : func(cmd *cobra.Command, args []string) {

			var sortItems []string

			size := viper.GetInt("sort-size")
			kind := viper.GetString("sort-kind")

			if kind == "" {
				sortItems = sort.DEAFAULT
			} else {
				sortItems = strings.Split(kind, ",")
			}

			for _, v := range sortItems {
				s, err := sort.New(v, size)
				if err != nil {
					exitByError(err)
				}
				sort.DoSort(s)
			}
		},
	}
)


func init() {

	sortCommand.Flags().IntP("size", "", 100000, "[optional]")
	sortCommand.Flags().StringP("sort", "", "", "[optional]")

	viper.BindPFlag("sort-size", sortCommand.Flags().Lookup("size"))
	viper.BindPFlag("sort-kind", sortCommand.Flags().Lookup("sort"))
	rootCmd.AddCommand(sortCommand)
}