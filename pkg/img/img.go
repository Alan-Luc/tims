package img

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	// os environs
	homeDir   string = os.Getenv("HOME")
	outputDir string = fmt.Sprintf("%s/Pictures/waifu2x/", homeDir)
	shell     string = os.Getenv("SHELL")

	// search vars
	searchDir1 string = fmt.Sprintf("%s%s", homeDir, "/Pictures")
	searchDir2 string = fmt.Sprintf("%s%s", homeDir, "/Dropbox")
	findImgArg string = "-name \"*.jpg\" -o -name \"*.png\" -o -name \"*.jpeg\""
)

type Img struct {
	FileName string
	FilePath string
}

// functions needed to implement list.Item interface
func (img Img) Title() string       { return img.FileName }
func (img Img) Description() string { return img.FilePath }
func (i Img) FilterValue() string   { return i.FileName }

type Images []Img

func (i *Images) Grep(file string) {
	searchCmd := fmt.Sprintf("find %s %s %s | grep %v", searchDir1, searchDir2, findImgArg, file)

	matches, err := exec.Command(shell, "-c", searchCmd).Output()
	if err != nil {
		fmt.Printf("no matches for \"%v\"\n", file)
		os.Exit(1)
	}

	// this pattern returns filename without extension
	fileMatchPat := regexp.MustCompile(`.*/([^/]+)\.(\w+)$`)
	filePaths := strings.Split(string(matches), "\n")

	if len(filePaths) > 1 {
		filePaths = filePaths[:len(filePaths)-1]
		for _, path := range filePaths {
			name := fileMatchPat.ReplaceAllString(path, "$1")
			img := Img{
				FileName: name,
				FilePath: path,
			}
			*i = append(*i, img)
		}
	}
}

func (i *Images) Scale(ind int) {
	imgList := *i
	if ind < 0 || ind > len(imgList) {
		log.Fatal("invalid selection")
	}

	img := imgList[ind]
	output := fmt.Sprintf("%s%s@2x.png", outputDir, img.FileName)

	cmd := exec.Command("waifu2x-ncnn-vulkan", "-i", img.FilePath, "-o", output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Images) Convert(ind int, format string) {
	imgList := *i
	if ind < 0 || ind > len(imgList) {
		log.Fatal("invalid selection")
	}

	img := imgList[ind]
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
