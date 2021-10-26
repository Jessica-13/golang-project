package main

import (
	"fmt"
)

type line struct {
	id string
	predecessor string
	distance int
  }

func (l *line) describe() {
	fmt.Printf("  %v -      %v      -    %v \n", l.id, l.predecessor, l.distance)
	/*for id, predecessor,distance := range ex_val_distance {
        fmt.Printf("%c - %c = %d\n", id, predecessor, distance)
    }*/
	/*
	for i := range ex_val_distance {
		fmt.Printf("%c - ", id)
		fmt.Printf("%c = ", predecessor)
        fmt.Printf("%d\n", distance)
    }*/
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

func main() {
	ll := &line{id: "A", predecessor: "R", distance: 1}
	/*
	//declaration empty slice
	var ex_val_distance line										//ex_val_distance = make([]int,3,8)   ---   ex_val_distance[0]=10
	// join element 10
	ex_val_distance=append(ex_val_distance,10)
	// join element 10
	ex_val_distance=append(ex_val_distance,20)
	// print content and dimensions
	fmt.Println(ex_val_distance)
	//fmt.Println(len(ex_val_distance))	//length
	//fmt.Println(cap(ex_val_distance))	//capacity
	
	var ex_val_distance[10] string
	for x := 0; x < 3; x++ {
		ex_val_distance[x] = "A"
	}
	for i := range ex_val_distance {
		ll.setId("B")
		ll.setPredecessor("C")
		ll.setDistance(3)
		ll.describe()
    }
	*/
	fmt.Println(" id - predecessor - distance ")
	ll.describe()
  }