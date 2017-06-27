/*
 *          Copyright 2017, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
	"os"
)

func generateText(parsedArgs *tParsedArguments) {
	if parsedArgs.outputStd {
		generateOutputToStd(parsedArgs.bytesCount, parsedArgs.cores, parsedArgs.threads)

	} else {
		generateOutputToFile(parsedArgs.bytesCount, parsedArgs.outputFile, parsedArgs.cores, parsedArgs.threads)
	}
}

func generateOutputToStd(bytesCount int, cores int, threads int) {
	str := "sample."

	if bytesCount > 150000 {
		bytesCount = 150
	}

	for i := 0; i < bytesCount; i += len(str) {
		bytesLeft := bytesCount - i

		if bytesLeft < len(str) {
			fmt.Print(str[:bytesLeft])

		} else {
			fmt.Print(str)
		}
	}
	fmt.Println()
}

func generateOutputToFile(bytesCount int, fileName string, cores int, threads int) {
	str := "sample."

	// open output file
	fo, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < bytesCount; i += len(str) {
		bytesLeft := bytesCount - i

		if bytesLeft < len(str) {
			_, err := fo.WriteString(str[:bytesLeft])

			if err != nil {
				panic(err)
			}

		} else {
			_, err := fo.WriteString(str)

			if err != nil {
				panic(err)
			}
		}
	}
}
