package term

import (
	"fmt"
)

const (
	bold       = "1"
	dim        = "2"
	blue       = "0;34"
	green      = "0;32"
	red        = "0;31"
	yellow     = "0;33"
	cyan       = "0;36"
	boldBlue   = "1;34"
	boldGreen  = "1;32"
	boldRed    = "1;31"
	boldYellow = "1;33"
	boldCyan   = "1;36"
	dimBlue    = "2;34"
	dimGreen   = "2;32"
	dimRed     = "2;31"
	dimYellow  = "2;33"
	dimCyan    = "2;36"
)

// Bold returns a bold string
func Bold(msg string) string {
	return Colour(bold, msg)
}

// Boldf returns a bold string, except that the msg will be formated according
// to the format specifier
//
func Boldf(msg string, v ...interface{}) string {
	return Colourf(bold, msg, v)
}

// Dim returns a dim string
func Dim(msg string) string {
	return Colour(dim, msg)
}

// Dimf returns a dim string, except that the msg will be formated according
// to the format specifier
//
func Dimf(msg string, v ...interface{}) string {
	return Colourf(dim, msg, v)
}

// Blue returns a blue string
func Blue(msg string) string {
	return Colour(blue, msg)
}

// Bluef returns a blue string, except that the msg will be formated according
// to the format specifier
//
func Bluef(msg string, v ...interface{}) string {
	return Colourf(blue, msg, v)
}

// Green returns a Green string
func Green(msg string) string {
	return Colour(green, msg)
}

// Greenf returns a Green string, except that the msg will be formated according
// to the format specifier
//
func Greenf(msg string, v ...interface{}) string {
	return Colourf(green, msg, v)
}

// Red returns a Red string
func Red(msg string) string {
	return Colour(red, msg)
}

// Redf returns a Red string, except that the msg will be formated according
// to the format specifier
//
func Redf(msg string, v ...interface{}) string {
	return Colourf(red, msg, v)
}

// Yellow returns a Yellow string
func Yellow(msg string) string {
	return Colour(yellow, msg)
}

// Yellowf returns a Yellow string, except that the msg will be formated according
// to the format specifier
//
func Yellowf(msg string, v ...interface{}) string {
	return Colourf(yellow, msg, v)
}

// Cyan returns a Cyan string
func Cyan(msg string) string {
	return Colour(cyan, msg)
}

// Cyanf returns a Cyan string, except that the msg will be formated according
// to the format specifier
//
func Cyanf(msg string, v ...interface{}) string {
	return Colourf(cyan, msg, v)
}

// BoldBlue returns a BoldBlue string
func BoldBlue(msg string) string {
	return Colour(boldBlue, msg)
}

// BoldBluef returns a BoldBlue string, except that the msg will be formated according
// to the format specifier
//
func BoldBluef(msg string, v ...interface{}) string {
	return Colourf(boldBlue, msg, v)
}

// BoldGreen returns a BoldGreen string
func BoldGreen(msg string) string {
	return Colour(boldGreen, msg)
}

// BoldGreenf returns a BoldGreen string, except that the msg will be formated according
// to the format specifier
//
func BoldGreenf(msg string, v ...interface{}) string {
	return Colourf(boldGreen, msg, v)
}

// BoldRed returns a BoldRed string
func BoldRed(msg string) string {
	return Colour(boldRed, msg)
}

// BoldRedf returns a BoldRed string, except that the msg will be formated according
// to the format specifier
//
func BoldRedf(msg string, v ...interface{}) string {
	return Colourf(boldRed, msg, v)
}

// BoldYellow returns a BoldYellow string
func BoldYellow(msg string) string {
	return Colour(boldYellow, msg)
}

// BoldYellowf returns a BoldYellow string, except that the msg will be formated according
// to the format specifier
//
func BoldYellowf(msg string, v ...interface{}) string {
	return Colourf(boldYellow, msg, v)
}

// BoldCyan returns a BoldCyan string
func BoldCyan(msg string) string {
	return Colour(boldCyan, msg)
}

// BoldCyanf returns a BoldCyan string, except that the msg will be formated according
// to the format specifier
//
func BoldCyanf(msg string, v ...interface{}) string {
	return Colourf(boldCyan, msg, v)
}

// DimBlue returns a DimBlue string
func DimBlue(msg string) string {
	return Colour(dimBlue, msg)
}

// DimBluef returns a DimBlue string, except that the msg will be formated according
// to the format specifier
//
func DimBluef(msg string, v ...interface{}) string {
	return Colourf(dimBlue, msg, v)
}

// DimGreen returns a DimGreen string
func DimGreen(msg string) string {
	return Colour(dimGreen, msg)
}

// DimGreenf returns a DimGreen string, except that the msg will be formated according
// to the format specifier
//
func DimGreenf(msg string, v ...interface{}) string {
	return Colourf(dimGreen, msg, v)
}

// DimRed returns a DimRed string
func DimRed(msg string) string {
	return Colour(dimRed, msg)
}

// DimRedf returns a DimRed string, except that the msg will be formated according
// to the format specifier
//
func DimRedf(msg string, v ...interface{}) string {
	return Colourf(dimRed, msg, v)
}

// DimYellow returns a DimYellow string
func DimYellow(msg string) string {
	return Colour(dimYellow, msg)
}

// DimYellowf returns a DimYellow string, except that the msg will be formated according
// to the format specifier
//
func DimYellowf(msg string, v ...interface{}) string {
	return Colourf(dimYellow, msg, v)
}

// DimCyan returns a DimCyan string
func DimCyan(msg string) string {
	return Colour(dimCyan, msg)
}

// DimCyanf returns a DimCyan string, except that the msg will be formated according
// to the format specifier
//
func DimCyanf(msg string, v ...interface{}) string {
	return Colourf(dimCyan, msg, v)
}

// Colour returns a string
func Colour(code string, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}

// Colourf returns a string, except that the msg will be formated according
// to the format specifier
//
func Colourf(code string, msg string, v []interface{}) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, fmt.Sprintf(msg, v...))
}

// HighlightSource returns a string colour to indicate that the contents
// represents source code.
func HighlightSource(source string) string {
	return Cyan(source)
}

// HighlightSourcef returns a string colour to indicate that the contents
// represents source code.
func HighlightSourcef(source string, v ...interface{}) string {
	return Cyanf(source, v...)
}
