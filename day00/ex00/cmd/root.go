package cmd

import (
	ans "ex00/anscombe"
	rdr "ex00/reader"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

func processFlags() {
	// if no flags are given, the default scheme is set
	if !(Flags.Mean || Flags.Median || Flags.Mode || Flags.Psd || Flags.Ssd) {
		Flags.Mean = true
		Flags.Median = true
		Flags.Mode = true
		Flags.Psd = true
	}
	if Flags.Psd && Flags.Ssd {
		err_str := fmt.Errorf("the --psd and --ssd cannot coexist")
		fmt.Fprintln(os.Stderr, err_str)
		os.Exit(2)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ex00",
	Short: "APG1-Bootcamp",
	Long: `
The School 21 Go Language Bootcamp -> Day00/ex00 - Anscombe's quartet.

The numbers are given from the system standard input stream.
To stop the input, you may press Ctrl+D or any combination alike.
`,
	Run: func(cmd *cobra.Command, args []string) {
		processFlags()

		nums, err := rdr.ReadNumbers()
		if len(nums) == 0 {
			fmt.Println("No numbers were given.")
			os.Exit(0)
		}
		fmt.Println() // to separate the outputs of reading and processing sections

		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadNumbersError")
			os.Exit(3)
		}

		if Flags.Mean {
			mean, err := ans.GetMean(nums)
			if err != nil {
				fmt.Fprintln(os.Stderr, "GetMean Error")
				os.Exit(4)
			}
			fmt.Printf("Mean: %.2f\n", mean)
		}
		if Flags.Median {
			median, err := ans.GetMedian(nums)
			if err != nil {
				fmt.Fprintln(os.Stderr, "GetMedian Error")
				os.Exit(5)
			}
			fmt.Printf("Median: %.2f\n", median)
		}
		if Flags.Mode {
			mode, err := ans.GetMode(nums)
			if err != nil {
				fmt.Fprintln(os.Stderr, "GetMode Error")
				os.Exit(6)
			}
			fmt.Printf("Mode: %.2f\n", mode)
		}
		if Flags.Psd {
			psd, err := ans.GetPopulationStandardDeviation(nums)
			if err != nil {
				fmt.Fprintln(os.Stderr, "GetPopulationStandardDeviation Error")
				os.Exit(7)
			}
			fmt.Printf("SD: %.2f\n", psd)
		}
		if Flags.Ssd {
			ssd, err := ans.GetSampleStandardDeviation(nums)
			if err != nil {
				fmt.Fprintln(os.Stderr, "GetSampleStandardDeviation Error")
				os.Exit(8)
			}
			fmt.Printf("SD: %.2f\n", ssd)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&Flags.Mean, "mean", false, "Compute the Mean")
	rootCmd.PersistentFlags().BoolVar(&Flags.Median, "median", false, "Compute the Median")
	rootCmd.PersistentFlags().BoolVar(&Flags.Mode, "mode", false, "Compute the Mode")
	rootCmd.PersistentFlags().BoolVar(&Flags.Psd, "psd", false, "Compute the Population Standard Deviation")
	rootCmd.PersistentFlags().BoolVar(&Flags.Ssd, "ssd", false, "Compute the Sample Standard Deviation")
}
