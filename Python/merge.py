import os, re

rootpath = "resources/mergesort/"

def GetLeaves(dirpath):
	files_name = os.listdir(rootpath + dirpath)
	
	print dirpath
	print files_name
	return files_name

def MergeSort(dirpath):
	leaves = GetLeaves(dirpath)
	sortedpath = dirpath + "sorted/"
	createFolder(sortedpath)
	for i in range (0,len(leaves)):
		SortFile(dirpath + leaves[i])

	leaves = GetLeaves(sortedpath)
	mergesort(sortedpath,leaves,0,"")

def mergesort(path,m,cont,path_to):
	if len(m) <= 1:
		return m

	mid = len(m) / 2
	left = m[:mid]
	right = m[mid:]

	left = mergesort(path,left,cont+1,path_to+`cont`)
	right = mergesort(path,right,cont+2,path_to+`cont`)

	return merge(path,left, right,cont,path_to)

def merge(path,left, right, cont,path_to):
	result = []
	# name = `cont` + "_" + (left[0]).split(".")[0] + "_" + (right[0]).split(".")[0] + ".merged.sorted"
	# if cont == 0:
	name = path_to + `cont` + ".merged.sorted"
	result.append(name)
	left_file = openfile(path + left[0],'r')
	right_file = openfile(path + right[0],'r')
	new_file = openfile(path + name,'w')

	pos_left = 0
	pos_right = 0

	left_word = ""
	right_word = ""
	while 1:
		if len(left_word) == 0:
			left_word = left_file.readline()
		if len(right_word) == 0:
			right_word = right_file.readline()
		print "lesf: "+left_word + " " + `-len(left_word)`
		print " right: "+right_word+ " " + `-len(right_word)`

		if len(left_word) > 0 and len(right_word) > 0:
			if left_word <= right_word:
				new_file.write(left_word)
				left_word = ""
			else:
				new_file.write(right_word)
				right_word = ""
			
		elif len(left_word) > 0:
			new_file.write(left_word)
			left_word = ""
		elif len(right_word) > 0:
			new_file.write(right_word)
			right_word = ""
		else:
			break;

	left_file.close()
	right_file.close()
	new_file.close()

	return result

def createFolder(filename):
	folder = filename.split(".")[0];
	os.makedirs(rootpath+folder)

def openfile(filepath,opt):
	return open(rootpath+filepath,opt)

def FilterFile(filepath, pattern):
	createdFile = rootpath +filepath + ".filtered"
	filteredFile = open(createdFile,'a')
	
	with openfile(filepath,'r') as f:
		line = f.readline()
		while len(line) > 0:
			if re.match(pattern,line):
				filteredFile.write(line.lower())
			line = f.readline()

	filteredFile.close()

def CreateLeaves(filepath, leafSize):
	folder = rootpath + filepath.split("/")[0] + "/leaves/"
	createFolder(filepath.split("/")[0] + "/leaves/")
	file = openfile(filepath,'r')
	
	leafcont = 0
	linecont = 0
	filteredFile = ""
	line = file.readline()
	while len(line) > 0:
		if linecont%leafSize == 0:
			if linecont > 0:
				filteredFile.close()
			filteredFile = open(folder +"leaf" + `leafcont`,'a')
			leafcont = leafcont + 1
		
		line = file.readline()
		filteredFile.write(line)
		linecont = linecont + 1

	file.close()
	filteredFile.close()

def SortFile(filepath):
	with open(rootpath + filepath,'r') as f:
		fileData = f.read()
		data = sorted(fileData.split('\n'))
		if len(data[0]) == 0:
			data = data[1:] #delete '\n'
		data = "\n".join(data) + "\n"

		print data
		
		print "filepath:", filepath
		fullpath = filepath.split("/"); 
		print "fullpath:", fullpath
		fullpath.append("sorted")
		l = len(fullpath)
		fullpath[l-2],fullpath[l-1] = fullpath[l-1], fullpath[l-2]
		sortedpath = "/".join(fullpath)
		
		print "path: "+ sortedpath
		with open(rootpath + sortedpath + ".sorted",'w') as sortedFile:			
			for i in range(0,len(data)):
				sortedFile.write(data[i])

def GetSortedFile(name):
	return rootpath + name.split(".")[0] +"/leaves/sorted/0.merged.sorted"