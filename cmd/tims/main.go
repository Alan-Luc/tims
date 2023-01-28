package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	path := "bocchi"

	grep := exec.Command("grep", path)
	find := exec.Command("find", "/home/alan/Pictures")

	pipe, _ := find.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe

	find.Start()

	grepOut, _ := grep.Output()

	arr := strings.Split(string(grepOut), "\n")
	// re := regexp.MustCompile(`[^\/]*\.(\w+)`)
	re := regexp.MustCompile(`.*/([^/]+)\.(\w+)$`)
	pathArr := []string{}
	nameArr := []string{}

	for _, path := range arr {
		pathArr = append(pathArr, path)
		match := re.ReplaceAllString(path, "$1")
		nameArr = append(nameArr, match)
	}
	// format := ".jpg"

	// test := fmt.Sprintf("%s%s", nameArr[1], format)

	fmt.Println(string(grepOut))
	fmt.Println(nameArr)
	fmt.Println(pathArr)

	// cmd := exec.Command("convert", pathArr[1], test)
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
