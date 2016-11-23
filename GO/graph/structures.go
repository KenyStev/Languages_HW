package graph

type Node struct{
	Value interface{} `json:"value"`
}

type Edge struct{
	From Node 	`json:"from"`
	To Node 	`json:"to"`
	Weight int 	`json:"weight"`
}

type Graph struct{
	Edges []Edge 	`json:"edges"`
}

func (g *Graph) Add(new_Edge Edge) {
	g.Edges = append(g.Edges, new_Edge)
}

func (g Graph) GetEdges() []Edge {
	return g.Edges
}

func (g Graph) GetIndexes() map[Node]int {
	index := 0
	indexedMap := make(map[Node]int)
	for _,Edge_g := range g.GetEdges() {
		indexedMap[Edge_g.From] = 0
		indexedMap[Edge_g.To] = 0
	}
	for i,_ := range indexedMap {
		indexedMap[i] = index
		index++
	}
	return indexedMap
}

func (s Graph) Len() int {
    return len(s.Edges)
}

func (s Graph) Swap(i, j int) {
    s.Edges[i], s.Edges[j] = s.Edges[j], s.Edges[i]
}

func (s Graph) Less(i, j int) bool {
    return s.Edges[i].Weight < s.Edges[j].Weight
}