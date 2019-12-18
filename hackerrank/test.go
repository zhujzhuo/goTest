package main

import (
    "fmt"
    "strings"
    "strconv"
)

// Complete the staircase function below.
func staircase(n int32) {
      var i,j,spaces,stair int32
      for i,j=1,n-1;i<=n;i++{
          result:=""
          for spaces=0;spaces<j;spaces++{
              result+=" "
          }
          for stair=0;stair<i;stair++{
              result+="#"
          }
          fmt.Println(result)
          j-=1
      } 

}

func timeconv(s string) string{
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
    staircase(6)
    fmt.Println("12:05:45AM",timeconv("12:05:45AM"))
    fmt.Println("01:05:45AM",timeconv("01:05:45AM"))
    fmt.Println("12:05:45PM",timeconv("12:05:45PM"))
    fmt.Println("01:05:45PM",timeconv("01:05:45PM"))
}

