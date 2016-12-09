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
# import os

# os.makedirs("directory")

with open('uploads/emails.txt','r') as f:
	line = f.readline()
	while len(line) > 0:
		print line
		line = f.readline()