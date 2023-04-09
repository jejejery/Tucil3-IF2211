package structs
import "fmt"

type Graph struct {
    numNodes int
    adjMatrix [][] float64;
}

func CreateGraph(graph *Graph, numNodes int) {
    graph.numNodes = numNodes
    graph.adjMatrix = make([][]float64, numNodes)
    for i := 0; i < numNodes; i++ {
        graph.adjMatrix[i] = make([]float64, numNodes)
    }
}

func AddEdge(graph *Graph, source int, destination int, weight float64) {
    graph.adjMatrix[source][destination] = weight
	graph.adjMatrix[destination][source] = weight
}

func RemoveEdge(graph *Graph, source int, destination int) {
    graph.adjMatrix[source][destination] = 0
	graph.adjMatrix[destination][source] = 0
}


func PrintGraph(graph Graph) {
    for i := 0; i < graph.numNodes; i++ {
        for j := 0; j < graph.numNodes; j++ {
            fmt.Printf("%.3f ",graph.adjMatrix[i][j])
        }
        fmt.Println()
    }
}

