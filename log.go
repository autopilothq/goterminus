package term

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

var (
	warningLabel = Yellow("WARNING")

	// Use
	controlCharStart = regexp.MustCompile(`\\[0-9]{3}\[[0-9];[0-9]{1,2}m`)
	controlCharEnd   = regexp.MustCompile(`\\[0-9]{3}\[m`)
)

// Info outputs generate info.
func Info(msg string) {
	fmt.Println(msg)
}

// Infof is identical to Info, except that the msg will be formated according
// to the format specifier
//
func Infof(msg string, v ...interface{}) {
	fmt.Printf(msg+"\n", v...)
}

// Action outputs a specific named action, and a accompying message
//
// For example
//    Action("Created", "Resource `Legacy`")
//
func Action(actionName string, msg string) {
	fmt.Printf("  %25s  %s\n", Green(actionName), msg)
}

// Actionf is identical to Action, except that the msg will be formated according
// to the format specifier
//
// For example
//    Actionf("Created", "Resource `%s`", resourceName)
//
func Actionf(actionName string, msg string, v ...interface{}) {
	colorAction := Green(actionName)
	fmt.Printf("  %25s  %s\n", colorAction, fmt.Sprintf(msg, v...))
}

// Warn outputs the message, highlighted to indicate that it's a warning
//
// For example
//    Warn("Resource `Legacy` already existed, ignored")
//
func Warn(msg string) {
	fmt.Printf("%s %s\n", warningLabel, msg)
}

// Warnf is identical to Warn, except that the msg will be formated according
// to the format specifier
//
// For example
//    Warnf("Resource `%s` already existed, ignored", resourceName)
//
func Warnf(msg string, v ...interface{}) {
	fmt.Printf("%s %s\n", warningLabel, fmt.Sprintf(msg, v...))
}

// Error outputs the message, highlighted to indicate that it's an error
//
// For example
//    Error("Resource `Legacy` was invalid")
//
func Error(msg string) {
	colouredMsg := Redf("\n\tERROR %s", msg)
	fmt.Println(colouredMsg)
}

// Errorf is identical to Error, except that the msg will be formated according
// to the format specifier
//
// For example
//    Errorf("Resource `%s` was invalid: %v", resourceName, err)
//
func Errorf(msg string, v ...interface{}) {
	colouredMsg := Redf("\n\tERROR "+msg, v...)
	fmt.Println(colouredMsg)
}

// List renders a list of strings in one or more columns
func List(items []string, numColumns int, columnWidth int) {
	rows := make([][]string, 0)
	size := len(items)

	// If there's too little content then we'll just render it in a single column
	if size < numColumns {
		for _, item := range items {
			fmt.Println(item)
		}

		fmt.Println("")
		return
	}

	itemsPerCol := int(math.Ceil(float64(size) / float64(numColumns)))
	offset := 0

	for ri := 0; ri < itemsPerCol; ri++ {
		row := make([]string, numColumns)

		for col := 0; col < numColumns; col++ {
			if offset < size {
				row[col] = items[offset]
				offset++
			} else {
				offset++
				continue
			}
		}

		rows = append(rows, row)
	}

	Table(nil, rows...)
}

func Table(headers []string, rows ...[]string) {
	var numColumns int

	// We use the number of headers to know how many columns to render, but
	// headers are optional so we need a fallback way of doing this.
	if headers != nil {
		numColumns = len(headers)
	} else {
		// TODO probably should find the row with the max number of columns
		numColumns = len(rows[0])
	}

	maxCellWidth := findMaxCellWidth(numColumns, headers, rows) + 4
	w := tabwriter.NewWriter(os.Stdout, maxCellWidth, 8, 4, ' ', 0)
	w.Write([]byte("\n"))

	if headers != nil {
		w.Write([]byte(strings.Join(headers, "\t") + "\f"))
		border := strings.Repeat("-", maxCellWidth*numColumns+(4*numColumns-1)) + "\f"
		w.Write([]byte(border))
	}

	for _, row := range rows {
		w.Write([]byte(strings.Join(row, "\t") + "\f"))
	}

	w.Write([]byte("\n"))
	w.Flush()
}

func findMaxCellWidth(numColumns int, headers []string, rows [][]string) int {
	var maxWidth int

	if headers != nil {
		for ci := 0; ci < numColumns-1; ci++ {
			l := cellLength(headers[ci])
			if l > maxWidth {
				maxWidth = l
			}
		}
	}

	for _, row := range rows {
		for ci := 0; ci < numColumns-1; ci++ {
			l := cellLength(row[ci])
			if l > maxWidth {
				maxWidth = l
			}
		}
	}

	return maxWidth
}

func cellLength(s string) int {
	return len(controlCharEnd.ReplaceAllString(controlCharStart.ReplaceAllString(s, ""), ""))
}
