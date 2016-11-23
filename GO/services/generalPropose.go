package services

import (
	"strings"
	"os"; "io"
    "strconv"
	"net/http"
	"log"
	"mime/multipart"
)

func Upload(name *multipart.FileHeader, savepath string) (int, string){
	log.Println("getting handle to file")
    file, err := name.Open()
    
    defer file.Close()
    if err != nil {
        return http.StatusInternalServerError, err.Error()
    }
    filename := strings.Split(name.Filename,".")[0]
    CreateFolder(savepath+filename)
	dst, err := os.Create(savepath+filename+"/" + name.Filename)
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

func Download(filepath,name string,writer http.ResponseWriter) {
	Openfile,_ := os.Open(filepath)
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+name)
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
}

func CreateFolder(folder string) {
	if err := os.Mkdir(folder,0777); err != nil{
		log.Println("no creo el folder: "+folder)
		return
	}
}