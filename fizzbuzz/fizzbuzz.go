package main
import fmt "fmt"

func main() {
    for i := 1; i<=100; i++ {
        switch {
        case i%15 == 0:
            fmt.Printf("FizzBuzz")
        case i%3 == 0:
            fmt.Printf("Fizz")
        case i%5 == 0:
            fmt.Printf("Buzz")
        default:
            fmt.Printf("%d", i)
        }
        fmt.Printf(" ")
    }
    fmt.Printf("\n")
}
