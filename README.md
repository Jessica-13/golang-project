# Dijkstra
A Golang implementation of Dijkstra's algorithm which is used to find the shortest paths between nodes in a graph.


## Design
We implement this algorithm through a client-server format. This allows several clients  to do the calculations in parallel, through the use of the functionality of the go routines. <br>

### Files :

- client.go <br> 
- server.go <br>

<hr />

## Description of the Dijkstra's algorithm
The algorith picks the unvisited node with the lowest distance, calculates the distance through it to each unvisited neighbor,and updates the neighbor's distance if smaller.

### STEPS:

 <strong>1)</strong> Mark all nodes unvisited. 
 
 <strong>2)</strong> Create a set of all the unvisited nodes called the unvisited set, in our case we are going to use a set for visited nodes, not for unvisited nodes.

 <strong>3)</strong> Assign to every node a tentative distance value: set it to zero for our initial node. Then set the initial node as current.

 <strong>4)</strong> For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node. 
 
 <strong>5)</strong> Compare the newly calculated tentative distance to the current assigned value and assign the smaller one. Otherwise, keep the current value.

 <strong>6)</strong> When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited. A visited node will never be checked again.

 <strong>7)</strong> Select next unvisited node that is marked with the smallest tentative distance, set it as the new "current node", and go back to step 3.

=> Dijkstra's algorithm, can be optimized using a priority queue => using <strong>heap</strong>.

### What's a Heap?

An almost complete tree that satisfies the heap property: 
- In a max heap, for any given node C, if P is a parent node of C, then the key (the value) of P is greater than or equal to the key of C. 
- In a min heap, the key of P is less than or equal to the key of C The node at the "top" of the heap (with no parents) is called the root node.

A heap can be thought of as a priority queue; the most important node will always be at the top, and when removed, its replacement will be the most important. 
This can be useful when coding algorithms that require certain things to processed in a complete order,but when you don't want to perform a full sort or need to know anything about the rest of the nodes. 

<hr />

## Implementation
First, call server.go by `go run server.go -port 8000` in a terminal. <em>Client connect to the server from the default port 8000.</em>
The program establishes a random graph which is displayed in the server terminal. 
Each vertex is assigned a random distance assignment between 1 and 20, and random connections are assigned between the vertices. <br>

<p align="center">
<a href="https://asciinema.org/a/k0YpaYoXPArwKaKAxiXdmcE2A?autoplay=1?loop=1" target="_blank"><img src="https://asciinema.org/a/k0YpaYoXPArwKaKAxiXdmcE2A.svg" /></a>
</p>

Then, call client.go by `go run client.go -host localhost -port 8000` in a new terminal.
Once a client has connected, the server terminal will display it. By entering two vertices in the client's terminal, the program allows to obtain the shortest distance between any two vertices entered by the user and shows which vertices the path of this shortest distance passes through. 
If the user enters a non-existent vertex, the distance will be displayed as 0.

<p align="center">
<a href="https://asciinema.org/a/RL8M8QD11rzIzdNCmKwrWFJbT?autoplay=1?loop=1" target="_blank"><img src="https://asciinema.org/a/RL8M8QD11rzIzdNCmKwrWFJbT.svg" /></a>
</p>

The server can connect to multiple clients. 

<p align="center">
<a href="https://asciinema.org/a/odi27lUT4CenVoppmNjeJ8Ky7?autoplay=1?loop=1" target="_blank"><img src="https://asciinema.org/a/odi27lUT4CenVoppmNjeJ8Ky7.svg" /></a>
</p>

The solution is accessible by both server.go and client.go.


## Requirements

Tested on
```
go v1.17
```
<hr />

## Credits

<p align="center">
  <img src="http://www.insa-lyon.fr/sites/www.insa-lyon.fr/files/logo-coul.jpg" width="350" alt="accessibility text">
</p>

<strong>INSA Lyon</strong>, Lyon Institute of Applied Sciences <br/> 
Department of Telecommunications, Services and Uses, 3TC, Group 1

Project related to the ELP module (Ecosystème des langages de programmation) - Golang.

### Referent Professor

Pierre François

### Authors

SALMA Bahar <br/>
SPERA Jessica <br/>
WAN Zihao <br/>

