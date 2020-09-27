package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the gameOfThrones function below.
func gameOfThrones(s string) string {

var sArr = []rune(s)
    var sIntArr = make([]int32, 126)
    var oddCount int32 =0

    //read all letters
    for _, e:=range sArr{
        sIntArr[e] = sIntArr[e]+1
    }

    var i int32=0
    for i=97;i<123;i++ {
        if sIntArr[i]%2 !=0{
            oddCount++
        }
    }
    if oddCount>1{
        return "NO"
    }
    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := gameOfThrones(s)

    fmt.Fprintf(writer, "%s\n", result)

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
