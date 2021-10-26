package main

import (
	"fmt"
)

/*
		Graph initialization
*/

/*
		Dijkstra integration <- calculDisMin.go
*/

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

func main() {
	//test values
	ex_val_id := [5]string{"A", "B", "C", "G", "F"}
	ex_val_predecessor := [5]string{"R", "C", "D", "C", "C"}
	ex_val_distance := [5]int{1, 3, 2, 3, 3}
	
	fmt.Println(" id - predecessor - distance ")
	fmt.Println("_____________________________")
	for i := 0; i < 5; i++ {
		ll := &line{id: ex_val_id[i], predecessor: ex_val_predecessor[i], distance: ex_val_distance[i]}
		ll.describe()
	}	
  }