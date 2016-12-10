from bitFuncs import *
import binascii
import struct
import math

rootpath = "resources/bitcode/"

def writeBits(fileout, bytes_arr):
	for byte_i in range(0,len(bytes_arr)):
		for i in range(0,8):
			print hasBit(bytes_arr[byte_i],i)
			imgByte = struct.unpack('b',fileout.read(1))[0]
			fileout.seek(-1,1)
			print "imgByte before: ",imgByte
			if hasBit(bytes_arr[byte_i],i):
				imgByte = setBit(imgByte,0)
			else:
				imgByte = clearBit(imgByte,0)
			
			print "imgByte after: ",imgByte
			fileout.write(struct.pack('b',imgByte))

def readBits(filein, bytes_arrlen):
	bytes_arr = []
	for bit in range (0,bytes_arrlen):
		new_byte = 0
		for i in range(0,8):
			fileByte = struct.unpack('b',filein.read(1))[0]
			if hasBit(fileByte,0):
				new_byte = setBit(new_byte,i)
			else:
				new_byte = clearBit(new_byte,i)
			
		bytes_arr.append(new_byte)
	
	return bytes_arr

def getImageOffsetStart(imagename):
	folder = imagename.split(".")[0] + "/"
	with open(rootpath + folder+imagename,'rb') as image:
		image.seek(28)
		bit = image.read(2)
		bitsPerColor = convertBytesToInt(bit,'h')

		print "bitsPerColor: ", bitsPerColor

		numColors = 0
		if bitsPerColor <= 8:
			numColors = math.pow(2,bitsPerColor)
			print "numColors: ",numColors

		headerSize = 54
		colorTableSize = 4*numColors
		imageOffset = headerSize + colorTableSize

		print "imageOffset dec: ",imageOffset,", hex: ",imageOffset
		return imageOffset

def saveMessage(filename,message):
	messageFile = open(filename,'w')
	messageFile.write(message)
	messageFile.close()

def convertBytesToInt(bit,t):
	return struct.unpack(t,bit)[0]

def GetHidden(name):
	return rootpath + name.split(".")[0] + "/hidden_" + name

def GetMessage(name):
	return rootpath + name.split(".")[0] + "/message.txt"

# getImageOffsetStart("oceano.bmp")
# saveMessage("prueba.txt",'Hola')

# file = open("uploads/oceano.bmp",'r+b')
# bytes_arr = struct.pack('I',64)
# print bytes_arr
# file.seek(54)
# writeBits(file,bytearray(bytes_arr))
# file.close()

# file2 = open("uploads/oceano.bmp",'r+b')
# file2.seek(54)
# bytes_arr2 = bytearray(readBits(file2,4))
# print "num: ",struct.unpack('I',bytes_arr2)[0]
# file2.close()