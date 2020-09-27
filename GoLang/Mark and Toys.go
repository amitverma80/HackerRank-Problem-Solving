package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/rand"
)

func quickSot(prices []int32)[]int32{
    arrLen := int32(len(prices))

    if arrLen<2{
        return prices
    }

    left, right:= 0, arrLen-1

    pivot:=rand.Int()%len(prices)

    prices[pivot], prices[right]=prices[right], prices[pivot]

    for  index, _:=range prices{
        if prices[index]<prices[right]{
            prices[left], prices[index] = prices[index], prices[left]
            left++
        }
    }

    prices[left], prices[right] = prices[right],prices[left]
    quickSot(prices[:left])
    quickSot(prices[left+1:])

    return prices
}

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
    pricesSort:= quickSot(prices)
    var toyCounter, price  int32=0,0
    for _, e:=range pricesSort{
        if price+e<k{
            price=price+e
            toyCounter++
        }else {
            break
        }
    }
    return toyCounter
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    pricesTemp := strings.Split(readLine(reader), " ")

    var prices []int32

    for i := 0; i < int(n); i++ {
        pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
        checkError(err)
        pricesItem := int32(pricesItemTemp)
        prices = append(prices, pricesItem)
    }

    result := maximumToys(prices, k)

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
