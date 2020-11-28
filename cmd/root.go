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

	// bind and check s2t flag
	s2t.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	s2t.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	s2t.MarkFlagRequired(source)
	s2t.MarkFlagRequired(destination)

	// bind and check t2s flag
	t2s.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	t2s.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	t2s.MarkFlagRequired(source)
	t2s.MarkFlagRequired(destination)

	// bind and check s2tw flag
	s2tw.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	s2tw.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	s2tw.MarkFlagRequired(source)
	s2tw.MarkFlagRequired(destination)

	// bind and check tw2s flag
	tw2s.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	tw2s.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	tw2s.MarkFlagRequired(source)
	tw2s.MarkFlagRequired(destination)

	// bind and check s2hk flag
	s2hk.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	s2hk.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	s2hk.MarkFlagRequired(source)
	s2hk.MarkFlagRequired(destination)

	// bind and check hk2s flag
	hk2s.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	hk2s.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	hk2s.MarkFlagRequired(source)
	hk2s.MarkFlagRequired(destination)

	// bind and check s2twp flag
	s2twp.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	s2twp.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	s2twp.MarkFlagRequired(source)
	s2twp.MarkFlagRequired(destination)

	// bind and check tw2sp flag
	tw2sp.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	tw2sp.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	tw2sp.MarkFlagRequired(source)
	tw2sp.MarkFlagRequired(destination)

	// bind and check t2tw flag
	t2tw.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	t2tw.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	t2tw.MarkFlagRequired(source)
	t2tw.MarkFlagRequired(destination)

	// bind and check t2hk flag
	t2hk.Flags().StringVarP(&src, source, "s", "", "input file location (REQUIRED)")
	t2hk.Flags().StringVarP(&dest, destination, "d", "", "output file location (REQUIRED)")
	t2hk.MarkFlagRequired(source)
	t2hk.MarkFlagRequired(destination)
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
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
