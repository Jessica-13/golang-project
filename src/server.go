// A TCP server written in Go.

package main

import hp "container/heap" 

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
    "math/rand"
    "strings"
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

func (h *minPath) Push(x interface{}) {
    *h = append(*h, x.(path))
}

/* Find the shortest path
old : is the list of the analyzed paths
n   : is an index representing the length of old [1,2]
x   : represents the element of old at the index (n-1) ie [0,1]
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

func (h *heap) push(p path) {
    hp.Push(h.values, p)
}

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

// Initialization of a new graph
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
    g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
    fmt.Println(" ", origin, "      ", destiny, "         ", weight)
}

// Definition anonymous fonction that set relations between edges : 
func (g *graph) getEdges(node string) []edge {

    return g.nodes[node]
}


func (g *graph) getPath(origin, destiny string) (int, []string) {   //gets the shortest path between origin and destiny
    h := newHeap()                                                  // definition of a new heap named h -> call fonction newHeap()
    h.push(path{value: 0, nodes: []string{origin}}) 
    visited := make(map[string]bool)                                // Mark all nodes unvisited. 

    // nodes control :
    for len(*h.values) > 0 {
        // Find the nearest yet to visit node
        p := h.pop()
        node := p.nodes[len(p.nodes)-1]

        if visited[node] {          // if already visited
            continue                    // do nothing
        }

        if node == destiny {        // when we reached the destination
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
            }
        }
        visited[node] = true        // once a node is checked, it's marked as visited
    }
    return 0, nil
}
// END graph definition



var originVertexInput string 
var destinationVertexInput string

var okOrigin bool
var okDestination bool

// ---

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8000, "The port to listen on; default is 8000.")

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

    min := 1
    max := 30

    nbVertexGraph := 5

    // MAKE THE GRAPH 

    
    // WITHOUT GO ROUTINES
    timeI := time.Now()     // TIME START
    for vertex1 := 0; vertex1 < nbVertexGraph; vertex1++ {      // nombre vertex
        for vertex2 := vertex1; vertex2 < nbVertexGraph; vertex2++ {      // graph complet
            vertex1S := strconv.FormatInt(int64(vertex1), 10)
            vertex2S := strconv.FormatInt(int64(vertex2), 10)
            distance := 0
            if vertex1 != vertex2 {
                distance = rand.Intn(max - min) + min
            }
            graph.addEdge(vertex1S, vertex2S, distance)
        }
    }

    timeF := time.Now()     // TIME STOP
    fmt.Println("Difference time make algo : ", timeF.Sub(timeI))       // OUTPUT TIME
    
	// ---

	flag.Parse()

	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go handleConnection(conn, graph)
	}
}

func handleConnection(conn net.Conn, graph *graph) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)


	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		take := handleMessage(scanner.Text(), conn)      // Taking input from client 
		fmt.Print("Shortest path calculation result (distance - path) : ")
         
        takeSplit := strings.Split(take, " ")   // split the input into origin and destination
        
        stringOrigin := takeSplit[0]
		stringDestination := takeSplit[1]

        // OUTPUT INTO SERVER
		fmt.Println(graph.getPath(stringOrigin, stringDestination))
		fmt.Println(" ")

        // for reply - OUTPUT INTO CLIENT
	// The shortest distance and the nodes of the shortest path returned by the getPath function are converted into string form and returned to the client
        AAA,BBB := graph.getPath(stringOrigin, stringDestination)
        t := strconv.Itoa(AAA)
        justString := strings.Join(BBB, " ")
        newmessage := string("Shortest path calculation result. Distance : " + t + " - Path : " + justString)
        conn.Write([]byte(newmessage + "\n"))

	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) string {
	fmt.Println("> " + message)

	if len(message) > 0 && message[0] == '/' {
		switch {
		case message == "/time":
			resp := "It is " + time.Now().String() + "\n"
			fmt.Print("< " + resp)
			conn.Write([]byte(resp))

		case message == "/quit":
			fmt.Println("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			fmt.Println("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			os.Exit(0)

		default:
			conn.Write([]byte("Unrecognized command.\n"))
		}
	}

	return message
}
