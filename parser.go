/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"strconv"
)

func parseArguments(clArgs []string) *tParsedArguments {
	parsedArgs := new(tParsedArguments)

	if len(clArgs) == 1 {
		parsedArgs.helpArg = true

	} else if len(clArgs) == 2 {
		parsedArgs = parseOneArgument(parsedArgs, clArgs[1])

	} else if len(clArgs) == 3 {
		parsedArgs = parseTwoArguments(parsedArgs, clArgs[1], clArgs[2])

	} else if len(clArgs) > 5 {
		parsedArgs.tooMany = true

	} else {
		parsedArgs = parseTwoArguments(parsedArgs, clArgs[1], clArgs[2])
	}
	return parsedArgs
}

func parseOneArgument(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if parseHelp(argument) {
		parsedArgs.helpArg = true

	} else if parseVersion(argument) {
		parsedArgs.versionArg = true

	} else if parseCopyright(argument) {
		parsedArgs.copyrightArg = true

	} else {
		parsedArgs.unknownArgs = append(parsedArgs.unknownArgs, argument)
	}
	return parsedArgs
}

func parseTwoArguments(parsedArgs *tParsedArguments, argument1 string, argument2 string) *tParsedArguments {
	bytesCount := parseBytesCount(argument1)

	if argument2 == "std" {
		parsedArgs.outputStd = true

	} else {
		parsedArgs.outputFile = argument2
	}

	if bytesCount > 0 {
		parsedArgs.bytesCount = bytesCount

	} else {
		parsedArgs = parseOneArgument(parsedArgs, argument1)
	}
	return parsedArgs
}

func parseBytesCount(argument string) int {
	lastByte := argument[len(argument)-1]
	multiplier := parseMultiplierByte(lastByte)

	if multiplier > 0 {
		if multiplier > 1 {
			if len(argument) > 1 {
				argument = argument[:len(argument)-1]
			} else {
				argument = "1"
			}
		}
		bytesCount, err := strconv.Atoi(argument)

		if err == nil {
			return bytesCount * multiplier

		} else {
			return 0
		}

	} else {
		return 0
	}
}

func parseMultiplierByte(multiplierByte byte) int {
	if multiplierByte == 'K' || multiplierByte == 'k' {
		return 1024

	} else if multiplierByte == 'M' || multiplierByte == 'm' {
		return 1024 * 1024

	} else if multiplierByte == 'G' || multiplierByte == 'g' {
		return 1024 * 1024 * 1024

	} else if multiplierByte == 'T' || multiplierByte == 't' {
		return 1024 * 1024 * 1024 * 1024

	} else if multiplierByte >= '0' && multiplierByte <= '9' {
		return 1

	} else {
		return 0
	}
}

func parseHelp(argument string) bool {
	return argument == "--help" || argument == "-help" || argument == "-h" || argument == "--usage" || argument == "-usage"
}

func parseVersion(argument string) bool {
	return argument == "--version" || argument == "-version" || argument == "-v"
}

func parseCopyright(argument string) bool {
	return argument == "--copyright" || argument == "-copyright"
}
