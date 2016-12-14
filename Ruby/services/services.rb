require 'json'
require_relative '../merge/merge'
require_relative '../bitcode/bitcode'
require_relative '../graph/kruskal'

def SortEmails(filename)
	puts filename
	name = "#{filename.split(".")[0]}"
	FilterFile("#{name}/#{filename}",/\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z/i)
	CreateLeaves("#{name}/#{filename}.filtered",5)
	MergeSort("#{name}/leaves/")

	GetSortedFile(filename)
end

def HideMessage(filename, message)
	hideMessage(filename,message)
	GetHidden(filename)
end

def SeekMessage(filename)
	seekMessage(filename)
	GetMessage(filename)
end

def Kruskal(json_graph)
	tree = ApplyKruskal(json_graph)
	tree.to_json
end

# SortEmails("2milmails.txt")

# HideMessage("oceano.bmp","que pedo")

# SeekMessage("hidden_oceano.bmp")

# f = File.open("graph/graph.json",'r')
# data = f.read

# puts Kruskal(JSON.parse(data))