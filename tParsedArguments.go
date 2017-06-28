/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

const (
	result_HELP int = iota
	result_VERSION
	result_COPYRIGHT
	result_OUTPUT_STD
	result_OUTPUT_FILE
	result_TOO_MANY
	result_INVALID_ARGUMENTS
	result_WRONG_COMBINATION
)

type tParsedArguments struct {
	result           int
	bytesCount       int
	fileName         string
	coresCount       int
	threadsCount     int
	invalidArguments []string
}
