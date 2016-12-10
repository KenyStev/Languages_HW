def bytesManipulate():	
	with open("uploads/emails.txt", "r+b") as f:
	    # fourbytes = [ord(b) for b in f.read(4)]
	    # fourbytes[0] = fourbytes[1]  # whatever, manipulate your bytes here
	    # fourbytes = ['k','e','n','y']
	    f.seek(5)
	    f.write("keny")

def test():
	with open("uploads/emails.txt", "rb") as f:
	    byte = f.read(1)
	    while byte != "":
	        print byte
	        byte = f.read(1)


def test2():
	# bitlist = ['AB', 'EC', 'CD', 'AB', 'ED', 'EB', 'DB', 'AB', 'EC']
	# bitstring = ''
	# for bit in bitlist:
	#     bitstring += r'\x' + bit
	bitout = open('uploads/emails.txt', 'ab')
	bitout.seek(0)
	bitout.write("keny")
# test()
# test2()
import os

# os.makedirs("directory")
def readline_by_line():
	with open('uploads/emails.txt','r') as f:
		line = f.readline()
		while len(line) > 0:
			print line
			line = f.readline()

def try_array():
	with open('uploads/emails.txt','r') as f:
		data = f.read()
		print data
		data = sorted(data.split('\n'))
		print data
		data = "\n".join(data)
		data = data[1:]
		print data

def joinArray():
	rootpath = ["mergesort","emails","leaves","sorted"]
	sortedpath = "/".join(rootpath) + "/"
	print sortedpath

# print os.listdir("uploads")

# for x in xrange(0,10):
# 	print x

# with open("uploads/oceano.bmp") as f:
# 	bytes = f.read(5)
# 	for b in range(0,len(bytes)):
# 		print bytes[b]

# f = open("uploads/text.txt")
# cont = 0
# while cont < 5:
# 	b = f.read(1)
# 	print b
# 	cont = cont + 1
# print "fuera"
# f.close()

# numbers = {'first': 2, 'second': 3, 'third': 1, 'four': 4}
# sorted_nus = sorted(numbers,key=numbers.__getitem__)

# print numbers
# print sorted_nus

# for x in sorted_nus:
# 	print x

import json, services

with open("graph.json") as f_data:
	g = json.load(f_data)
	print services.Kruskal(g)