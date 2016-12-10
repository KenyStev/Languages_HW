from bottle import route, run, get, post, request, static_file
import services as s
import os

SORTHPATH = "resources/mergesort/"
BITCODEPATH = "resources/bitcode/"

def upload(req, path):
    file = req.files.get('file')
    folder = createDir(path,file.filename)
    file.save(folder+file.filename)

def download(filename):
    return static_file(filename, root='', download=filename)

def createDir(path,filename):
	name = filename.split(".")[0]
	os.makedirs(path+name)

	return path+name+"/"

@post('/api/sort')
def sortFile():
	upload(request,SORTHPATH)
	return download(s.SortEmails(request.files.file.filename))

@post('/api/bitcode')
def hideMessage():
	upload(request,BITCODEPATH)
	return download(s.HideMessage(request.files.file.filename,request.forms.get('message')))

@post('/api/bitcode/seek')
def hideMessage():
	upload(request,BITCODEPATH)
	return download(s.SeekMessage(request.files.file.filename))

@post('/api/kruskal')
def doKruskal():
	json_graph = request.json
	return s.Kruskal(json_graph)

@post('/upload')
def uploadFile():
    upload(request,"uploads/")

run(host='localhost', port=8080)