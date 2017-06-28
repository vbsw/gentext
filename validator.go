/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

func validCommandLineInfoOnly(parsedArgs *tParsedArguments) bool {
	return parsedArgs.result != result_OUTPUT_STD && parsedArgs.result != result_OUTPUT_FILE
}

func validHelp(parsedArgs *tParsedArguments) bool {
	return parsedArgs.result == result_HELP
}

func validVersion(parsedArgs *tParsedArguments) bool {
	return parsedArgs.result == result_VERSION
}

func validCopyright(parsedArgs *tParsedArguments) bool {
	return parsedArgs.result == result_COPYRIGHT
}
