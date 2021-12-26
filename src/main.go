package main

import hp "container/heap" 

import (
	"fmt"
	"math/rand"
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
    /*
    fmt.Println("try h : ", h) 
    try h :  [{0 [S]}]
    try h :  [{0 [S]}]
    try h :  [{4 [S B]}]
    try h :  [{4 [S B]} {2 [S C]}]
    try h :  [{2 [S C]} {4 [S B]}]
    try h :  [{4 [S B]} {3 [S C B]}]
    try h :  [{3 [S C B]} {4 [S B]}]
    try h :  [{4 [S B]} {8 [S C B D]}]
    try h :  [{4 [S B]} {8 [S C B D]}]
    try h :  [{8 [S C B D]}]
    */
    
    /*
    fmt.Println("try len(h) : ", h) 
    try len(h) :  [{0 [S]}]
    try len(h) :  [{0 [S]}]
    try len(h) :  [{4 [S B]}]
    try len(h) :  [{4 [S B]} {2 [S C]}]
    try len(h) :  [{2 [S C]} {4 [S B]}]
    try len(h) :  [{4 [S B]} {3 [S C B]}]
    try len(h) :  [{3 [S C B]} {4 [S B]}]
    try len(h) :  [{4 [S B]} {8 [S C B D]}]
    try len(h) :  [{4 [S B]} {8 [S C B D]}]
    try len(h) :  [{8 [S C B D]}]
    */
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
    /*
    fmt.Println("try h[i].value : ", h[i].value)
    fmt.Println("try h[j].value : ", h[j].value)
    fmt.Println("try h[i].value < h[j].value : ", h[i].value < h[j].value)

    try h[i].value :  2
    try h[j].value :  4
    try h[i].value < h[j].value :  true
    try h[i].value :  3
    try h[j].value :  4
    try h[i].value < h[j].value :  true
    try h[i].value :  8
    try h[j].value :  4
    try h[i].value < h[j].value :  false
    */
    return h[i].value < h[j].value 
}

/* Definition anonymous fonction that : 
input  -> h (minPath type)
-> call the Swap() fonction
output -> If the second h[j] is the shorter one, 
          inversion of the two to be able to compare it with the next value
*/
func (h minPath) Swap(i, j int) { 
    /*
    fmt.Println("try : ", h[i])
    fmt.Println("try : ", h[j])

    try :  {0 [S]}
    try :  {0 [S]}
    try :  {4 [S B]}
    try :  {2 [S C]}
    try :  {2 [S C]}
    try :  {4 [S B]}
    try :  {4 [S B]}
    try :  {3 [S C B]}
    try :  {3 [S C B]}
    try :  {4 [S B]}
    try :  {4 [S B]}
    try :  {8 [S C B D]}
    try :  {8 [S C B D]}
    try :  {8 [S C B D]}
    */
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
    /*
    fmt.Println("try : ", *h)

    try :  [{0 [S]}]
    try :  [{4 [S B]}]
    try :  [{4 [S B]} {2 [S C]}]
    try :  [{4 [S B]} {3 [S C B]}]
    try :  [{4 [S B]} {8 [S C B D]}]
    */
}

/* Find the shortest path
old : is the list of the analyzed paths
n   : is an index representing the length of old [1,2]
x   : represents the element of old at the index (n-1) ie [0,1]
*/
func (h *minPath) Pop() interface{} {
    old := *h
    /*
    fmt.Println("try : ", *h)
    try :  [{0 [S]}]
    try :  [{4 [S B]} {2 [S C]}]
    try :  [{4 [S B]} {3 [S C B]}]
    try :  [{8 [S C B D]} {4 [S B]}]
    try :  [{8 [S C B D]}]
    */
    n := len(old)
    /*
    fmt.Println("try : ", n)
    try :  1
    try :  2
    try :  2
    try :  2
    try :  1
    */
    x := old[n-1]
    /*
    fmt.Println("try : ", x)
    try :  {0 [S]}
    try :  {2 [S C]}
    try :  {3 [S C B]}
    try :  {4 [S B]}
    try :  {8 [S C B D]}
    */
    *h = old[0 : n-1]
    /*
    fmt.Println("try : ", *h)
  
    try :  []
    try :  [{4 [S B]}]
    try :  [{4 [S B]}]
    try :  [{8 [S C B D]}]
    try :  []
    */
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
    /*
    fmt.Println("try : ", heap{values: &minPath{}})

    try :  {0xc00000c090}
    */
    return &heap{values: &minPath{}}
}

/*
*/
func (h *heap) push(p path) {
    hp.Push(h.values, p)
    /*
    fmt.Println("try h.values : ", *h.values)   // the possibilities

    try h.values :  [{0 [S]}]
    try h.values :  [{4 [S B]}]
    try h.values :  [{2 [S C]} {4 [S B]}]
    try h.values :  [{3 [S C B]} {4 [S B]}]
    try h.values :  [{4 [S B]} {8 [S C B D]}]

    fmt.Println("try p : ", p)  // the chosen path

    try p :  {0 [S]}
    try p :  {4 [S B]}
    try p :  {2 [S C]}
    try p :  {3 [S C B]}
    try p :  {8 [S C B D]}
    */
}

