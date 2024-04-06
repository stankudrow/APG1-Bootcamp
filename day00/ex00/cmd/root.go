package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// capitalised fields are exportable
type CliFlags struct {
	Mean   bool
	Median bool
	Mode   bool
	Psd    bool
	Ssd    bool
}

var Flags = CliFlags{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ex00",
	Short: "APG1-Bootcamp",
	Long:  `The School 21 Go Language Bootcamp -> Day00/ex00`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// if no flags are given, the default scheme is set
	if !(Flags.Mean || Flags.Median || Flags.Mode || Flags.Psd || Flags.Ssd) {
		Flags.Mean = true
		Flags.Median = true
		Flags.Mode = true
		Flags.Psd = true
	}
	if Flags.Psd && Flags.Ssd {
		err_str := fmt.Errorf("the --psd or --ssd flag can be specified, not both")
		fmt.Fprintln(os.Stderr, err_str)
		os.Exit(2)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&Flags.Mean, "mean", false, "Compute the Mean")
	rootCmd.PersistentFlags().BoolVar(&Flags.Median, "median", false, "Compute the Median")
	rootCmd.PersistentFlags().BoolVar(&Flags.Mode, "mode", false, "Compute the Mode")
	rootCmd.PersistentFlags().BoolVar(&Flags.Psd, "psd", false, "Compute the Population Standard Deviation")
	rootCmd.PersistentFlags().BoolVar(&Flags.Ssd, "ssd", false, "Compute the Sample Standard Deviation")
}
