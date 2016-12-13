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

def read_bytes
	f = File.open("uploads/oceano.bmp", "r")
	f.seek(10)
	num = f.read(2).unpack('S')
	puts num
	puts (num[0] + 3)
	f.close
end

read_bytes
