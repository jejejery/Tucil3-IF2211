package structs

import "fmt"
type Graph struct {
    nodes int
    adj_matrix [][]float64
}

func CreateGraph(nodes int) *Graph {
    graph := &Graph{}
    graph.nodes = nodes
    graph.adj_matrix = make([][]float64, nodes)
    for i := 0; i < nodes; i++ {
        graph.adj_matrix[i] = make([]float64, nodes)
    }
    return graph
}

func (graph *Graph) AddEdge(source int, destination int, weight float64) {
    graph.adj_matrix[source][destination] = weight
	graph.adj_matrix[destination][source] = weight
}

func (graph *Graph) RemoveEdge(source int, destination int) {
    graph.adj_matrix[source][destination] = 0
	graph.adj_matrix[destination][source] = 0
}


func (graph *Graph) PrintGraph() {
    for i := 0; i < graph.nodes; i++ {
        for j := 0; j < graph.nodes; j++ {
            fmt.Printf("%.3f ",graph.adj_matrix[i][j])
        }
        fmt.Println()
    }
}

