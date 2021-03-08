package main

import (
	"fmt"
	"os/exec"
)

func main() {
	app := "/home/darren/go/bin/pdfparser"

	arg0 := "-p"
	arg1 := ""
	arg2 := "dataset/27213833-01112021.pdf"
	arg3 := "output"
	arg4 := "-f"

	cmd := exec.Command(app, arg0, arg1, arg4, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Failed")
		fmt.Println(err.Error())
		return
	}

	fmt.Println()
	fmt.Print(string(stdout))
}
