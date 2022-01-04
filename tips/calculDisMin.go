package main

import (
    "fmt"
)

const nbVertex int = 
const maxWeight int = 10000 

var getShortPath = [nbVertex]int{    }

func main() {
 
    var TablePathMin int       //Store the smallest value of the node in the getShortPath that is not traversed
    var Vx int                 //Store the subscript of the node in the getShortPath that is not traversed
    var isgetPath [MAXVEX]bool //Record whether the node has found the minimal path from v0 to vx
 
    // Get the array of weights for the row v0
    for v := 0; v < len(   ); v++ {
        getShortPath[v] = 
    }
    getShortPath[0] = 0
    isgetPath[0] = true
 
    //from v1 ~ vn
    for v := 1; v < len(  ); v++ {
        TablePathMin = maxWeight
 
        //find the smallest value of the node in the getShortPath that is not traversed
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
