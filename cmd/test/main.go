package main

import (
	"fmt"

	"github.com/nachocano/bitbucket-source/pkg/reconciler"
)

func main() {
	hook, err := reconciler.CallHook("key", "test", "test", "test", []string{"push"})
	fmt.Println("test %v, %v", hook, err)
}
