package graph

import(
	"encoding/json"
	"sort"
	"log"
)

func ApplyKruskal(g Graph) *Graph {
	new_Graph := new(Graph)
	indexes := g.GetIndexes()
	size := len(indexes)
	conections := make([]int,size)
	for i := 0; i < size; i++ {
		conections[i] = -1
	}
	sort.Sort(g)
	queue_sorted := g.GetEdges()
	log.Println(indexes)
	log.Printf("old_Graph: %v",g)
	for _,item := range queue_sorted {
		ini := getEndIndex(conections,indexes[item.From]);
        fin := getEndIndex(conections,indexes[item.To]);
        if ini != fin {
            new_Graph.Add(item)
            if((-conections[ini])>(-conections[fin])) {
                conections[ini]+=conections[fin];
                conections[fin]=ini;
            }else{
                conections[fin]+=conections[ini];
                conections[ini]=fin;
            }
        }else{
            log.Printf("ya estan en el mismo arbol - forma ciclo: %v\n",item)
        }
	}
	log.Printf("new_Graph: %v",new_Graph)
	return new_Graph
}

func getEndIndex(arr []int, i int) int {
	index := i
    for (arr[index]>=0) {
        index = arr[index]
    }
	return index;
}

func GraphToJson(g Graph) string {
	json_format,_ := json.Marshal(g)
	return string(json_format)
}

func JsonToGraph(json_format []byte) *Graph {
	var decoded *Graph
	json.Unmarshal(json_format,&decoded)
	return decoded
}