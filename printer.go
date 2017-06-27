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
	if parsedArgs.tooMany {
		printTooManyArguments()

	} else if len(parsedArgs.unknownArgs) > 0 {
		printUnknownArguments(parsedArgs.unknownArgs)

	} else if validHelp(parsedArgs) {
		printHelp()

	} else if validVersion(parsedArgs) {
		printVersion()

	} else if validCopyright(parsedArgs) {
		printCopyright()

	} else {
		printWrongCombination()
	}
}

func printTooManyArguments() {
	fmt.Println("Error: too many arguments")
}

func printUnknownArguments(unknownArguments []string) {
	if len(unknownArguments) == 1 {
		fmt.Print("Error: unknown argument")

	} else {
		fmt.Print("Error: unknown arguments")
	}

	for _, arg := range unknownArguments {
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
	fmt.Println("  gentext SIZE OUTPUT {THREAD-OPTIONS}")
	fmt.Println("OPTION")
	fmt.Println("  -help       prints this help")
	fmt.Println("  -version    prints version numer of gentext")
	fmt.Println("  -copyright  prints copyright of gentext")
	fmt.Println("THREAD-OPTION")
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
