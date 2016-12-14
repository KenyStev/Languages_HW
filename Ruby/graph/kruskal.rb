require 'json'

def GetIndexes(graph)
	index = 0
	indexedMap = Hash.new
	puts "edges: ",graph["edges"]
	edges = graph["edges"]
	for i in (0..edges.length-1)
		puts("edge: ",edges[i])
		edge = edges[i]
		indexedMap[edge["from"]["value"]] = 0
		puts("indexFrom: ",edge["from"]["value"])
		puts("indexFrom map: ",indexedMap[edge["from"]["value"]])
		indexedMap[edge["to"]["value"]] = 0
		puts("indexTo: ",edge["to"]["value"])
		puts("indexTo map: ",indexedMap[edge["to"]["value"]])
	end
	indexedMap.each do |key,value|
		print("key: ",key, " value: ",indexedMap[key],"\n")
		indexedMap[key] = index
		index = index + 1
	end
	indexedMap
end

def getEndIndex(arr, i)
	index = i
	while arr[index] >= 0
		index = arr[index]
	end
	index
end
=begin

def ApplyKruskal(graph)
	new_Graph = []
	indexes = GetIndexes(graph)
	size = indexes.length
	conections = []
	size.times do
		conections.push(-1)
	end

	puts("-- ApplyKruskal --")
	print("indexes: ", indexes,"\n")
	print("conections: ", conections,"\n")
	
	queue_sorted = graph["edges"].sort
	print(indexes)
	print("old_Graph: ",graph)
	print "\n"
	print("sorted: ",queue_sorted)
	for i in range (0,len(queue_sorted)):
		item = queue_sorted[i]
		print item
		ini = getEndIndex(conections,indexes.get(item["from"]["value"]));
		fin = getEndIndex(conections,indexes.get(item["to"]["value"]));
		if (ini != fin):
			new_Graph.append(item)
			if((-conections[ini])>(-conections[fin])):
				conections[ini]+=conections[fin]
				conections[fin]=ini
			else:
				conections[fin]+=conections[ini]
				conections[ini]=fin
		else:
			print("ya estan en el mismo arbol - forma ciclo: ",item)

	print("new_Graph: ",new_Graph)
	return new_Graph
end
=end

f = File.open("graph/graph.json",'r')
data = f.read

# puts GetIndexes(JSON.parse(data))
d = JSON.parse(data)

puts d["edges"].sort_by {|hsh| hsh["weight"]}