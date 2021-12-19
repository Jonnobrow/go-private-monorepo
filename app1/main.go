package main

import (
	"github.com/jonnobrow/go-private-monorepo/lib1"
	"github.com/jonnobrow/go-private-monorepo/nested/lib2"
)

func main() {
	lib1.OutputName("Jonathan")
	lib1.OutputJob("Associate DevOps Engineer")
	lib2.OutputAge(22)
}
