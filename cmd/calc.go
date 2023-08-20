package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-tutorial/tutorial/calc"
	"github.com/spf13/cobra"
)


var (
	calcCommand = &cobra.Command{
		Use: "calc [action `sum|mul|div|sub`] [integer to num1] [integer to num2]",
		Short :`Calculation cmd of two number`,
		Long  : `Calculation that performs arithmetic operations of two numbers.`,
		Args: func(cmd *cobra.Command, args []string) error {

			// validation the number of argument
			if len(args) != 3 {
				return fmt.Errorf("Error: `%s` is insufficient arguments", args)
			}

			// validate second argument correct type.
			_, err := strconv.Atoi(strings.TrimSpace(args[1]))
			if err != nil {
				return fmt.Errorf("Error: `%s` connot converted to integer", args[1])
			}

			// validate third argument correct type.
			_, err2 := strconv.Atoi(strings.TrimSpace(args[2]))
			if err2 != nil {
				return fmt.Errorf("Error: `%s` connot converted to integer", args[2])
			}
			return nil
		},
		Run   : func(cmd *cobra.Command, args []string) {
			// parse sting integer to integer
			num1, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			num2, _ := strconv.Atoi(strings.TrimSpace(args[2]))

			c := &calc.Calc{num1 , num2}
			result, err := c.CalCuLate(args[0])

			if err != nil {
				exitByError(err)
			}

			fmt.Println(args[0], " Result: ",  result)
		},
	}
)


func init() {
	rootCmd.AddCommand(calcCommand)
}