/*
*/
func (h *heap) pop() path {
    i := hp.Pop(h.values)
    /*
    fmt.Println("try i : ", i)

    try i :  {0 [S]}
    try i :  {2 [S C]}
    try i :  {3 [S C B]}
    try i :  {4 [S B]}
    try i :  {8 [S C B D]}
    
    fmt.Println("try i.(path) : ", i.(path))

    try i.(path) :  {0 [S]}
    try i.(path) :  {2 [S C]}
    try i.(path) :  {3 [S C B]}
    try i.(path) :  {4 [S B]}
    try i.(path) :  {8 [S C B D]}
    */
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

/* Definition anonymous fonction that set edges : 
input  -> g (*graph type)
output -> Adds in sequence to g.nodes[origin] an edge variable consisting of :
            a node (valeur destiny) 
            and a weight (valeur weight)
Almost the same for g.nodes[destiny] :
            a node (valeur origin) 
            and a weight (valeur weight)
*/
func (g *graph) addEdge(origin, destiny string, weight int) {
    g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
    /*
    fmt.Println("try g.nodes[origin] : ", g.nodes[origin])

    try g.nodes[origin] :  [{B 4}]
    try g.nodes[origin] :  [{B 4} {C 2}]
    try g.nodes[origin] :  [{S 4} {C 1}]
    try g.nodes[origin] :  [{S 4} {C 1} {D 5}]  
    */
    g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
    /*
    fmt.Println("try g.nodes[destiny] : ", g.nodes[destiny])

    try g.nodes[destiny] :  [{S 4}]
    try g.nodes[destiny] :  [{S 2}]
    try g.nodes[destiny] :  [{S 2} {B 1}]
    try g.nodes[destiny] :  [{B 5}]
    */

    // OUTPUT
	// fmt.Println(g.nodes[origin], "  ", g.nodes[destiny], "  ", weight) // not good
    fmt.Println(" ", origin, "      ", destiny, "         ", weight)
}

/* Definition anonymous fonction that set relations between edges : 
*/
func (g *graph) getEdges(node string) []edge {
    /*
    fmt.Println("try g.nodes[node] : ", g.nodes[node])

    try g.nodes[node] :  [{B 4} {C 2}]          // the neighboring vertices of S
    try g.nodes[node] :  [{S 2} {B 1}]          // the neighboring vertices of C
    try g.nodes[node] :  [{S 4} {C 1} {D 5}]    // the neighboring vertices of D
    */
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


// MAKE BIGGER GRAPH :
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"  // 52 characters

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}


func main() {
    fmt.Println(" ")
	fmt.Println("Dijkstra")
    fmt.Println(" ")
    fmt.Println(" ")
    fmt.Println("Chart description graph")
    fmt.Println(" ")
    fmt.Println("_____________________________")
    fmt.Println(" id - predecessor - distance ")
    fmt.Println("_____________________________")

   
    graph := newGraph() 

    // BIGGER GRAPH :
    min := 1
    max := 20
    vertex1 := RandStringBytes(1)
    for i := 0; i < 10; i++ {   // all vertices linked (x10)
        vertex2 := RandStringBytes(1)
        distance := rand.Intn(max - min) + min
        //fmt.Println(vertex1, vertex2, distance)
        graph.addEdge(vertex1, vertex2, distance)
        vertex1 = vertex2
    }	

    for i := 0; i < 10; i++ {   // more random connections (x10)
        vertex1 := RandStringBytes(1)
        vertex2 := RandStringBytes(1)
        distance := rand.Intn(max - min) + min
        graph.addEdge(vertex1, vertex2, distance)
    }


    /* Example
    graph.addEdge("S", "B", 4)
    graph.addEdge("S", "C", 2)
    graph.addEdge("B", "C", 1)
    graph.addEdge("B", "D", 5)
    
	
	graph.addEdge("C", "D", 8)
    graph.addEdge("C", "E", 10)
    graph.addEdge("D", "E", 2)
    graph.addEdge("D", "T", 6)
    graph.addEdge("E", "T", 2)
    fmt.Println(graph.getPath("S", "T"))
	*/

    // var then variable name then variable type
    var originVertexInput string 
    var destinationVertexInput string

    var okOrigin bool
    okOrigin = true
    var okDestination bool
    okDestination = true

    fmt.Println(" ")
    for okOrigin {
        fmt.Print("Please enter origin vertex  : ") 
        fmt.Scanln(&originVertexInput)      // Taking input from user 
        for i := range letterBytes {
            stringLetterBytes := string(letterBytes[i])
            if  originVertexInput == stringLetterBytes {
                okOrigin = false
            }
        }
    }	

    for okDestination {
        fmt.Print("Please enter destination vertex  : ")  
        fmt.Scanln(&destinationVertexInput) // Taking input from user
        for i := range letterBytes {
            stringLetterBytes := string(letterBytes[i])
            if  destinationVertexInput == stringLetterBytes {
                okDestination = false
            }
        }
    }	

    fmt.Print("Shortest path calculation result (distance - path) : ")
	fmt.Println(graph.getPath(originVertexInput, destinationVertexInput))
    fmt.Println(" ")

}