package merge

import (
	"os"
	"strings"
	"strconv"
	"bufio"
	"fmt"
	"regexp"
)

func sort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = sort(left)
	right = sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func CreateFolder(filename string) {
	folder := strings.Split(filename,".")[0];
	if err := os.Mkdir("resources/"+folder,0777); err != nil{
		fmt.Println("no creo el folder")
		return
	}
}

func FilterFile(filepath, pattern string) {
	file,err := os.Open(filepath)
	defer file.Close()
	if err != nil{
		fmt.Println("no abrio archivo")
		return
	}
	scanner := bufio.NewScanner(file)
	filteredFile,_ := os.Create(filepath + ".filtered")
	defer filteredFile.Close()
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if matched,_ := regexp.MatchString(pattern,line); matched{
			filteredFile.WriteString(line+"\n")
		}
	}
}

func CreateLeaves(filepath string, leafSize int) {
	folder := strings.Split(filepath,"/")[0]
	file,err := os.Open(filepath)
	defer file.Close()
	if err != nil{
		fmt.Println("no abrio archivo")
		return
	}
	leafcont := 0
	scanner := bufio.NewScanner(file)
	var filteredFile *os.File
	var line string
	for linecont := 0; scanner.Scan(); linecont++ {
		if linecont%leafSize == 0 {
			filteredFile,_ = os.Create(folder +"/leaf" + strconv.Itoa(leafcont))
			defer filteredFile.Close()
			leafcont++
		}
		line = scanner.Text()
		filteredFile.WriteString(line+"\n")
	}
}