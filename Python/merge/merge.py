import os, re

rootpath = "resources/mergesort/"

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
		data = "\n".join(data)
		# data = data[1:] #delete '\n'
		
		print "filepath:", filepath
		fullpath = filepath.split("/"); 
		print "fullpath:", fullpath
		fullpath.append("sorted")
		l = len(fullpath)
		fullpath[l-2],fullpath[l-1] = fullpath[l-1], fullpath[l-2]
		sortedpath = "/".join(fullpath)
		
		print "path: "+ sortedpath
		# createFolder(sortedpath)
		with open(rootpath + sortedpath + ".sorted",'a') as sortedFile:			
			for i in range(1,len(data)):
				sortedFile.write(data[i])

# FilterFile("emails/emails.txt","(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
# CreateLeaves("emails/emails.txt.filtered",5)
SortFile("emails/leaves/leaf0")