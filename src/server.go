/*
A TCP server written in Go.
*/
package main

import hp "container/heap" 

import (
	"math/rand"
)

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// ---

type path struct {
    value int
    nodes []string
}

type minPath []path

func (h minPath) Len() int { 
	return len(h) 
}

func (h minPath) Less(i, j int) bool { 
    return h[i].value < h[j].value 
}

func (h minPath) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i] 
}

func (h *minPath) Push(x interface{}) {
    *h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type heap struct {
    values *minPath
}

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

type edge struct {
    node   string
    weight int
}

type graph struct {
    nodes map[string][]edge
}

func newGraph() *graph {
    return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
    g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
    g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
    fmt.Println(" ", origin, "      ", destiny, "         ", weight)
}

func (g *graph) getEdges(node string) []edge {

    return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
    h := newHeap()
    h.push(path{value: 0, nodes: []string{origin}}) 
    visited := make(map[string]bool)

    for len(*h.values) > 0 {
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
                h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
            }
        }
        visited[node] = true
    }
    return 0, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

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
    max := 20
    vertex1 := RandStringBytes(1)
    for i := 0; i < 10; i++ {   // all vertices linked (x10)
        vertex2 := RandStringBytes(1)
        distance := rand.Intn(max - min) + min
        //fmt.Println(vertex1, vertex2, distance)
        graph.addEdge(vertex1, vertex2, distance)
        vertex1 = vertex2
    }	

    for i := 0; i < 20; i++ {   // more random connections (x20)
        vertex1 := RandStringBytes(1)
        vertex2 := RandStringBytes(1)
        distance := rand.Intn(max - min) + min
        graph.addEdge(vertex1, vertex2, distance)
    }

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
				
		stringOrigin := string(take[0])
		stringDestination := string(take[1])
		fmt.Println(graph.getPath(stringOrigin, stringDestination))
		fmt.Println(" ")
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


    