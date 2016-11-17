package main

import(
	// "fmt"
	"./merge"
	
	// "io/ioutil"
	// "log"
)

func main() {
	merge.FilterFile("emails/emails.txt","(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves("emails/emails.txt.filtered",5)
	// merge.SortFile("emails/leaves/leaf0")
	// merge.GetLeaves("resources/emails/leaves/")
	merge.MergeSort("emails/leaves/")	
}


/*//read file
import (
	"fmt"
	"os"
	"bufio"
	// "regexp"
	// "io/ioutil"
	)

func main() {
	// r,_ := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	// file,_ := ioutil.ReadFile("emails.txt")
	// fmt.Println(r.FindAllString(string(file),-1))

	cont := 0
	pos := 0
	file,_ := os.Open("resources/emails/emails.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cont > 3 {
			file.Seek(int64(pos),0)
			scanner = bufio.NewScanner(file)
			if cont > 10{
				break
			}
		}else{
			pos += len(line) + 1 
		}
		cont++
		fmt.Println(line)
		fmt.Println(int64(-len(line)))
	}
}*/