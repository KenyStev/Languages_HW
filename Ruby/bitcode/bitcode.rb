$rootpath = "resources/bitcode/"

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

def GetHidden(name):
	"#{$rootpath}#{name.split(".")[0]}/hidden_#{name}"
end

def GetMessage(name):
	"#{$rootpath}#{name.split(".")[0]}/message.txt"
end