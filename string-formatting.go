package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("%v\n", p) //this prints an instance of our point struct.  {1 2}

    fmt.Printf("%+v\n", p) //If the value is a struct, the %+v variant will include the struct’s field names. {x:1 y:2}

    fmt.Printf("%#v\n", p) //The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.    main.point{x:1, y:2}

    fmt.Printf("%T\n", p)  // To print the type of a value, use %T.    main.point

    fmt.Printf("%t\n", true)  // Formatting booleans is straight-forward.     true

    fmt.Printf("%d\n", 123) // There are many options for formatting integers. Use %d for standard, base-10 formatting.   123 

    fmt.Printf("%b\n", 14) //This prints a binary representation.   1110

    fmt.Printf("%c\n", 33) // This prints the character corresponding to the given integer.  !

    fmt.Printf("%x\n", 456) //%x provides hex encoding.  1c8

    fmt.Printf("%f\n", 78.9) //There are also several formatting options for floats. For basic decimal formatting use %f.  78.900000

    fmt.Printf("%e\n", 123400000.0) //%e and %E format the float in (slightly different versions of) scientific notation.  1.234000e+08
    fmt.Printf("%E\n", 123400000.0) //1.234000E+08

    fmt.Printf("%s\n", "\"string\"") //For basic string printing use %s.   "string"

    fmt.Printf("%q\n", "\"string\"") //To double-quote strings as in Go source, use %q.  "\"string\""

    fmt.Printf("%x\n", "hex this") //As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.   6865782074686973

    fmt.Printf("%p\n", &p)// To print a representation of a pointer, use %p. 0x42135100

    fmt.Printf("|%6d|%6d|\n", 12, 345)// |    12|   345|

    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)  //|  1.20|  3.45|

    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //|1.20  |3.45  |

    fmt.Printf("|%6s|%6s|\n", "foo", "b") //|   foo|     b|

    fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //|foo   |b     | 

    s := fmt.Sprintf("a %s", "string")  // So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.
    fmt.Println(s) //a string

    fmt.Fprintf(os.Stderr, "an %s\n", "error")  //You can format+print to io.Writers other than os.Stdout using Fprintf.   an error
}
