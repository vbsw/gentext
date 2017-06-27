/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

type tParsedArguments struct {
	err error
	helpArg bool
	versionArg bool
	copyrightArg bool
	unknownArgs []string
	bytesCount int
	outputFile string
	outputStd bool
	cores int
	threads int
	tooMany bool
}
