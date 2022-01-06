package main

import (
    "fmt"
)

const nbVertex int = 
const maxWeight int = 10000 

var getShortPath = [nbVertex]int{    }

func main() {
 
    var TablePathMin int       //Stores the smallest value of the node in the getShortPath that is unvisited
    var Vx int                 //Stores the subscript of the node in the getShortPath that is unvisited
    var isgetPath [MAXVEX]bool //Records whether the node has found the minimal path from v0 to vx or not
 
    // Get the array of weights for the row v0
    for v := 0; v < len(   ); v++ {
        getShortPath[v] = 
    }
    getShortPath[0] = 0
    isgetPath[0] = true
 
    //from v1 ~ vn
    for v := 1; v < len(  ); v++ {
        TablePathMin = maxWeight
 
        //finds the smallest value of the node in the getShortPath that is unvisited
        for w := 0; w < len(   ); w++ {
            if !isgetPath[w] && getShortPath[w] < TablePathMin {
                Vx = w
                TablePathMin = getShortPath[w]
            }
        }
        isgetPath[Vx] = true
        for j := 0; j < len(  ); j++ {
            if !isgetPath[j] && TablePathMin+graph[Vx][j] < getShortPath[j] {
                getShortPath[j] = TablePathMin + 
            }
        }
    }
