package cmd

import (
	"strings"

	"github.com/go-tutorial/tutorial/sort"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var (
	sortCommand = &cobra.Command{
		Use: "sort",
		Short :`Compare sort algorithms`,
		Long  : `Compare sort algorithms, 
					quick 
					bubble 
					heap
					gnome 
					merge 
					tree`,
		Run   : func(cmd *cobra.Command, args []string) {
			var sortItems []string

			size := viper.GetInt("sort-size")
			kind := viper.GetString("sort-kind")

			if kind == "" {
				sortItems = sort.DEFAULT
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