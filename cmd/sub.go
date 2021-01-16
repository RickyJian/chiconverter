package cmd

import (
	"bufio"
	"fmt"
	"os"

	"chiconverter/utils"
	"github.com/RickyJian/gocc"
	"github.com/spf13/cobra"
)

const (
	// newLineDelimiter defines new line delimiter
	newLineDelimiter = '\n'
	// s2t defines simplified to traditional subcommand description
	s2t = "s2t"
	// t2s defines traditional to simplified subcommand description
	t2s = "t2s"
	// s2tw defines simplified to traditional(TW) subcommand description
	s2tw = "s2tw"
	// tw2s defines traditional(TW) simplified subcommand description
	tw2s = "tw2s"
	// s2t defines simplified to traditional subcommand description
	s2hk = "s2hk"
	// s2t defines simplified to traditional subcommand description
	hk2s = "hk2s"
	// s2t defines simplified to traditional subcommand description
	s2twp = "s2twp"
	// s2t defines simplified to traditional subcommand description
	tw2sp = "tw2sp"
	// s2t defines simplified to traditional subcommand description
	t2tw = "t2tw"
	// s2t defines simplified to traditional subcommand description
	t2hk = "t2hk"
)

// run defines subcommand run function
var run = func(cmd *cobra.Command, args []string) {
	log.Info().Msg("convert start")
	defer log.Info().Msg("convert finish")

	destName, err := utils.DestFileName(src, dest)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	texts, err := utils.ReadAll(src)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	contents, err := utils.Convert(NewConverter(cmd.Name(), texts))
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	file, err := os.Create(destName)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < len(contents); i++ {
		writer.WriteString(contents[i])
		writer.WriteByte(newLineDelimiter)
	}
}

// NewConverter returns converter
func NewConverter(subCommand string, texts [][]byte) utils.Converter {
	switch subCommand {
	case s2t:
		return S2T(texts)
	case t2s:
		return T2S(texts)
	case s2tw:
		return S2TW(texts)
	case tw2s:
		return TW2S(texts)
	case s2hk:
		return S2HK(texts)
	case hk2s:
		return HK2S(texts)
	case s2twp:
		return S2TWP(texts)
	case tw2sp:
		return TW2SP(texts)
	case t2tw:
		return T2TW(texts)
	case t2hk:
		return T2HK(texts)
	default:
		panic(fmt.Errorf("%s convert unimplemented", subCommand))
	}
}

// subS2T defines sub command to convert simplified to traditional
var subS2T = &cobra.Command{
	Use:   s2t,
	Short: "Convert Simplified Chinese to Traditional Chinese.",
	Long:  `Convert Simplified Chinese to Traditional Chinese.`,
	Run:   run,
}

// S2T defines simplified to traditional type
type S2T [][]byte

// Convert simplified convert to traditional chinese
func (c S2T) Convert() (map[int]string, error) {
	cvt, err := gocc.New(s2t)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subT2S defines sub command to convert traditional to simplified
var subT2S = &cobra.Command{
	Use:   t2s,
	Short: "Convert Traditional Chinese to Simplified Chinese.",
	Long:  `Convert Traditional Chinese to Simplified Chinese.`,
	Run:   run,
}

// T2S defines traditional to simplified type
type T2S [][]byte

// Convert traditional convert to simplified chinese
func (c T2S) Convert() (map[int]string, error) {
	cvt, err := gocc.New(t2s)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subS2TW defines sub command to convert simplified to traditional(TW).
var subS2TW = &cobra.Command{
	Use:   s2tw,
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard).`,
	Run:   run,
}

// S2TW defines simplified to traditional(TW) type
type S2TW [][]byte

// Convert simplified to traditional(TW) chinese
func (c S2TW) Convert() (map[int]string, error) {
	cvt, err := gocc.New(s2tw)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subTW2S defines sub command to convert traditional chinese (taiwan standard) to simplified chinese
var subTW2S = &cobra.Command{
	Use:   tw2s,
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese.`,
	Run:   run,
}

// TW2S defines traditional(TW) to simplified type
type TW2S [][]byte

// Convert traditional(TW) to simplified chinese
func (c TW2S) Convert() (map[int]string, error) {
	cvt, err := gocc.New(tw2s)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subS2HK defines sub command to convert simplified chinese to traditional chinese (hong kong standard)
var subS2HK = &cobra.Command{
	Use:   s2hk,
	Short: "Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Hong Kong Standard).`,
	Run:   run,
}

// S2HK defines simplified to traditional(HK) type
type S2HK [][]byte

// Convert simplified to traditional(HK) chinese
func (c S2HK) Convert() (map[int]string, error) {
	cvt, err := gocc.New(s2hk)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subHK2S defines sub command to convert traditional chinese (hong kong standard) to simplified chinese
var subHK2S = &cobra.Command{
	Use:   hk2s,
	Short: "Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.",
	Long:  `Convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese.`,
	Run:   run,
}

// HK2S defines traditional(HK) to simplified type
type HK2S [][]byte

// Convert traditional(HK) to simplified chinese
func (c HK2S) Convert() (map[int]string, error) {
	cvt, err := gocc.New(hk2s)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subS2TWP defines sub command to convert simplified chinese to traditional chinese (taiwan standard) with taiwanese idiom
var subS2TWP = &cobra.Command{
	Use:   s2twp,
	Short: "Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.",
	Long:  `Convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom.`,
	Run:   run,
}

// S2TWP defines simplified to traditional(TWP) type
type S2TWP [][]byte

// Convert simplified to traditional(TWP) chinese
func (c S2TWP) Convert() (map[int]string, error) {
	cvt, err := gocc.New(s2twp)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subTW2SP defines sub command to convert traditional chinese (taiwan standard) to simplified chinese with mainland chinese idiom
var subTW2SP = &cobra.Command{
	Use:   tw2sp,
	Short: "Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.",
	Long:  `Convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom.`,
	Run:   run,
}

// TW2SP defines traditional(TW) to simplified(P) type
type TW2SP [][]byte

// Convert traditional(TW) to simplified(P) chinese
func (c TW2SP) Convert() (map[int]string, error) {
	cvt, err := gocc.New(tw2sp)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// subT2TW defines sub command to convert traditional chinese (openCC standard) to taiwan standard
var subT2TW = &cobra.Command{
	Use:   t2tw,
	Short: "Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Taiwan Standard.`,
	Run:   run,
}

// T2TW defines traditional to traditional(TW) type
type T2TW [][]byte

// Convert traditional(TW) to simplified(P) chinese
func (c T2TW) Convert() (map[int]string, error) {
	cvt, err := gocc.New(t2tw)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}

// T2HK defines sub command to convert traditional chinese (openCC standard) to hong kong standard
var subT2HK = &cobra.Command{
	Use:   t2hk,
	Short: "Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.",
	Long:  `Convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard.`,
	Run:   run,
}

// T2HK defines traditional to traditional(HK) type
type T2HK [][]byte

// Convert traditional to traditional(HK) chinese
func (c T2HK) Convert() (map[int]string, error) {
	cvt, err := gocc.New(t2hk)
	if err != nil {
		return map[int]string{}, err
	}
	return utils.Batch(c, cvt.Convert)
}
