package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func staircase(n int32) {
      var i,j,spaces,stair int32
      for i,j=1,n-1;i<=n;i++{
          for spaces=0;spaces<j;spaces++{
              fmt.Print(" ")
          }
          for stair=0;stair<i;stair++{
              fmt.Print("#")
          }
          fmt.Println()
          j-=1
      } 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
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

