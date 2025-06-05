package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	myMovie := Movie{"喜剧之王", 2000, 10, []string{"周星驰,张柏芝"}}
	marshal, err := json.Marshal(myMovie)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%s \n", marshal)

	move := Movie{}
	err = json.Unmarshal(marshal, &move)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%v \n", move)
}
