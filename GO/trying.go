package main

import(
	//"fmt"
	"./merge"
)

func main() {
	merge.CreateFolder("pruebas")
	merge.FilterFile("emails/emails.txt","(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves("emails/emails.txt.filtered",3)
}


/* //read file
import (
	"fmt"
	"regexp"
	"io/ioutil"
	)

func main() {
	r,_ := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	file,_ := ioutil.ReadFile("emails.txt")
	fmt.Println(r.FindAllString(string(file),-1))
}*/