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

func HideMessages(imagename,message string) {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()

	imageOut,_ := os.Create(rootpath+folder+"hiden_"+imagename)
	defer imageOut.Close()
	io.Copy(imageOut,image)

	start := getImageOffsetStart(imagename)
	imageOut.Seek(int64(start),0)

	messagebits := []byte(message)
	for _,byte_i := range messagebits {
		for i := 0; i < 8; i++ {
			log.Printf("%t",hasBit(byte_i,uint(i)))
			imgByte := make([]byte,1)
			imageOut.Read(imgByte)
			imageOut.Seek(int64(-1),1)
			log.Printf("imgByte before: %x",imgByte)
			if hasBit(byte_i,uint(i)) {
				imgByte[0] = setBit(imgByte[0],0)
			}else{
				imgByte[0] = clearBit(imgByte[0],0)
			}
			log.Printf("imgByte after: %x",imgByte)
			imageOut.Write(imgByte)
		}
	}

	/*image.Seek(10,0)
	bit := make([]byte,1)
	image.Read(bit)
	log.Printf("%x",bit[0])*/

	// image.Seek(int64(bit[0]),0)
	// image.Read(bit)
	// log.Printf("%x",bit[0])
	// image.Seek(int64(-1),1)

	// len_chars := len(message)
}

func getImageSize(imagename string) uint64 {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()
	
	image.Seek(34,0)
	bit := make([]byte,4)
	image.Read(bit)
	bits := make([]byte,8)
	copy(bits,bit)

	num := binary.LittleEndian.Uint64(bits)
	log.Printf("%d",num)

	return num
}

func getImageOffsetStart(imagename string) uint64 {
	folder := strings.Split(imagename,".")[0] + "/"
	image,_ := os.Open(rootpath + folder+imagename)
	defer image.Close()
	
	image.Seek(28,0)
	bit := make([]byte,2)
	image.Read(bit)
	bits := make([]byte,8)
	copy(bits,bit)
	bitsPerColor := binary.LittleEndian.Uint64(bits)
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