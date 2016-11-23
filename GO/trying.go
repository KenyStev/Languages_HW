package main

import (
	// "./graph"
	"./services"
	"fmt")

func main() {
	myJSON := "{\"edges\":[{\"from\":{\"value\":\"0\"},\"to\":{\"value\":\"2\"},\"weight\":8},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"3\"},\"weight\":6},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"6\"},\"weight\":9},{\"from\":{\"value\":\"2\"},\"to\":{\"value\":\"4\"},\"weight\":5},{\"from\":{\"value\":\"2\"},\"to\":{\"value\":\"6\"},\"weight\":1},{\"from\":{\"value\":\"4\"},\"to\":{\"value\":\"6\"},\"weight\":6},{\"from\":{\"value\":\"5\"},\"to\":{\"value\":\"7\"},\"weight\":6},{\"from\":{\"value\":\"6\"},\"to\":{\"value\":\"7\"},\"weight\":9},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"5\"},\"weight\":3},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"2\"},\"weight\":4}]}"
	// myGraph := graph.JsonToGraph(myJSON)
	// fmt.Printf("trying0:\n%v\n",myGraph)
	// myJson := graph.GraphToJson(*myGraph)
	// fmt.Printf("trying1:\n%v\n",myJson)
	fmt.Println(services.Kruskal(myJSON))
}

/*import (
	"./graph"
	"encoding/json"
	"fmt"
	)

// type Node struct {
// 	Value string
// }

// var m map[string]int

func main() {
	/*g := graph.Graph{[]graph.Edge{graph.Edge{graph.Node{"4"},graph.Node{"5"},8},
					graph.Edge{graph.Node{"4"},graph.Node{"7"},2},
					graph.Edge{graph.Node{"7"},graph.Node{"5"},9}}}*/

	/*g := graph.Graph{[]graph.Edge{graph.Edge{graph.Node{"0"},graph.Node{"2"},8},
					graph.Edge{graph.Node{"1"},graph.Node{"3"},6},
					graph.Edge{graph.Node{"1"},graph.Node{"6"},9},
					graph.Edge{graph.Node{"2"},graph.Node{"4"},5},
					graph.Edge{graph.Node{"2"},graph.Node{"6"},1},
					graph.Edge{graph.Node{"4"},graph.Node{"6"},6},
					graph.Edge{graph.Node{"5"},graph.Node{"7"},6},
					graph.Edge{graph.Node{"6"},graph.Node{"7"},9},
					graph.Edge{graph.Node{"1"},graph.Node{"5"},3},
					graph.Edge{graph.Node{"1"},graph.Node{"2"},4}}}
*/
	// new_g := graph.ApplyKruskal(g)

	/*json_enc,_ := json.Marshal(g)
	// json_enc_ng,_ := json.Marshal(new_g)
	fmt.Println(string(json_enc))*/
	// fmt.Println(string(json_enc_ng))

	// myJSON := []byte("{\"edges\":[{\"from\":{\"value\":\"0\"},\"to\":{\"value\":\"2\"},\"weight\":8},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"3\"},\"weight\":6},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"6\"},\"weight\":9},{\"from\":{\"value\":\"2\"},\"to\":{\"value\":\"4\"},\"weight\":5},{\"from\":{\"value\":\"2\"},\"to\":{\"value\":\"6\"},\"weight\":1},{\"from\":{\"value\":\"4\"},\"to\":{\"value\":\"6\"},\"weight\":6},{\"from\":{\"value\":\"5\"},\"to\":{\"value\":\"7\"},\"weight\":6},{\"from\":{\"value\":\"6\"},\"to\":{\"value\":\"7\"},\"weight\":9},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"5\"},\"weight\":3},{\"from\":{\"value\":\"1\"},\"to\":{\"value\":\"2\"},\"weight\":4}]}")

	/*var stru graph.Graph
	json.Unmarshal(myJSON,&stru)
	fmt.Println(stru)*/

	// m := g.GetIndexes()
	// m = make(map[string]int)
	// m["8"] = 8
	// m["5"] = 5
	// m["6"] = 6
	// m["4"] = 4
	// m["5"] = 5
	// m["5"] = 5
	
	// for i,val := range m {
	// 	fmt.Printf("%s: %d\n",i, val)
	// }
	// fmt.Printf("len: %d\n",len(m))
// }/*

/*import(
	"./bitcode"
)

func main() {
	bitcode.HideMessage("oceano.bmp","KE")
	bitcode.SeekMessage("hidden_oceano.bmp")
}*/

/*import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}*/

/*import(
	// "fmt"
	"./merge"
	
	// "io/ioutil"
	// "log"
)

func main() {
	merge.FilterFile("emails/emails.txt","(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves("emails/emails.txt.filtered",5)
	// merge.SortFile("emails/leaves/leaf0")
	merge.GetLeaves("resources/emails/leaves/")
	merge.MergeSort("emails/leaves/")	
}*/


/*//read file
import (
	"fmt"
	"os"
	"bufio"
	// "regexp"
	// "io/ioutil"
	)

func main() {
	// r,_ := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	// file,_ := ioutil.ReadFile("emails.txt")
	// fmt.Println(r.FindAllString(string(file),-1))

	cont := 0
	pos := 0
	file,_ := os.Open("resources/emails/emails.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if cont > 3 {
			file.Seek(int64(pos),0)
			scanner = bufio.NewScanner(file)
			if cont > 10{
				break
			}
		}else{
			pos += len(line) + 1 
		}
		cont++
		fmt.Println(line)
		fmt.Println(int64(-len(line)))
	}
}*/