import os, re

rootpath = "resources/mergesort/"

def createFolder(filename):
	folder = filename.split(".")[0];
	os.makedirs(rootpath+folder)

def FilterFile(filepath, pattern):
	# file = open(filepath,'r')
	# defer file.Close()
	# scanner := bufio.NewScanner(file)
	createdFile = rootpath +filepath + ".filtered"
	filteredFile = open(createdFile,'a')
	# defer filteredFile.Close()
	# if err != nil{
	# 	log.Println("no se pudo crear: "+createdFile)
	# }
	with open(rootpath+filepath,'r') as f:
		line = f.readline()
		while len(line) > 0:
			if re.match(pattern,line):
				filteredFile.write(line.lower())
			line = f.readline()

	filteredFile.close()


FilterFile("emails/emails.txt","(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")