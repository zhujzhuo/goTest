package main 

import  (
"fmt"
)

func IAppend(s1,s2 []byte) []byte{
   s1_len := len(s1)
   s2_len := len(s2)
   total_len := len(s1)+len(s2)
   fmt.Println(s1_len,s2_len)
   s :=make([]byte,total_len)
   for k,v:=range s1 {
     s[k]=v
   }
   for k,v:=range s2 {
    s[s1_len+k]=v
   }
   return s
}

func main(){
  s1 := []byte{1,2,3,'n'}
  s2 := []byte{4,5,'a'}
  s  := IAppend(s1,s2)
  fmt.Println(s)
}
