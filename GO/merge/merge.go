package merge

import (
	"os"
	"strings"
	"strconv"
	"bufio"
	"log"
	"regexp"
	"io/ioutil"
	"sort"
)

const(
	rootpath = "resources/mergesort/"
)

func GetLeaves(dirpath string) []string {
	files, err := ioutil.ReadDir(rootpath + dirpath)
	if err != nil {
		log.Println("No pudo abrir el dir: " + dirpath)
	}
	var files_name []string
	for _, file := range files {
		files_name = append(files_name,file.Name())
	}
	log.Println(dirpath)
	log.Println(files_name)
	return files_name
}

func MergeSort(dirpath string) {
	leaves := GetLeaves(dirpath)
	var sortedpath string
	sortedpath = dirpath + "sorted/"
	createFolder(sortedpath)
	for _, file := range leaves {
		SortFile(dirpath + file)
	}
	leaves = GetLeaves(sortedpath)
	mergesort(sortedpath,leaves,0)
}

func mergesort(path string,m []string, cont int) []string{
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = mergesort(path,left,cont+1)
	right = mergesort(path,right,cont+2)

	return merge(path,left, right,cont)
}

func merge(path string,left, right []string, cont int) []string {
	var result []string
	name := strconv.Itoa(cont) + ".merged.sorted"
	result = append(result, name)
	left_file := openfile(path + left[0])
	defer left_file.Close()
	right_file := openfile(path + right[0])
	defer right_file.Close()
	new_file,_ := os.Create(rootpath + path + name)
	defer new_file.Close()

	scan_left := bufio.NewScanner(left_file)
	scan_right := bufio.NewScanner(right_file)
	pos_left := 0; pos_right := 0

	// var left_word,right_word string
	for {
		scan_left.Scan(); scan_right.Scan()
		left_word := scan_left.Text()
		right_word := scan_right.Text()
		log.Print("lesf: "+left_word + " " + strconv.Itoa(-len(left_word)))
		log.Println(" right: "+right_word+ " " + strconv.Itoa(-len(right_word)))

		if len(left_word) > 0 && len(right_word) > 0 {
			if left_word <= right_word {
				new_file.WriteString(left_word+"\n")
				pos_left += len(left_word) + 1
				right_file.Seek(int64(pos_right),0)
				scan_right = bufio.NewScanner(right_file)
			} else {
				new_file.WriteString(right_word+"\n")
				pos_right += len(right_word)+ 1
				left_file.Seek(int64(pos_left),0)
				scan_left = bufio.NewScanner(left_file)
			}
		} else if len(left_word) > 0 {
			new_file.WriteString(left_word+"\n")
		} else if len(right_word) > 0 {
			new_file.WriteString(right_word+"\n")
		}else{
			break;
		}
	}

	return result
}

func createFolder(filename string) {
	folder := strings.Split(filename,".")[0];
	if err := os.Mkdir(rootpath+folder,0777); err != nil{
		log.Println("no creo el folder: "+filename)
		return
	}
}

func openfile(filepath string) *os.File {
	file,err := os.Open(rootpath + filepath)
	
	if err != nil{
		log.Println("no abrio archivo "+filepath)
		return nil
	}
	return file
}

func FilterFile(filepath, pattern string) {
	file := openfile(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	createdFile := rootpath +filepath + ".filtered"
	filteredFile,err := os.Create(createdFile)
	defer filteredFile.Close()
	if err != nil{
		log.Println("no se pudo crear: "+createdFile)
	}
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if matched,_ := regexp.MatchString(pattern,line); matched{
			filteredFile.WriteString(strings.ToLower(line)+"\n")
		}
	}
}

func CreateLeaves(filepath string, leafSize int) {
	folder := rootpath + strings.Split(filepath,"/")[0] + "/leaves/"
	createFolder(strings.Split(filepath,"/")[0] + "/leaves/")
	file := openfile(filepath)
	defer file.Close()
	leafcont := 0
	scanner := bufio.NewScanner(file)
	var filteredFile *os.File
	var line string
	for linecont := 0; scanner.Scan(); linecont++ {
		if linecont%leafSize == 0 {
			filteredFile,_ = os.Create(folder +"leaf" + strconv.Itoa(leafcont))
			defer filteredFile.Close()
			leafcont++
		}
		line = scanner.Text()
		filteredFile.WriteString(line+"\n")
	}
}

func SortFile(filepath string) {
	fileData,_ := ioutil.ReadFile(rootpath + filepath)
	data := strings.Split(string(fileData),"\n")
	sort.Strings(data)
	fullpath := strings.Split(filepath,"/"); 
	fullpath = append(fullpath, "sorted")
	l := len(fullpath)
	fullpath[l-2],fullpath[l-1] = fullpath[l-1], fullpath[l-2]
	sortedpath := rootpath
	for i,path := range fullpath{
		if i<l-1{
			sortedpath += path + "/"
		}
	}
	log.Println("path: "+ sortedpath)
	// createFolder(sortedpath)
	sortedFile,_ := os.Create(sortedpath + fullpath[l-1] + ".sorted")
	defer sortedFile.Close()
	data = data[1:] //delete '\n'
	for _,line := range(data) {
		sortedFile.WriteString(line+"\n")
	}
}

func GetSortedFile(name string) string{
	return rootpath + strings.Split(name,".")[0] +"/leaves/sorted/0.merged.sorted"
}