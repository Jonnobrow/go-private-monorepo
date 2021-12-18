package lib1

import "fmt"

// OutputName prints the provided name to stdout in the following format
// Your name is %s
func OutputName(name string) {
    outputX("name", name)
}

// OutputJob prints the provided job to stdout in the following format
// Your job is %s
func OutputJob(job string) {
    outputX("job", job)
}

func outputX(thing string, value string) {
    fmt.Printf("Your %s is %s\n", thing, value)
}
