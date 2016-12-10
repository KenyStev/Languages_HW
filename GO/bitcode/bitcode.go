package bitcode

import(
	"log"
	"os"; "io"
	"math"
	"strings"
	"encoding/binary"
)

const(
	rootpath = "resources/bitcode/"
)

func HideMessage(imagename,message string) {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()
	saveMessage(rootpath+folder+"message.txt",message)

	imageOut,_ := os.Create(rootpath+folder+"hidden_"+imagename)
	defer imageOut.Close()
	io.Copy(imageOut,image)

	start := getImageOffsetStart(imagename)
	imageOut.Seek(int64(start),0)

	log.Println("-----> len <-----")
	log.Printf("msg len: %d",uint32(len(message)*8))
	messageLen := make([]byte, 4)
    binary.LittleEndian.PutUint32(messageLen, uint32(len(message)))
    log.Println(messageLen)
	writeBits(imageOut,messageLen)
	log.Println("-----> msg <-----")
	messagebits := []byte(message)
	writeBits(imageOut,messagebits)
}

func SeekMessage(imagename string){
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath+folder+imagename)
	defer image.Close()

	start := getImageOffsetStart(imagename)
	image.Seek(int64(start),0)
	log.Println("-----> len <-----")
	messageLenBytes := readBits(image,4)
	log.Println(messageLenBytes)
	messageLen := convertBytesToInt(messageLenBytes)
	log.Printf("msg len: %d",messageLen)
	log.Println("-----> msg <-----")
	message := readBits(image,messageLen)
	log.Printf("msg: %s",string(message))
	
	saveMessage(rootpath+folder+"message.txt",string(message))
}

func writeBits(fileout *os.File, bytes []byte) {
	for _,byte_i := range bytes {
		for i := 0; i < 8; i++ {
			log.Printf("%t",hasBit(byte_i,uint(i)))
			imgByte := make([]byte,1)
			fileout.Read(imgByte)
			fileout.Seek(int64(-1),1)
			log.Printf("imgByte before: %x",imgByte)
			if hasBit(byte_i,uint(i)) {
				imgByte[0] = setBit(imgByte[0],0)
			}else{
				imgByte[0] = clearBit(imgByte[0],0)
			}
			log.Printf("imgByte after: %x",imgByte)
			fileout.Write(imgByte)
		}
	}
}

func readBits(filein *os.File, byteslen uint64) []byte {
	bytes := make([]byte,byteslen)
	for bit := uint64(0); bit<byteslen; bit++ {
		var new_byte byte
		for i := uint(0); i < 8; i++ {
			fileByte := make([]byte,1)
			filein.Read(fileByte)
			if hasBit(fileByte[0],0) {
				new_byte = setBit(new_byte,i)
			}else{
				new_byte = clearBit(new_byte,i)
			}
		}
		bytes[bit] = new_byte
	}
	return bytes
}

func getImageSize(imagename string) uint64 {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()
	
	image.Seek(34,0)
	bit := make([]byte,4)
	image.Read(bit)

	num := convertBytesToInt(bit)

	return num
}

func getImageOffsetStart(imagename string) uint64 {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()
	
	image.Seek(28,0)
	bit := make([]byte,2)
	image.Read(bit)
	bitsPerColor := convertBytesToInt(bit)
	log.Printf("bitsPerColor: %d",bitsPerColor)

	var numColors uint64
	if bitsPerColor <= 8 {
		numColors = uint64(math.Pow(2,float64(bitsPerColor)))
		log.Printf("numColors: %d",numColors)
	}

	headerSize := uint64(54)
	colorTableSize := 4*numColors
	imageOffset := headerSize + colorTableSize

	log.Printf("imageOffset dec: %d, hex: %x",imageOffset,imageOffset)
	return imageOffset
}

func saveMessage(filename,message string) {
	messageFile,_ := os.Create(filename)
	defer messageFile.Close()
	messageFile.WriteString(string(message))
}

func convertBytesToInt(bit []byte) uint64{
	bits := make([]byte,8)
	copy(bits,bit)

	num := binary.LittleEndian.Uint64(bits)
	log.Printf("%d",num)
	return num
}

func GetHidden(name string) string{
	return rootpath + strings.Split(name,".")[0] + "/hidden_" + name
}

func GetMessage(name string) string{
	return rootpath + strings.Split(name,".")[0] + "/message.txt"
}

func GET(imagename string) uint64 {
	return getImageOffsetStart(imagename)
}