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
	fmt.Printf(" %v - %v - %v \n", l.id, l.predecessor, l.distance)
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
	ll.describe()
	//ll.setDistance(3)
	//fmt.Println(ll.id)
	//ll.setPredecessor("D")
	//fmt.Println(ll.predecessor)
  }