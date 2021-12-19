package lib1

import "fmt"

// OutputName prints the provided name to stdout in the following format
// Your name is %s
func OutputName(name string) {
	fmt.Println(outputX("name", name))
}

// OutputJob prints the provided job to stdout in the following format
// Your job is %s
func OutputJob(job string) {
	fmt.Println(outputX("job", job))
}

func outputX(thing string, value string) string {
	return fmt.Sprintf("Your %s is %s\n", thing, value)
}
