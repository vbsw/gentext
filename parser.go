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
		parsedArgs.result = result_HELP

	} else if len(clArgs) == 2 {
		parsedArgs = parseOneArgument(parsedArgs, clArgs[1])

	} else if len(clArgs) == 3 {
		parsedArgs = parseTwoArguments(parsedArgs, clArgs[1], clArgs[2])

	} else if len(clArgs) == 4 {
		parsedArgs = parseThreeArguments(parsedArgs, clArgs[1], clArgs[2], clArgs[3])

	} else if len(clArgs) == 5 {
		parsedArgs = parseFourArguments(parsedArgs, clArgs[1], clArgs[2], clArgs[3], clArgs[4])

	} else if len(clArgs) > 5 {
		parsedArgs.result = result_TOO_MANY

	} else {
		parsedArgs = parseTwoArguments(parsedArgs, clArgs[1], clArgs[2])
	}
	return parsedArgs
}

func parseOneArgument(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if parseHelp(argument) {
		parsedArgs.result = result_HELP

	} else if parseVersion(argument) {
		parsedArgs.result = result_VERSION

	} else if parseCopyright(argument) {
		parsedArgs.result = result_COPYRIGHT

	} else {
		parsedArgs = parseWrongOneArgument(parsedArgs, argument)
	}
	return parsedArgs
}

func parseTwoArguments(parsedArgs *tParsedArguments, argument1, argument2 string) *tParsedArguments {
	bytesCount := parseBytesCount(argument1)

	if argument2 == "std" {
		parsedArgs.result = result_OUTPUT_STD

	} else {
		parsedArgs.result = result_OUTPUT_FILE
		parsedArgs.fileName = argument2
	}

	if bytesCount > 0 {
		parsedArgs.bytesCount = bytesCount

	} else {
		parsedArgs = parseWrongTwoArguments(parsedArgs, argument1)
	}
	return parsedArgs
}

func parseThreeArguments(parsedArgs *tParsedArguments, argument1, argument2, argument3 string) *tParsedArguments {
	parsedArgs = parseTwoArguments(parsedArgs, argument1, argument2)

	if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
		parsedArgs = parseGeneratorOption(parsedArgs, argument3)

	} else {
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument3)
	}
	return parsedArgs
}

func parseFourArguments(parsedArgs *tParsedArguments, argument1, argument2, argument3, argument4 string) *tParsedArguments {
	parsedArgs = parseTwoArguments(parsedArgs, argument1, argument2)

	if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
		parsedArgs = parseGeneratorOption(parsedArgs, argument3)

		if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
			parsedArgs = parseGeneratorOption(parsedArgs, argument4)

		} else {
			parsedArgs = parseWrongThreeArguments(parsedArgs, argument4)
		}

	} else {
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument3)
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument4)
	}
	return parsedArgs
}

func parseFiveArguments(parsedArgs *tParsedArguments, argument1, argument2, argument3, argument4, argument5 string) *tParsedArguments {
	parsedArgs = parseTwoArguments(parsedArgs, argument1, argument2)

	if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
		parsedArgs = parseGeneratorOption(parsedArgs, argument3)

		if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
			parsedArgs = parseGeneratorOption(parsedArgs, argument4)

			if parsedArgs.result == result_OUTPUT_STD || parsedArgs.result == result_OUTPUT_FILE {
				parsedArgs = parseGeneratorOption(parsedArgs, argument5)

			} else {
				parsedArgs = parseWrongThreeArguments(parsedArgs, argument5)
			}

		} else {
			parsedArgs = parseWrongThreeArguments(parsedArgs, argument4)
			parsedArgs = parseWrongThreeArguments(parsedArgs, argument5)
		}

	} else {
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument3)
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument4)
		parsedArgs = parseWrongThreeArguments(parsedArgs, argument5)
	}
	return parsedArgs
}

