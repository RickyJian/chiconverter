/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

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
)

func init() {
	// bind sub command
	rootCmd.AddCommand(s2t)
	rootCmd.AddCommand(t2s)
	rootCmd.AddCommand(s2tw)
	rootCmd.AddCommand(tw2s)
	rootCmd.AddCommand(s2hk)
	rootCmd.AddCommand(hk2s)
	rootCmd.AddCommand(s2twp)
	rootCmd.AddCommand(tw2sp)
	rootCmd.AddCommand(t2tw)
	rootCmd.AddCommand(t2hk)

	// bind persistence flag
	rootCmd.PersistentFlags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	rootCmd.PersistentFlags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	rootCmd.MarkPersistentFlagRequired(source)
	rootCmd.MarkPersistentFlagRequired(destination)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chiconverter",
	Short: "A chinese converter",
	Long: `A command line tool to convert chinese from
  Simplified to Traditional(Standard/TW/HK)
  Traditional(Standard/TW/HK) to Simplified
  Traditional(Standard) to Traditional(TW/HK)
Default translation is Simplified to Traditional(Standard).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var s2t = &cobra.Command{
	Use:   "s2t",
	Short: "Convert Simplified Chinese to Traditional Chinese.",
	Long:  `Convert Simplified Chinese to Traditional Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var t2s = &cobra.Command{
	Use:   "t2s",
	Short: "Convert Traditional Chinese to Simplified Chinese.",
	Long:  `Convert Traditional Chinese to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var s2tw = &cobra.Command{
	Use:   "s2tw",
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var tw2s = &cobra.Command{
	Use:   "s2tw",
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var s2hk = &cobra.Command{
	Use:   "s2hk",
	Short: "Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var hk2s = &cobra.Command{
	Use:   "s2tw",
	Short: "Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var s2twp = &cobra.Command{
	Use:   "s2twp",
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var tw2sp = &cobra.Command{
	Use:   "tw2sp",
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var t2tw = &cobra.Command{
	Use:   "tw2sp",
	Short: "Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var t2hk = &cobra.Command{
	Use:   "t2hk",
	Short: "Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
