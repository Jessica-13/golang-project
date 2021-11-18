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
-> call the Swap() fonction
output -> If the second h[j] is the shorter one, 
          inversion of the two to be able to compare it with the next value
*/
func (h minPath) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i] 
}

/* Pointer on Golang
The & operator is used to find the memory address of a variable.
The * operator gives access to the values present in the memory address.
*/

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

/* Initialization of a new heap
 return -> values of the heap
 Because :
 *heap <= &heap{values: &minPath{}}
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

/* Map. 
A Go map is a lookup table. 
It returns a value from a key—or tests if one exists. 
(We specify the type of keys and values in a map.)

A map cannot be sorted directly. 
But if we get a slice of the keys from a map, 
we can sort that slice and loop over it, accessing the map's values.

To get an element, we access it by name. 
To loop over the entire map's contents, we use a for-loop—each key 
and value can be accessed separately.
*/

/* Definition of "graph" structure :
non-default type consisting of map like (map[key-type]val-type)
  - key-type -> string
  - val-type -> []edge -> node   string
                          weight int
*/
type graph struct {
    nodes map[string][]edge
}

/* Initialization of a new graph
*/
func newGraph() *graph {
    return &graph{nodes: make(map[string][]edge)}
}

/* Definition anonymous fonction that : 
input  -> g (*graph type)
output -> 
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
    h := newHeap()  // definition of a new heap named h -> call fonction newHeap()
    // fmt.Println("h before first push : ", h) -> &{0xc00000c090}
    // fmt.Println("h before first push : ", *h) -> same
    // fmt.Println("h before first push : ", &h) -> 0xc00000e030
    h.push(path{value: 0, nodes: []string{origin}}) 
    // fmt.Println("h before first push : ", *h) -> same
    // fmt.Println("h after first push : ", h) -> &{0xc00000c090}
    // fmt.Println("h before first push : ", &h) -> 0xc00000e030
    visited := make(map[string]bool)    // Mark all nodes unvisited. 

    // nodes control :
    for len(*h.values) > 0 {
        // Find the nearest yet to visit node
        p := h.pop()
        /* fmt.Println("try p : ", p)
        try p :  {0 [S]}
        try p :  {2 [S C]}
        try p :  {3 [S C B]}
        try p :  {4 [S B]}
        try p :  {8 [S C B D]}
        */
        node := p.nodes[len(p.nodes)-1]
        /* fmt.Println("try node : ", node)
        try node :  S
        try node :  B
        try node :  B
        try node :  D
        */

        if visited[node] {  // if already visited
            continue    // do nothing
        }

        if node == destiny {    // when we reached the destination
            // fmt.Println("try : ", p.value, p.nodes) -> try :  8 [S C B D]
            // p.value = 8
            // p.nodes = [S C B D]
            return p.value, p.nodes
        }

        for _, e := range g.getEdges(node) {
            if !visited[e.node] {   //  if not visited yet
                h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
                // h.push() -> assign the value
                // path{}   -> for the path
                /* 
                We calculate the total spent so far plus the cost and the path of getting here :
                
                (cost) value: p.value + e.weight
                (path) nodes: append([]string{}, append(p.nodes, e.node)...)
                */

                /*
                fmt.Println(" try h : ", h)
                try h :  &{0xc00000c090}
                try h :  &{0xc00000c090}
                try h :  &{0xc00000c090}
                try h :  &{0xc00000c090}
                */

                /*
                fmt.Println(" try path : ", p.value)
                try path :  0
                try path :  0
                try path :  2
                try path :  3
                */

                /*
                fmt.Println(" try path : ", e.weight)
                try path :  4
                try path :  2
                try path :  1
                try path :  5
                */

                /*
                fmt.Println(" try path : ", p.value + e.weight)
                try path :  4
                try path :  2
                try path :  3
                try path :  8
                */

                /*
                fmt.Println(" try path : ", append([]string{}, append(p.nodes, e.node)...))
                try path :  [S B]
                try path :  [S C]
                try path :  [S C B]
                try path :  [S C B D]
                */

                /*
                fmt.Println(" try path : ", append([]string{}))
                try path :  []
                try path :  []
                try path :  []
                try path :  []
                */

                /*
                fmt.Println(" try path : ", append(p.nodes, e.node)...)
                ./main.go:282:28: too many arguments in call to fmt.Println
                have (string, ...string)
                want (...interface {})
                */

                /*
                fmt.Println(" try path : ", append(p.nodes, e.node))
                try path :  [S B]
                try path :  [S C]
                try path :  [S C B]
                try path :  [S C B D]
                */
            }
        }
        visited[node] = true    // once a node is checked, it's marked as visited
    }
    return 0, nil
}

// END graph definition


// START OUTPUT definition

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

 // END SECTION OUTPUT
  

func main() {
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
	
	/*test values
	ex_val_id := [5]string{"A", "B", "C", "G", "F"}
	ex_val_predecessor := [5]string{"R", "C", "D", "C", "C"}
	ex_val_distance := [5]int{1, 3, 2, 3, 3}
	*/
    /*
	fmt.Println("_____________________________")
	fmt.Println(" id - predecessor - distance ")
	fmt.Println("_____________________________")
	for i := 0; i < 5; i++ {
		ll := &line{id: ex_val_id[i], predecessor: ex_val_predecessor[i], distance: ex_val_distance[i]}
		ll.describe()
	}
    */
/* BIGGER GRAPH : to see later
	min := 1
	max := 20
	for i := 0; i < 20; i++ {
		dist = rand.Intn(max - min) + min)
		graph.addEdge("S", "B", dist)
	}	
*/

  }