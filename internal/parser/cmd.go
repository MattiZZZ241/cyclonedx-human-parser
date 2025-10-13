package parser

import (
	"fmt"
	"strings"
)

func parseArgs(args []string, opts *Options) {
	if len(args) == 0 {
		return
	}

	arg := args[0]

	// Check for flags
	if strings.HasPrefix(arg, "-") {
		switch arg {
		case "-v", "--verbose":
			opts.Verbose = true
		case "-a", "--all":
			opts.ShowAll = true
		default:
			fmt.Printf("Unknown flag: %s\n", arg)
		}
		parseArgs(args[1:], opts) // Recur with remaining args
	} else {
		// Assume it's a file path or argument
		opts.OtherArgs = append(opts.OtherArgs, arg)
		parseArgs(args[1:], opts)
	}
}
