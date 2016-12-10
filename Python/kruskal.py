import json
from operator import itemgetter

def GetIndexes(graph):
	index = 0
	indexedMap = {}
	print("edges: ",graph["edges"])
	edges = graph["edges"]
	for i in range (0,len(edges)):
		print("edge: ",edges[i])
		edge = edges[i]
		indexedMap[edge["from"]["value"]] = 0
		print("indexFrom: ",edge["from"]["value"])
		print("indexFrom map: ",indexedMap.get(edge["from"]["value"]))
		indexedMap[edge["to"]["value"]] = 0
		print("indexTo: ",edge["to"]["value"])
		print("indexTo map: ",indexedMap.get(edge["to"]["value"]))
	
	for key in indexedMap:
		print("key: ",key, " value: ",indexedMap[key])
		indexedMap[key] = index
		index = index + 1
		
	return indexedMap

def getEndIndex(arr, i):
	index = i
	while arr[index] >= 0:
		index = arr[index]

	return index

def ApplyKruskal(graph):
	new_Graph = []
	indexes = GetIndexes(graph)
	size = len(indexes)
	conections = []
	for i in range(0,size):
		conections.append(-1)

	print("-- ApplyKruskal --")
	print("indexes: ", indexes)
	print("conections: ", conections)
	
	queue_sorted = sorted(graph["edges"],key=itemgetter('weight'))
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