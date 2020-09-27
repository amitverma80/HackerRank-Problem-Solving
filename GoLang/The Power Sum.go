package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the powerSum function below.
func powerSum(X int32, N int32, numb int32) int32 {
        var pwr = int32(math.Pow(float64(numb), float64(N)))

    if pwr<X{
        return powerSum(X, N, numb+1) + powerSum(X-pwr, N, numb+1)
    }

    if pwr==X{
        return 1
    } else {
        return 0
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    XTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    X := int32(XTemp)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)

    result := powerSum(X, N, 1)

    fmt.Fprintf(writer, "%d\n", result)

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