func parseWrongOneArgument(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if len(argument) > 2 {
		argFirstTwoLetters := argument[:2]

		if argFirstTwoLetters == "-c" || argFirstTwoLetters == "-t" {
			number, err := strconv.Atoi(argument[2:])

			if (err != nil || number < 0) && !contains(parsedArgs.invalidArguments, argument) {
				parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
			}

		} else {
			parsedArgs = parseWrongBytesCountArgument(parsedArgs, argument)
		}

	} else {
		parsedArgs = parseWrongBytesCountArgument(parsedArgs, argument)
	}
	return parsedArgs
}

func parseWrongTwoArguments(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if parseOption(argument) {
		parsedArgs.result = result_WRONG_COMBINATION

	} else if len(argument) > 2 {
		argFirstTwoLetters := argument[:2]

		if argFirstTwoLetters == "-c" || argFirstTwoLetters == "-t" {
			number, err := strconv.Atoi(argument[2:])

			if err == nil && number >= 0 {
				parsedArgs.result = result_WRONG_COMBINATION

			} else {
				parsedArgs.result = result_INVALID_ARGUMENTS
				parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
			}

		} else {
			parsedArgs.result = result_INVALID_ARGUMENTS
			parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
		}

	} else {
		parsedArgs.result = result_INVALID_ARGUMENTS
		parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
	}
	return parsedArgs
}

func parseWrongThreeArguments(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if len(argument) > 2 {
		argFirstTwoLetters := argument[:2]

		if argFirstTwoLetters == "-c" || argFirstTwoLetters == "-t" {
			number, err := strconv.Atoi(argument[2:])

			if (err != nil || number < 0) && !contains(parsedArgs.invalidArguments, argument) {
				parsedArgs.result = result_INVALID_ARGUMENTS
				parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
			}

		} else if !parseOption(argument) && !contains(parsedArgs.invalidArguments, argument) {
			parsedArgs.result = result_INVALID_ARGUMENTS
			parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
		}

	} else if !contains(parsedArgs.invalidArguments, argument) {
		parsedArgs.result = result_INVALID_ARGUMENTS
		parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
	}
	return parsedArgs
}

func parseGeneratorOption(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	if len(argument) > 2 {
		argFirstTwoLetters := argument[:2]

		if argFirstTwoLetters == "-c" {
			coresCount, err := strconv.Atoi(argument[2:])

			if err == nil && coresCount >= 0 {
				parsedArgs.coresCount = coresCount

			} else {
				parsedArgs.result = result_INVALID_ARGUMENTS
				parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
			}

		} else if argFirstTwoLetters == "-t" {
			threadsCount, err := strconv.Atoi(argument[2:])

			if err == nil && threadsCount >= 0 {
				parsedArgs.threadsCount = threadsCount

			} else {
				parsedArgs.result = result_INVALID_ARGUMENTS

				if !contains(parsedArgs.invalidArguments, argument) {
					parsedArgs.result = result_INVALID_ARGUMENTS
					parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
				}
			}

		} else if parseOption(argument) {
			parsedArgs.result = result_WRONG_COMBINATION

		} else {
			parsedArgs = parseWrongBytesCountArgument(parsedArgs, argument)
		}

	} else {
		parsedArgs = parseWrongBytesCountArgument(parsedArgs, argument)
	}
	return parsedArgs
}

func parseWrongBytesCountArgument(parsedArgs *tParsedArguments, argument string) *tParsedArguments {
	bytesCount := parseBytesCount(argument)

	if bytesCount > 0 {
		parsedArgs.result = result_WRONG_COMBINATION

	} else if !contains(parsedArgs.invalidArguments, argument) {
		parsedArgs.result = result_INVALID_ARGUMENTS
		parsedArgs.invalidArguments = append(parsedArgs.invalidArguments, argument)
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

func parseOption(argument string) bool {
	return parseHelp(argument) || parseVersion(argument) || parseCopyright(argument)
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

func contains(strArray []string, str string) bool {
	for _, s := range strArray {
		if s == str {
			return true
		}
	}
	return false
}
