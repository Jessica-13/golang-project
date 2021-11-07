/* 
Secondly, the logic for the graph: 
we use a struct that contains a map to keep the edges among the nodes, 
with functions to add the edges and get all edges from one node.

nous utilisons une structure qui contient une map pour conserver les edges entre les nœuds
avec des fonctions pour ajouter les edges et obtenir toutes les edges d'un nœud.
*/

package main

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
}

func (g *graph) getEdges(node string) []edge {
    return g.nodes[node]
}


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



