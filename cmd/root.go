package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/RickyJian/chiconverter/utils"
	"github.com/RickyJian/gocc"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	// source defines persistence flag `source`
	source = "source"
	// source defines persistence flag `destination`
	destination = "destination"
)

var (
	// src defines input file location
	src string
	// dest defines output location
	dest string
	// log defines global log
	log zerolog.Logger
)

func init() {
	// bind sub command
	rootCmd.AddCommand(subS2T)
	rootCmd.AddCommand(subT2S)
	rootCmd.AddCommand(subS2TW)
	rootCmd.AddCommand(subTW2S)
	rootCmd.AddCommand(subS2HK)
	rootCmd.AddCommand(subHK2S)
	rootCmd.AddCommand(subS2TWP)
	rootCmd.AddCommand(subTW2SP)
	rootCmd.AddCommand(subT2TW)
	rootCmd.AddCommand(subT2HK)
	rootCmd.AddCommand(subDownload)

	// bind and check subS2T flag
	subS2T.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subS2T.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subS2T.MarkFlagRequired(source)
	subS2T.MarkFlagRequired(destination)

	// bind and check subT2S flag
	subT2S.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subT2S.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subT2S.MarkFlagRequired(source)
	subT2S.MarkFlagRequired(destination)

	// bind and check subS2TW flag
	subS2TW.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subS2TW.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subS2TW.MarkFlagRequired(source)
	subS2TW.MarkFlagRequired(destination)

	// bind and check subTW2S flag
	subTW2S.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subTW2S.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subTW2S.MarkFlagRequired(source)
	subTW2S.MarkFlagRequired(destination)

	// bind and check subS2HK flag
	subS2HK.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subS2HK.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subS2HK.MarkFlagRequired(source)
	subS2HK.MarkFlagRequired(destination)

	// bind and check subHK2S flag
	subHK2S.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subHK2S.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subHK2S.MarkFlagRequired(source)
	subHK2S.MarkFlagRequired(destination)

	// bind and check subS2TWP flag
	subS2TWP.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subS2TWP.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subS2TWP.MarkFlagRequired(source)
	subS2TWP.MarkFlagRequired(destination)

	// bind and check subTW2SP flag
	subTW2SP.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subTW2SP.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subTW2SP.MarkFlagRequired(source)
	subTW2SP.MarkFlagRequired(destination)

	// bind and check subT2TW flag
	subT2TW.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subT2TW.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subT2TW.MarkFlagRequired(source)
	subT2TW.MarkFlagRequired(destination)

	// bind and check subT2HK flag
	subT2HK.Flags().StringVarP(&src, source, "s", "", "input file name (REQUIRED)")
	subT2HK.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	subT2HK.MarkFlagRequired(source)
	subT2HK.MarkFlagRequired(destination)

	// init chinese config
	gocc.SetConfigPath(textConfigPath)

	// init log setting
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: utils.LogTimeLayout}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("|  %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("FIELD: %s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chiconverter",
	Short: "A chinese converter",
	Long: `A command line tool to convert chinese from
  Simplified to Traditional(Standard/TW/HK)
  Traditional(Standard/TW/HK) to Simplified
  Traditional(Standard) to Traditional(TW/HK).`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
