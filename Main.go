package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
	"strings"
	"strconv"
)

func readLines(filename string) ([]string, error) {
    var lines []string
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    buf := bytes.NewBuffer(file)
    for {
        line, err := buf.ReadString('\n')
		
        if len(line) == 0 {
            if err != nil {
                if err == io.EOF {
                    break
                }
                return lines, err
            }
        }
        lines = append(lines, line)
        if err != nil && err != io.EOF {
            return lines, err
        }
    }
    return lines, nil
}

func main() {
    // a flat file that has 339276 lines of text in it for a size of 62.1 MB
    filename := "track.txt"
    lines, err := readLines(filename)
	for index,element := range lines {
		fmt.Println(index,element)
		result := strings.Split(element, " ")
		var startingNumber,endingNumber,transferCode int
		startingNumber, err = strconv.Atoi(result[0])
		fmt.Println(startingNumber)
		endingNumber, err = strconv.Atoi(result[1])
		fmt.Println(endingNumber)
		var runes []rune= []rune(result[2])
		var statusCode rune = runes[0]
		fmt.Println(statusCode)
		transferCode, err = strconv.Atoi(result[3])
		fmt.Println(transferCode)
	}
    fmt.Println(len(lines))
    if err != nil {
        fmt.Println(err)
        return
    }
}