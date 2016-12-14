require 'fileutils'
require_relative 'bitFuncs'

$rootpath = "resources/bitcode/"

def HideMessage(imagename,message)
	folder = "#{imagename.split(".")[0]}/"
	saveMessage("#{$rootpath}#{folder}message.txt",message)

	FileUtils.cp("#{$rootpath}#{folder}#{imagename}","#{$rootpath}#{folder}hidden_#{imagename}")
	imageOut = File.open("#{$rootpath}#{folder}hidden_#{imagename}",'r+b')

	start = getImageOffsetStart(imagename)
	imageOut.seek(start)

	puts "-----> len <-----"
	print "msg len: ",(message.length*8),"\n"
	messageLen = intToArr(message.length)

	puts messageLen
	writeBits(imageOut,messageLen)
	puts "-----> msg <-----"
	messagebits = message.bytes
	writeBits(imageOut,messagebits)

	imageOut.close
end

def SeekMessage(imagename)
	folder ="#{imagename.split(".")[0]}/"
	image = File.open("#{$rootpath}#{folder}#{imagename}")

	start = getImageOffsetStart(imagename)
	image.seek(start)
	puts "-----> len <-----"
	messageLenBytes = readBits(image,4)
	puts messageLenBytes
	messageLen = arrToInt(messageLenBytes)
	print "msg len: ",messageLen,"\n"
	puts "-----> msg <-----"
	
	message = readBits(image,messageLen).map { |e| e.chr }.join
	print "msg: ",message,"\n"
	
	saveMessage("#{$rootpath}#{folder}message.txt",message)
	image.close
end

def writeBits(fileout, bytes_arr)
	print bytes_arr,"\n"
	for byte_i in (0..(bytes_arr.length-1))
		for i in 0..7
			byte_Arr = bytes_arr[byte_i]
			puts hasBit(byte_Arr,i)
			imgByte = fileout.getbyte
			fileout.seek(-1,IO::SEEK_CUR)
			print "imgByte before: ",imgByte,"\n"
			if hasBit(byte_Arr,i)
				imgByte = setBit(imgByte,0)
			else
				imgByte = clearBit(imgByte,0)
			end
			
			print "imgByte after: ",imgByte,"\n"
			fileout << imgByte.chr
		end
	end
end

def readBits(filein, bytes_arrlen)
	bytes_arr = []
	for bit in (0..(bytes_arrlen-1))
		new_byte = 0
		for i in (0..7)
			fileByte = filein.getbyte
			puts hasBit(fileByte,0)
			print "fileByte: ",fileByte,"\n"
			print "new_byte before: ",new_byte,"\n"
			if hasBit(fileByte,0)
				new_byte = setBit(new_byte,i)
			else
				new_byte = clearBit(new_byte,i)
			end
			print "new_byte after: ",new_byte,"\n"
		end
			
		bytes_arr.push(new_byte)
	end
	print "bytes_arr: ",bytes_arr,"\n"

	return bytes_arr
end

def getImageOffsetStart(imagename)
	folder = "#{imagename.split(".")[0]}/"
	File.open("#{$rootpath}#{folder}#{imagename}",'r') do |image|
		image.seek(28)
		bit = image.read(2)
		bitsPerColor = convertBytesToInt(bit)

		print "bitsPerColor: ", bitsPerColor

		numColors = 0
		if bitsPerColor <= 8
			numColors = 2**bitsPerColor
			print "numColors: ",numColors
		end
		headerSize = 54
		colorTableSize = 4*numColors
		imageOffset = headerSize + colorTableSize

		print "imageOffset dec: ",imageOffset,", hex: ",imageOffset
		imageOffset
	end
end

def saveMessage(filename,message)
	messageFile = File.open(filename,'w')
	messageFile.write(message)
	messageFile.close()
end

def convertBytesToInt(bit)
	bit.unpack('S')[0]
end

def GetHidden(name)
	"#{$rootpath}#{name.split(".")[0]}/hidden_#{name}"
end

def GetMessage(name)
	"#{$rootpath}#{name.split(".")[0]}/message.txt"
end

def intToArr(n)
	bytes = []
	bytes.push(n & 0xFF)
	bytes.push(n >> 8 & 0xFF)
	bytes.push(n >> 16 & 0xFF)
	bytes.push(n >> 24 & 0xFF)
	bytes
end

def arrToInt(arr)
	arr[0] | (arr[1] << 8) | (arr[2] << 16) | (arr[3] << 24)
end

# fo = File.open("#{$rootpath}oceano/oceano.bmp", "r+b")
# fo.seek(54)
# writeBits(fo,intToArr(127))
# fo.seek(54)
# num = readBits(fo,4,54)
# print "num: ",arrToInt(num),"\n"
# fo.close

# leer mensaje
# fo = File.open("#{$rootpath}oceano/oceano.bmp", "r+b")
# fo.seek(54)
# writeBits(fo,"Hola".bytes)
# fo.seek(54)
# num = readBits(fo,4,54)
# print "num: ",num,"\n"
# print "num: ",num.map { |e| e.chr }.join,"\n"
# fo.close

# HideMessage("oceano.bmp","Adiso papaito")
SeekMessage("hidden_oceano.bmp")