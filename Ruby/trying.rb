# Read line by line
# myFile=File.open("uploads/emails.txt","r")
# myFile2=File.open("uploads/emails2.txt","r")
# line = myFile.gets
# line2 = myFile2.gets
# state = line || line2
# while(state)
# 	puts line
# 	puts line2
# 	line = myFile.gets
# 	line2 = myFile2.gets
# 	state = line || line2
# end

# myFile.close
# myFile2.close

# Write File
# myFile=File.open("uploads/new_file.txt","w")
# myFile.write("Hola\n")
# myFile.write("Mundo\n")

def openfile(filepath,opt)
	0
end

def test
	puts openfile " "," "
	# openfile("uploads/emails.txt","r") do |myFile|
	# 	puts "entro"
	# 	while(line = myFile.gets)
	# 		puts line
	# 	end
	# 	puts "salio"
	# end
end

# test
# puts Dir.entries("resources")

require_relative 'bitcode/bitFuncs'

def read_bytes
	f = File.open("uploads/oceano.bmp", "r")
	f.seek(10,IO::CUR)
	num = f.read(2).unpack('S')
	puts num
	puts (num[0] + 3)
	f.close
end

# read_bytes
# fo = File.open("uploads/try.b", "wb")
# fo.write([54].pack "I")
# fo.close
=begin

fo = File.open("uploads/try.b", "r+b")
# num = fo.read(4).unpack("S")[0]
# puts num
# num = clearBit(num,0)
# puts num
# fo.seek(-4,IO::SEEK_CUR)
fo.seek(5)
r = [58].pack "I"
fo << r.chr
fo.seek(4)
r = [57].pack "I"
fo << r.chr
fo.close
=end

# 5.times do |i|
# 	puts i
# end
=begin

f = File.new("uploads/try.b",'r+b')   #=> #<File:testfile>
b = f.getbyte              #=> 0x38
puts b
f.ungetbyte(10)             #=> nil
f << 10.chr
f.seek(-1,IO::SEEK_CUR)
puts f.getbyte                  #=> 0x38
=end

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

# print intToArr(1875)
# puts arrToInt(intToArr(1875))
# filteredFile = File.open("resources/bitcode/2milmails/2milmails.txt.filtered",'w')