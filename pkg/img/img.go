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
)

type Img struct {
	FileName string
	FilePath string
}

// methods needed to implement list.Item interface
func (img Img) Title() string       { return img.FileName }
func (img Img) Description() string { return img.FilePath }
func (i Img) FilterValue() string   { return i.FileName }

type Images []Img

func (i *Images) Find(file string) {
	findImgArg := fmt.Sprintf(
		"-type f \\( -iname \"*%s*.jpg\" -o -iname \"*%s*.png\" -o -iname \"*%s*.jpeg\" \\)",
		file,
		file,
		file,
	)
	searchCmd := fmt.Sprintf("find %s %s %s", searchDir1, searchDir2, findImgArg)

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
	} else {
		fmt.Printf("no matches for \"%v\"\n", file)
		os.Exit(1)
	}
}

func (i *Images) Scale(ind int) {
	imgList := *i
	if ind < 0 || ind > len(imgList) {
		log.Fatal("invalid selection")
	}

	img := imgList[ind]
	output := fmt.Sprintf("%s%s@2x.png", outputDir, img.FileName)

	cmd := exec.Command("waifu2x-ncnn-vulkan", "-i", img.FilePath, "-o", output, "-n", "2")
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
	// newPathPat := regexp.MustCompile(`([^/]+)(\.\w+)$`)
	// newPath := newPathPat.ReplaceAllString(imgPath, "$1")
	// newPath = fmt.Sprintf("%s.%s", newPath, format)

	// cmd := exec.Command("convert", imgPath, newPath)
	cmd := exec.Command("mogrify", "-format", format, imgPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Images) Rotate(ind int, angle string) {
	imgList := *i
	if ind < 0 || ind > len(imgList) {
		log.Fatal("invalid selection")
	}

	img := imgList[ind]
	imgPath := img.FilePath

	cmd := exec.Command("mogrify", "-rotate", angle, imgPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Images) Monochrome(ind int) {
	imgList := *i
	if ind < 0 || ind > len(imgList) {
		log.Fatal("invalid selection")
	}

	img := imgList[ind]
	imgPath := img.FilePath

	// this pattern returns entire path without extension
	newPathPat := regexp.MustCompile(`([^/]+)(\.\w+)$`)
	newPath := newPathPat.ReplaceAllString(imgPath, "$1")
	newPath = fmt.Sprintf("%sBW.png", newPath)

	cmd := exec.Command("convert", imgPath, "-colorspace", "Gray", newPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
