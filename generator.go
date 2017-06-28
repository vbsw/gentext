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
	parsedArgs.threadsCount = adjustThreadsCount(parsedArgs.coresCount, parsedArgs.threadsCount)

	if parsedArgs.result == result_OUTPUT_STD {
		generateOutputToStd(parsedArgs.bytesCount, parsedArgs.coresCount, parsedArgs.threadsCount)

	} else {
		generateOutputToFile(parsedArgs.fileName, parsedArgs.bytesCount, parsedArgs.coresCount, parsedArgs.threadsCount)
	}
}

func adjustThreadsCount(coresCount, threadsCount int) int {
	if coresCount > 1 && threadsCount < 2 {
		return int(float32(coresCount) * 1.5)

	} else {
		return coresCount
	}
}

func generateOutputToStd(bytesCount, coresCount, threadsCount int) {
	if coresCount > 1 {
		generateOutputToStdMultiThreaded(bytesCount, coresCount, threadsCount)

	} else {
		generateOutputToStdSingleThreaded(bytesCount)
	}
}

func generateOutputToFile(fileName string, bytesCount, coresCount, threadsCount int) {
	if coresCount > 1 {
		generateOutputToFileMultiThreaded(fileName, bytesCount, coresCount, threadsCount)

	} else {
		generateOutputToFileSingleThreaded(fileName, bytesCount)
	}
}

func generateOutputToStdMultiThreaded(bytesCount, coresCount, threadsCount int) {
	generateOutputToStdSingleThreaded(bytesCount)
}

func generateOutputToStdSingleThreaded(bytesCount int) {
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

func generateOutputToFileMultiThreaded(fileName string, bytesCount, coresCount, threadsCount int) {
	generateOutputToFileSingleThreaded(fileName, bytesCount)
}

func generateOutputToFileSingleThreaded(fileName string, bytesCount int) {
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
