package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //This tests whether a pattern matches a string.
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")  // Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch"))  //This finds the match for the regexp.

    fmt.Println(r.FindStringIndex("peach punch"))  //This also finds the first match but returns the start and end indexes for the match instead of the matching text.    [0 5]

    fmt.Println(r.FindStringSubmatch("peach punch")) //  The Submatch variants include information about both the whole-pattern matches and the submatches within those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).     [peach ea]  

    fmt.Println(r.FindStringSubmatchIndex("peach punch")) //Similarly this will return information about the indexes of matches and submatches.   [0 5 1 3]

    fmt.Println(r.FindAllString("peach punch pinch ptt", -1))  //The All variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.   [peach punch pinch]

    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //These All variants are available for the other functions we saw above as well.    [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

    fmt.Println(r.FindAllString("peach punch pinch", 2))  //Providing a non-negative integer as the second argument to these functions will limit the number of matches.  [peach punch]

    fmt.Println(r.Match([]byte("peach")))  

    r = regexp.MustCompile("p([a-z]+)ch")    //When creating constants with regular expressions you can use the MustCompile variation of Compile. A plain Compile won’t work for constants because it has 2 return values.   
    fmt.Println(r)  //p([a-z]+)ch
    fmt.Println(r.FindString("peach punch"))

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //a <fruit>

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper) //The Func variant allows you to transform matched text with a given function.
    fmt.Println(string(out)) //a PEACH
    fmt.Println(string(bytes.ToUpper(in)))  // A PEACH
}
