/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"os"
)

func main() {
	parsedArgs := parseArguments(os.Args)

	if validCommandLineInfoOnly(parsedArgs) {
		printCommandLineInfo(parsedArgs)

	} else {
		generateText(parsedArgs)
	}
}
