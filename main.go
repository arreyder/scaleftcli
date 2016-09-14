package main

import (
	"fmt"
	"github.com/arreyder/scaleftutil"
	"os"
)

func main() {
	operation := os.Args[1]
	pattern := os.Args[2]
	if operation == "pattern" {
		if pattern == "" {
			fmt.Println("\nYou must supply a pattern for the pattern operation.\n")
		} else {
			err := DeleteByPattern(pattern)
			if err != nil {
				fmt.Println("\nError deleting servers with pattern:%s error:%v\n", pattern, err)
			} else {
				fmt.Println("\nSuccess.\n")
			}
		}
	} else {
		fmt.Println("You must supply an option and right now all we understand is 'pattern' folled by a substring to match hosts to be deleted.  There is no prompting.")
	}
}

func DeleteByPattern(pattern string) error {
	err := scaleftutil.DeleteServersByPattern(pattern)
	if err != nil {
		return fmt.Errorf("\nError deleting servers with pattern:%s error:%v\n", pattern, err)
	} else {
		fmt.Printf("\nDeleted all server with pattern: %s\n", pattern)
		return nil
	}
}
