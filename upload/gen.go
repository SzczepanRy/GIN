package main

import (
	"fmt"
	"os"
)

func gen() {

	c("tests", "test")
	c("users", "user")
	c("db", "log")
}
func c(name string, val string) {
	f, err := os.Create(fmt.Sprint(name, ".txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10000; i++ {
		_, err := f.WriteString(fmt.Sprint(i, "  ", val, i, "\n"))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
