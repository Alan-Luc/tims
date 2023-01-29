package img

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

const (
	scaleDir string = "/home/alan/Pictures/waifu2x/"
)

type img struct {
	FileName string
	FilePath string
}

type Images []img

func (i *Images) Grep(file string) {
	grep := exec.Command("grep", file)
	find := exec.Command("find", "/home/alan/Pictures")

	pipe, err := find.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	grep.Stdin = pipe
	if err := find.Start(); err != nil {
		log.Fatal(err)
	}

	grepOut, _ := grep.Output()
	pipe.Close()

	// this pattern returns filename without extension
	fileMatchPat := regexp.MustCompile(`.*/([^/]+)\.(\w+)$`)
	filePaths := strings.Split(string(grepOut), "\n")

	if len(filePaths) > 0 {
		for _, path := range filePaths {
			img := img{
				FileName: fileMatchPat.ReplaceAllString(path, "$1"),
				FilePath: path,
			}

			*i = append(*i, img)
		}
	}
}

func (i *Images) Scale(ind int) {
	list := *i
	if ind < 0 || ind > len(list) {
		log.Fatal("invalid selection")
	}

	img := list[ind]
	input := img.FileName
	output := fmt.Sprintf("%s%s@2x.png", scaleDir, input)

	cmd := exec.Command("waifu2x-ncnn-vulkan", "-i", input, "-o", output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Images) Convert(ind int, format string) {
	list := *i
	if ind < 0 || ind > len(list) {
		log.Fatal("invalid selection")
	}

	img := list[ind]
	imgPath := img.FilePath

	// this pattern returns entire path without extension
	newPathPat := regexp.MustCompile(`([^/]+)(\.\w+)$`)
	newPath := newPathPat.ReplaceAllString(imgPath, "$1")
	newPath = fmt.Sprintf("%s.%s", newPath, format)

	cmd := exec.Command("convert", imgPath, newPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
