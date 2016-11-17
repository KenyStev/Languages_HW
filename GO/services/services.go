package services

import(
	"strings"
	"os"; "io"
    "strconv"
	"net/http"
	"log"
	"mime/multipart"
	"../merge"
)

func Upload(name *multipart.FileHeader, savepath string) (int, string){
	merge.CreateFolder(name.Filename)

	log.Println("getting handle to file")
    file, err := name.Open()
    
    defer file.Close()
    if err != nil {
        return http.StatusInternalServerError, err.Error()
    }
    filename := strings.Split(name.Filename,".")[0]
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

func SortEmails(filename string) string{
	name := strings.Split(filename,".")[0]
	merge.FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves(name+"/"+filename+".filtered",5)
	merge.MergeSort(name+"/leaves/")

	return merge.GetSortedFile(filename)
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