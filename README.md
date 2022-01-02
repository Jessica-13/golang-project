# Dijkstra
A golang implementation of Dijkstra's algorithm which is used to find the shortest paths between nodes in a graph, based on the description at [Wikipedia](http://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm).
## Design
We implement this algorithm through a client-server format.<br>
- client.go <br> 
- server.go <br>

The solution is accessible via server.go.
## Implementation
First, call server.go by `go run server.go`. The program first establishes a random chain of ten vertices, each with a random distance assignment between 1 and 20, and then establishes a random connection between twenty pairs of vertices, again with a random distance assignment. This procedure generates a graph. The program allows to obtain the shortest distance between any two vertices entered by the user and shows which vertices the path of this shortest distance passes through. If the user enters a non-existent vertex, the distance will be displayed as 0.<br>


