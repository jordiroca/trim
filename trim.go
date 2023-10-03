/*

DESCRIPTION: trim, command line program in Go to trim input lines based on the script name.

VERSION: 23.10.2

OPTIONS:
-h, --help    Display this message.

AUTHOR: Jordi Roca
CREATED: 2023/10/02

GITHUB: https://github.com/jordiroca/trim

LICENSE: See LICENSE file.

*/

package main

import (
	"bufio" // buffered I/O
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
##     ##    ###    #### ##    ##
###   ###   ## ##    ##  ###   ##
#### ####  ##   ##   ##  ####  ##
## ### ## ##     ##  ##  ## ## ##
##     ## #########  ##  ##  ####
##     ## ##     ##  ##  ##   ###
##     ## ##     ## #### ##    ##
*/

func main() {
	scriptName := filepath.Base(os.Args[0])

	// Check '-h' or '--help'
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		displayHelp(scriptName)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		switch scriptName {
		case "ltrim", "lmow":
			// Remove leading whitespace (LEFT TRIM)
			// fmt.Fprintf(os.Stderr, "Removing leading WS\n")
			fmt.Println(strings.TrimLeftFunc(line, func(r rune) bool {
				return r == ' ' || r == '\t'
			}))
		case "rtrim", "rmow":
			// Remove trailing whitespace (RIGHT TRIM)
			// fmt.Fprintf(os.Stderr, "Removing trailing WS\n")
			fmt.Println(strings.TrimRightFunc(line, func(r rune) bool {
				return r == ' ' || r == '\t'
			}))
		case "trim", "mow":
			// fmt.Fprintf(os.Stderr, "Removing leading and trailing WS\n")
			fmt.Println(strings.TrimSpace(line))
		default:
			fmt.Fprintf(os.Stderr, "No funciona\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

/*
##     ## ######## ##       ########
##     ## ##       ##       ##     ##
##     ## ##       ##       ##     ##
######### ######   ##       ########
##     ## ##       ##       ##
##     ## ##       ##       ##
##     ## ######## ######## ##
*/

func displayHelp(scriptName string) {
	helpMessage := `
Usage: %s [OPTION]
Trim whitespaces from each line of input text based on the script name.

Script name and functionalities:
  trim:   Removes leading and trailing whitespaces.
  ltrim:  Removes leading whitespaces.
  rtrim:  Removes trailing whitespaces.

Options:
  -h, --help    Display this help message and exit.

Example usage:
  echo -e "  hello \n world  " | %s
`
	fmt.Printf(helpMessage, scriptName, scriptName)
}
