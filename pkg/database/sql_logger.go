package postgres

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const bold = "\033[1m"
const greenBg = "\033[42m"
const redBg = "\033[41m"
const greenText = "\033[32m"
const lightBlueText = "\033[96m"
const reset = "\033[0m"

func LogExecutionTime(label string, start time.Time) {
	log.Printf("%s took %v", label, time.Since(start))
}

func PrintSql(success bool, query string, start *time.Time, args ...interface{}) {
	formattedArgs := formatArgs(args)

	var bgColor string

	if success {
		bgColor = greenBg
	} else {
		bgColor = redBg
	}

	fmt.Println(
		bgColor + "[SQL]" + reset + " " +
			lightBlueText + bold + query + " " + strconv.FormatInt(time.Since(*start).Milliseconds(), 10) + "ms " + reset + greenText + formattedArgs + reset,
	)
}

func formatArgs(args []interface{}) string {
	formatted := "["
	for i, param := range args {
		formatted += fmt.Sprintf("'%v'", param)
		if i < len(args)-1 {
			formatted += ", "
		}
	}
	formatted += "]"
	return formatted
}
