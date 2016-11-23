package services

import(
	"strings"
	"../merge"
	"../bitcode"
	"../graph"
)

func SortEmails(filename string) string{
	name := strings.Split(filename,".")[0]
	merge.FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves(name+"/"+filename+".filtered",5)
	merge.MergeSort(name+"/leaves/")

	return merge.GetSortedFile(filename)
}

func HideMessage(filename, message string) string {
	bitcode.HideMessage(filename,message)
	return bitcode.GetHidden(filename)
}

func SeekMessage(filename string) string {
	bitcode.SeekMessage(filename)
	return bitcode.GetMessage(filename)
}

func Kruskal(json_graph *graph.Graph) string {
	// new_g := graph.JsonToGraph([]byte(json_graph))
	// log.Printf("json_graph:\n%v",new_g)
	tree := graph.ApplyKruskal(*json_graph)
	return graph.GraphToJson(*tree)
}