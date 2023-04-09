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

func GetAdjMatrix(graph Graph) [][] float64 {
    return graph.adjMatrix
}

func GetNumNodes(graph Graph) int {
    return graph.numNodes
}

func AddEdge(graph *Graph, source int, destination int, weight float64) {
    graph.adjMatrix[source][destination] = weight
}

func RemoveEdge(graph *Graph, source int, destination int) {
    graph.adjMatrix[source][destination] = 0
}

func GetValue(graph Graph, source int, destination int) float64 {
    return graph.adjMatrix[source][destination]
}

func PrintGraph(graph Graph) {
    for i := 0; i < graph.numNodes; i++ {
        for j := 0; j < graph.numNodes; j++ {
            fmt.Printf("%.3f ",graph.adjMatrix[i][j])
        }
        fmt.Println()
    }
}

