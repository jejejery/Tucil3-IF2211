package structs

type Graph struct {
    nodes int
    adj_matrix [][]int
}

func CreateGraph(nodes int) *Graph {
    graph := &Graph{}
    graph.nodes = nodes
    graph.adj_matrix = make([][]int, nodes)
    for i := 0; i < nodes; i++ {
        graph.adj_matrix[i] = make([]int, nodes)
    }
    return graph
}

func (graph *Graph) AddEdge(source int, destination int, weight int) {
    graph.adj_matrix[source][destination] = weight
	graph.adj_matrix[destination][source] = weight
}

func (graph *Graph) RemoveEdge(source int, destination int) {
    graph.adj_matrix[source][destination] = 0
	graph.adj_matrix[destination][source] = 0
}

func (graph *Graph) HasEdge(source int, destination int) bool {
    return graph.adj_matrix[source][destination] != 0
}
