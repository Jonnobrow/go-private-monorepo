package lib1

import "fmt"

// OutputName prints the provided name to stdout in the following format
// Your name is %s
func OutputName(name string) {
    fmt.Printf("Your name is %s", name)
}
