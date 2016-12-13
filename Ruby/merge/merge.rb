$rootpath = "resources/mergesort/"

def GetLeaves(dirpath)
	files_name = Dir.entries("#{$rootpath}#{dirpath}")
	
	puts "#{dirpath}"
	files_name.delete(".")
	files_name.delete("..")
	puts files_name
	files_name
end

def MergeSort(dirpath)
	leaves = GetLeaves(dirpath)
	sortedpath = "#{dirpath}sorted/"
	createFolder(sortedpath)
	for i in (0..(leaves.length-1))
		SortFile("#{dirpath}#{leaves[i]}")
	end

	leaves = GetLeaves(sortedpath)
	mergesort(sortedpath,leaves,0)
end

def mergesort(path,m,cont)
	if m.length <= 1
		return m
	end

	mid = m.length / 2
	left = m[0,mid]
	right = m[mid,m.length]

	left = mergesort(path,left,cont+1)
	right = mergesort(path,right,cont+2)

	return merge(path,left, right,cont)
end

def merge(path,left, right, cont)
	result = []
	name = "#{cont}.merged.sorted"
	unless cont == 0
		name = "#{cont}_#{(left[0]).split(".")[0]}_#{(right[0]).split(".")[0]}.merged.sorted"
	end
	result.push(name)
	left_file = File.open("#{$rootpath}#{path}#{left[0]}",'r')
	right_file = File.open("#{$rootpath}#{path}#{right[0]}",'r')
	new_file = File.open("#{$rootpath}#{path}#{name}",'w')

	pos_left = 0
	pos_right = 0

	left_word = nil
	right_word = nil
	while true
		if left_word == nil
			left_word = left_file.gets
		end
		if right_word == nil
			right_word = right_file.gets
		end

		if left_word != nil and right_word != nil
			puts "left: #{left_word} #{-left_word.length}"
			puts " right: #{right_word} #{-right_word.length}"
			if left_word <= right_word
				new_file.write(left_word)
				left_word = nil
			else
				new_file.write(right_word)
				right_word = nil
			end
		elsif left_word != nil
			puts "left: #{left_word} #{-left_word.length}"
			new_file.write(left_word)
			left_word = nil
		elsif right_word != nil
			puts " right: #{right_word} #{-right_word.length}"
			new_file.write(right_word)
			right_word = nil
		else
			break
		end
	end

	left_file.close()
	right_file.close()
	new_file.close()

	return result
end

def createFolder(filename)
	folder = filename.split(".")[0]
	Dir.mkdir("#{$rootpath}#{folder}")
end

def FilterFile(filepath, pattern)
	createdFile = "#{$rootpath}#{filepath}.filtered"
	filteredFile = File.open(createdFile,'w')
	
	f = File.open("#{$rootpath}#{filepath}",'r')
	while(line = f.gets)
		puts "entro"
		puts line
		begin
			if(line =~ pattern)
				filteredFile.write("#{line.downcase}\n")
				puts line
				puts "dentro if"
			elsif
				unless line.nil?
					raise line
				end
			end
		rescue
			line = line[0,line.length-1]
			retry
		end
		puts "paso"
	end
	f.close
	filteredFile.close
end

def CreateLeaves(filepath, leafSize)
	folder = "#{$rootpath}#{filepath.split("/")[0]}/leaves/"
	createFolder("#{filepath.split("/")[0]}/leaves/")
	file = File.open("#{$rootpath}#{filepath}",'r')
	
	leafcont = 0
	linecont = 0
	filteredFile = ""

	while(line = file.gets)
		if(linecont%leafSize == 0)
			if(linecont > 0)
				filteredFile.close
			end
			filteredFile = File.open("#{folder}leaf#{leafcont}",'a')
			leafcont += 1
		end
		filteredFile.write(line)
		linecont = linecont + 1
	end

	file.close
	filteredFile.close
end

def SortFile(filepath)
	File.open("#{$rootpath}#{filepath}",'r') do |f|
		fileData = f.readlines
		data = fileData.sort

		print data
		
		puts "filepath:", filepath
		fullpath = filepath.split("/"); 
		puts "fullpath:", fullpath
		fullpath.push("sorted")
		l = fullpath.length
		fullpath[l-2],fullpath[l-1] = fullpath[l-1], fullpath[l-2]
		sortedpath = fullpath.join("/")
		
		puts "path: "+ sortedpath
		File.open("#{$rootpath}#{sortedpath}.sorted",'w') do |sortedFile|
			for i in 0..(data.length-1)
				sortedFile.write(data[i])
			end
		end
	end
end

def GetSortedFile(name)
	"#{$rootpath}#{name.split('.')[0]}/leaves/sorted/0.merged.sorted"
end

# FilterFile('emails/emails.txt',/\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z/i)
# CreateLeaves('emails/emails.txt.filtered',5)
# GetLeaves("emails/leaves")
# SortFile("emails/leaves/leaf0")
# MergeSort("emails/leaves/")