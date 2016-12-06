var HashMap = require('hashmap');

var GetIndexes = function(graph) {
	let index = 0
	var indexedMap = new HashMap()
	console.log("edges: ",graph.edges)
	let edges = graph.edges
	for (var i in edges) {
		console.log("edge: ",edges[i])
		edge = edges[i]
		indexedMap.set(edge.from.value,0)
		console.log("indexFrom: ",edge.from.value)
		console.log("indexFrom map: ",indexedMap.get(edge.from.value))
		indexedMap.set(edge.to.value,0)
		console.log("indexTo: ",edge.to.value)
		console.log("indexTo map: ",indexedMap.get(edge.to.value))
	}
	indexedMap.forEach(function(value,key){
		console.log("key: ",key, " value: ",value)
		indexedMap.set(key,index)
		index++
	})
	return indexedMap
}

var getEndIndex = function(arr, i) {
    for (index = i;arr[index]>=0;) {
        index = arr[index]
    }
	return index;
}

var ApplyKruskal = function(graph) {
	new_Graph = []
	let indexes = GetIndexes(graph)
	let size = indexes.count()
	let conections = []
	for (i = 0; i < size; i++) {
		conections[i] = -1
	}
	console.log("-- ApplyKruskal --")
	console.log("indexes: ", indexes)
	console.log("conections: ", conections)
	// return
	graph.edges.sort((a,b) => {
		if (a.weight < b.weight)
			return -1
		else if (a.weight == b.weight)
			return 0
		else
			return 1
	})
	queue_sorted = graph.edges
	console.log(indexes)
	console.log("old_Graph: %v",graph)
	for (var i in queue_sorted) {
		item = queue_sorted[i]
		let ini = getEndIndex(conections,indexes.get(item.from.value));
        let fin = getEndIndex(conections,indexes.get(item.to.value));
        if (ini != fin) {
            new_Graph.push(item)
            if((-conections[ini])>(-conections[fin])) {
                conections[ini]+=conections[fin];
                conections[fin]=ini;
            }else{
                conections[fin]+=conections[ini];
                conections[ini]=fin;
            }
        }else{
            console.log("ya estan en el mismo arbol - forma ciclo: %v\n",item)
        }
	}
	console.log("new_Graph: %v",new_Graph)
	return new_Graph
}

exports.ApplyKruskal = ApplyKruskal

var test = function(){
	let data = `
				{
	"edges": [{
		"from": {
			"value": "0"
		},
		"to": {
			"value": "2"
		},
		"weight": 8
	}, {
		"from": {
			"value": "1"
		},
		"to": {
			"value": "3"
		},
		"weight": 6
	}, {
		"from": {
			"value": "1"
		},
		"to": {
			"value": "6"
		},
		"weight": 9
	}, {
		"from": {
			"value": "2"
		},
		"to": {
			"value": "4"
		},
		"weight": 5
	}, {
		"from": {
			"value": "2"
		},
		"to": {
			"value": "6"
		},
		"weight": 1
	}, {
		"from": {
			"value": "4"
		},
		"to": {
			"value": "6"
		},
		"weight": 6
	}, {
		"from": {
			"value": "5"
		},
		"to": {
			"value": "7"
		},
		"weight": 6
	}, {
		"from": {
			"value": "6"
		},
		"to": {
			"value": "7"
		},
		"weight": 9
	}, {
		"from": {
			"value": "1"
		},
		"to": {
			"value": "5"
		},
		"weight": 3
	}, {
		"from": {
			"value": "1"
		},
		"to": {
			"value": "2"
		},
		"weight": 4
	}]
}
				`
	let graph = JSON.parse(data)
	ApplyKruskal(graph)
	// let indexes = GetIndexes(graph)
	// console.log(indexes)
	// console.log(getEndIndex(indexes,5))
	// console.log(indexes)
}
// test()