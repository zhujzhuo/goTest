package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    /*
     * Write your code here.
     */
//分四种情况
//12:05:45AM  00:05:45
//01:05:45AM  01:05:45
//12:01:20PM  12:01:20
//01:01:20PM  13:01:20
     if strings.Contains(s, "AM") && strings.Split(s,":")[0]=="12"{
        return "00:"+strings.Split(s,":")[1]+":"+strings.Replace(strings.Split(s,":")[2],"AM","",1)
     }else if strings.Contains(s, "PM") && strings.Split(s,":")[0]=="12" {
        return strings.Replace(s,"PM","",1)
     }else if strings.Contains(s, "AM"){
        return strings.Replace(s,"AM","",1)
     }else{
        arItemTemp, _ := strconv.Atoi(strings.Split(s,":")[0])
        arItem := int(arItemTemp)+12
        str := strconv.Itoa(arItem)
        return str+":"+strings.Split(s,":")[1]+":"+strings.Replace(strings.Split(s,":")[2],"PM","",1)
      }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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

