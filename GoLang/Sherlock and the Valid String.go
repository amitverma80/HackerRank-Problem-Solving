package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the isValid function below.
func isValid(s string) string {    
    var sArr = []rune(s)
    var sIntArr = make([]int32, 123)
    var max, min int32 =0, 100000

    if len(s)==2{
        return "YES"
    }

    for _, e:=range sArr{
        sIntArr[e] = sIntArr[e]+1
        if sIntArr[e]>max{
            max = sIntArr[e]
        }
        if sIntArr[e]<min{
            min = sIntArr[e]
        }
    }
    if max==min{//as all characters are of equal numbers
        return "YES"
    }

    //there can be max three different digits to make it no
    var dig1, dig2, dig3, counter int32 = 0,0,0,0

    for counter =97;counter<123;counter++{
        if sIntArr[counter]!=0{
            if dig1 ==0 || dig1 == sIntArr[counter]{
                dig1 = sIntArr[counter]
            }
            if (dig2 ==0 || dig2 == sIntArr[counter]) && sIntArr[counter]!=dig1{
                if dig2 == sIntArr[counter]{
                    return "NO" // as coming 2nd time, so not required
                }
                dig2 = sIntArr[counter]
            }
            if (dig3 ==0 || dig3 == sIntArr[counter]) && sIntArr[counter]!=dig1 && sIntArr[counter]!=dig2{
                if dig3 == sIntArr[counter]{
                    return "NO" // as coming 2nd time, so not required
                }
                dig3 = sIntArr[counter]
            }
        }
    }

    if dig1>0 && dig2>0 && dig3>0{
        return "NO"//as we can not use three characters for reduction. Only one character can be reduced one time
    }
    if dig1>0 && dig2>0 && dig3==0{
        if dig1-dig2 ==1 || dig1-dig2==-1{
            return "YES"
        }
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

    result := isValid(s)

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
