// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import "fmt"

func main() {

    // The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // A classic initial/condition/after `for` loop.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` without a condition will loop repeatedly
    // until you `break` out of the loop or `return` from
    // the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // Show leap year
    for n := 1875; n <= 2017; n++ {
        if n%4 == 0 && n%100 != 0 || n%400 == 0 {
            fmt.Printf("%d, ", n)
        } else if n%100 == 0 {
            fmt.Println();
        }
    }
}
