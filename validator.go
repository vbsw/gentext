/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

func validCommandLineInfoOnly(parsedArgs *tParsedArguments) bool {
	if parsedArgs.helpArg {
		return true

	} else if parsedArgs.versionArg {
		return true

	} else if parsedArgs.copyrightArg {
		return true

	} else if len(parsedArgs.unknownArgs) > 0 {
		return true

	} else if parsedArgs.err != nil {
		return true

	} else if parsedArgs.tooMany {
		return true

	} else if parsedArgs.outputStd {
		return false

	} else if len(parsedArgs.outputFile) > 0 {
		return false

	} else {
		return true
	}
}

func validHelp(parsedArgs *tParsedArguments) bool {
	return parsedArgs.helpArg && !parsedArgs.versionArg && !parsedArgs.copyrightArg && !parsedArgs.outputStd && len(parsedArgs.outputFile) == 0
}

func validVersion(parsedArgs *tParsedArguments) bool {
	return parsedArgs.versionArg && !parsedArgs.helpArg && !parsedArgs.copyrightArg && !parsedArgs.outputStd && len(parsedArgs.outputFile) == 0
}

func validCopyright(parsedArgs *tParsedArguments) bool {
	return parsedArgs.copyrightArg && !parsedArgs.versionArg && !parsedArgs.helpArg && !parsedArgs.outputStd && len(parsedArgs.outputFile) == 0
}
