package services

import(
	"strings"
	"os"
	"io"
	"net/http"
	"log"
	"mime/multipart"
	"../merge"
)

func SetupEmails(name *multipart.FileHeader) (int, string){
	merge.CreateFolder(name.Filename)

	log.Println("getting handle to file")
    file, err := name.Open()
    
    defer file.Close()
    if err != nil {
        return http.StatusInternalServerError, err.Error()
    }
    filename := strings.Split(name.Filename,".")[0]
	dst, err := os.Create("resources/mergesort/"+filename+"/" + name.Filename)
	defer dst.Close()
	if err != nil {
	    return http.StatusInternalServerError, err.Error()
	}
	log.Println("file: "+name.Filename)

	log.Println("copying the uploaded file to the destination file")
	if _, err := io.Copy(dst, file); err != nil {
	    return http.StatusInternalServerError, err.Error()
	}
	return 200, "ok"
}

func SortEmails(filename string) string{
	name := strings.Split(filename,".")[0]
	merge.FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves(name+"/"+filename+".filtered",5)
	merge.MergeSort(name+"/leaves/")

	return "ok"
}