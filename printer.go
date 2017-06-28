/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
)

func printCommandLineInfo(parsedArgs *tParsedArguments) {
	if parsedArgs.result == result_HELP {
		printHelp()

	} else if parsedArgs.result == result_VERSION {
		printVersion()

	} else if parsedArgs.result == result_COPYRIGHT {
		printCopyright()

	} else if parsedArgs.result == result_TOO_MANY {
		printTooManyArguments()

	} else if parsedArgs.result == result_INVALID_ARGUMENTS {
		printInvalidArguments(parsedArgs.invalidArguments)

	} else if parsedArgs.result == result_WRONG_COMBINATION {
		printWrongCombination()
	}
}

func printTooManyArguments() {
	fmt.Println("Error: too many arguments")
}

func printInvalidArguments(invalidArguments []string) {
	if len(invalidArguments) == 1 {
		fmt.Print("Error: invalid argument")

	} else {
		fmt.Print("Error: invalid arguments")
	}

	for _, arg := range invalidArguments {
		fmt.Print(" ", arg)
	}
	fmt.Println()
}

func printWrongCombination() {
	fmt.Println("Error: wrong combination of arguments")
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  gentext [OPTIONS]")
	fmt.Println("  gentext SIZE OUTPUT {GENERATOR-OPTIONS}")
	fmt.Println("OPTION")
	fmt.Println("  -help       prints this help")
	fmt.Println("  -version    prints version numer of gentext")
	fmt.Println("  -copyright  prints copyright of gentext")
	fmt.Println("GENERATOR-OPTIONS")
	fmt.Println("  -cN         sets the number of logical CPUs to N")
	fmt.Println("  -tN         sets the number of threads to N")
	fmt.Println("Example")
	fmt.Println("  gentext 100 std        prints 100 bytes of text to standard output")
	fmt.Println("  gentext 100K test.txt  prints 100 kibibytes of text to file test.txt")
}

func printVersion() {
	fmt.Println("0.1.0")
}

func printCopyright() {
	fmt.Println("Copyright 2017, Vitali Baumtrok (vbsw@mailbox.org).")
	fmt.Println("gentext is distributed under the Boost Software License, version 1.0.")
	fmt.Println("(See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)")
}
