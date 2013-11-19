package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex
var n map[int]Vertex

func main() {
    m = make(map[string]Vertex)
    n = make(map[int]Vertex)
    
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    
    n[1] = Vertex{ 20,  20,}
    fmt.Println(m["Bell Labs"])
    fmt.Println(n[1])
    
}