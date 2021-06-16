package main

import (
"bufio"
"fmt"
"io"
"os"
"strconv"
"strings"
)



/*
 * Complete the 'numberOfItems' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY startIndices
 *  3. INTEGER_ARRAY endIndices
 */

const (
	item = "*"
	compartmentSepparation = "|"
)


func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {
	var result []int32

	for _, startIndex := range startIndices {
		for _, endIndex := range endIndices {
			subString := s[startIndex:endIndex]
			if !strings.Contains(subString, compartmentSepparation) ||
				!strings.Contains(subString, item) ||
				strings.Count(subString, compartmentSepparation) < 2 {
				continue
			}

			result = append(result, getNumberItemsInCompartments(subString))
		}
	}

	return result
}

func getNumberItemsInCompartments(s string) int32 {
	var totalOfItems int32 = 0
	pivotIndex := 0
	indOfCompartment := strings.Index(s, compartmentSepparation)
	for indOfCompartment != -1 {
		pivotIndex += indOfCompartment
		s = s[pivotIndex:]
		pivotIndex += strings.Index(s, compartmentSepparation)
		subString := s[:pivotIndex]
		totalOfItems += int32(strings.Count(subString, item))
		s = s[pivotIndex:]
	}

	return totalOfItems
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	startIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var startIndices []int32

	for i := 0; i < int(startIndicesCount); i++ {
		startIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		startIndicesItem := int32(startIndicesItemTemp)
		startIndices = append(startIndices, startIndicesItem)
	}

	endIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var endIndices []int32

	for i := 0; i < int(endIndicesCount); i++ {
		endIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		endIndicesItem := int32(endIndicesItemTemp)
		endIndices = append(endIndices, endIndicesItem)
	}

	result := numberOfItems(s, startIndices, endIndices)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

