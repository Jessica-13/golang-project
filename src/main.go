package main

import hp "container/heap" 

import (
	"fmt"
	//"math/rand"
)

// START heap definition

/* Definition of "path" structure :
non-default type consisting of an integer (named value) and a list of strings (named nodes)
*/
type path struct {
    value int
    nodes []string
}

/* Definition of "minPath" type :
non-default type consisting of an anonymous list of path 
*/
type minPath []path

/* Function closures
Go language provides a special feature known as an anonymous function. 
An anonymous function can form a closure. 
A closure is a special type of anonymous function that references variables 
declared outside of the function itself. 
It is similar to accessing global variables which are available 
before the declaration of the function.

Explanation: 
The closure references the variable GFG even after the newCounter() function 
has finished running but no other code outside of the newCounter() function 
has access to this variable.
*/

/* Definition anonymous fonction that : 
input  -> h (minPath type) 
-> call to the Len() function
output -> length of h (int type) 
*/
func (h minPath) Len() int { 
	return len(h) 
}

/* Definition anonymous fonction that : 
input  -> h (minPath type)
-> call to the Less() function 
output -> the shorter distance between the two provided
  return True -> h[i].value is the shorter one 
  return False -> h[j].value is the shorter one

h[i] -> case at index i of list h
.valure -> access to attribute value of the case h[]
*/
func (h minPath) Less(i, j int) bool { 
	return h[i].value < h[j].value 
}

/* Definition anonymous fonction that : 
input  -> h (minPath type)
output -> If the second h[j] is the shorter one, 
          inversion of the two to be able to compare it with the next value
*/
func (h minPath) Swap(i, j int)      { 
	h[i], h[j] = h[j], h[i] 
}

/*
*/
func (h *minPath) Push(x interface{}) {
    *h = append(*h, x.(path))
}

/*
*/
func (h *minPath) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

/* Definition of "heap" structure :
non-default type consisting of a pointer of minpath (named values)
*/
type heap struct {
    values *minPath
}

/*
*/
func newHeap() *heap {
    return &heap{values: &minPath{}}
}

/*
*/
func (h *heap) push(p path) {
    hp.Push(h.values, p)
}

/*
*/
func (h *heap) pop() path {
    i := hp.Pop(h.values)
    return i.(path)
}

// END heap definition

// START graphe definition

/* Definition of "edge" structure :
non-default type consisting of a string (named node) and an integer (named weight)
*/
type edge struct {
    node   string
    weight int
}

/*
*/
type graph struct {
    nodes map[string][]edge
}

/*
*/
func newGraph() *graph {
    return &graph{nodes: make(map[string][]edge)}
}

/*
*/
func (g *graph) addEdge(origin, destiny string, weight int) {
    g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
    g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

/*
*/
func (g *graph) getEdges(node string) []edge {
    return g.nodes[node]
}

/*
*/
func (g *graph) getPath(origin, destiny string) (int, []string) { //gets the shortest path between origin and destiny
    h := newHeap()
    h.push(path{value: 0, nodes: []string{origin}})
    visited := make(map[string]bool)

    for len(*h.values) > 0 {
        // Find the nearest yet to visit node
        p := h.pop()
        node := p.nodes[len(p.nodes)-1]

        if visited[node] {
            continue
        }

        if node == destiny {
            return p.value, p.nodes
        }

        for _, e := range g.getEdges(node) {
            if !visited[e.node] {
                // We calculate the total spent so far plus the cost and the path of getting here
                h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
            }
        }

        visited[node] = true
    }

    return 0, nil
}

// END graph definition


/* START OUTPUT definition

type line struct {
	id string
	predecessor string
	distance int
  }

func (l *line) setId(id string) {
	l.id = id
  }
  
func (l *line) setPredecessor(predecessor string) {
	l.predecessor = predecessor
  }  

func (l *line) setDistance(distance int) {
	l.distance = distance
  } 

func (l *line) describe() {
	fmt.Printf("  %v -      %v      -    %v \n", l.id, l.predecessor, l.distance)
  }

 END SECTION OUTPUT
 */
  

func main() {
	/*test values
	ex_val_id := [5]string{"A", "B", "C", "G", "F"}
	ex_val_predecessor := [5]string{"R", "C", "D", "C", "C"}
	ex_val_distance := [5]int{1, 3, 2, 3, 3}
	
	fmt.Println("_____________________________")
	fmt.Println(" id - predecessor - distance ")
	fmt.Println("_____________________________")
	for i := 0; i < 5; i++ {
		ll := &line{id: ex_val_id[i], predecessor: ex_val_predecessor[i], distance: ex_val_distance[i]}
		ll.describe()
	}*/

	fmt.Println("Dijkstra")
    // Example
    graph := newGraph() 

    graph.addEdge("S", "B", 4)
    graph.addEdge("S", "C", 2)
    graph.addEdge("B", "C", 1)
    graph.addEdge("B", "D", 5)
    
	/*
	graph.addEdge("C", "D", 8)
    graph.addEdge("C", "E", 10)
    graph.addEdge("D", "E", 2)
    graph.addEdge("D", "T", 6)
    graph.addEdge("E", "T", 2)
    fmt.Println(graph.getPath("S", "T"))
	*/
	fmt.Println(graph.getPath("S", "D"))
	
/* BIGGER GRAPH : to see later
	min := 1
	max := 20
	for i := 0; i < 20; i++ {
		dist = rand.Intn(max - min) + min)
		graph.addEdge("S", "B", dist)
	}	
*/

  }