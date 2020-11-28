package cmd

import (
	"github.com/spf13/cobra"
)

// s2t defines sub command to convert simplified chinese to traditional chinese
var s2t = &cobra.Command{
	Use:   "s2t",
	Short: "Convert Simplified Chinese to Traditional Chinese.",
	Long:  `Convert Simplified Chinese to Traditional Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// t2s defines sub command to convert traditional chinese to simplified chinese
var t2s = &cobra.Command{
	Use:   "t2s",
	Short: "Convert Traditional Chinese to Simplified Chinese.",
	Long:  `Convert Traditional Chinese to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// s2tw defines sub command to convert simplified chinese to traditional chinese (taiwan standard).
var s2tw = &cobra.Command{
	Use:   "s2tw",
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// tw2s defines sub command to convert traditional chinese (taiwan standard) to simplified chinese
var tw2s = &cobra.Command{
	Use:   "tw2s",
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// s2hk defines sub command to convert simplified chinese to traditional chinese (hong kong standard)
var s2hk = &cobra.Command{
	Use:   "s2hk",
	Short: "Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// s2tw defines sub command to convert traditional chinese (hong kong standard) to simplified chinese
var hk2s = &cobra.Command{
	Use:   "s2tw",
	Short: "Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// s2twp defines sub command to convert simplified chinese to traditional chinese (taiwan standard) with taiwanese idiom
var s2twp = &cobra.Command{
	Use:   "s2twp",
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// tw2sp defines sub command to convert traditional chinese (taiwan standard) to simplified chinese with mainland chinese idiom
var tw2sp = &cobra.Command{
	Use:   "tw2sp",
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// t2tw defines sub command to convert traditional chinese (openCC standard) to taiwan standard
var t2tw = &cobra.Command{
	Use:   "t2tw",
	Short: "Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// t2hk defines sub command to convert traditional chinese (openCC standard) to hong kong standard
var t2hk = &cobra.Command{
	Use:   "t2hk",
	Short: "Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
