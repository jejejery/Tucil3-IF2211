package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	structs "backend/lib/structs"
)

func ReadInput(datainfo *map[int]structs.MapValue, graph *structs.Graph) {
	f, err := os.Open("..\\..\\test\\test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	numNodesStr := scanner.Text()
	fmt.Println(numNodesStr)
	numNodes, errNumNodes := strconv.Atoi(numNodesStr)

	if (errNumNodes != nil) {
		log.Fatal(errNumNodes)
	}

	for i := 0; i < numNodes; i++ {
		scanner.Scan()
		name := scanner.Text()
		scanner.Scan()
		xstr := scanner.Text()
		x, errx := strconv.ParseFloat(xstr, 64)
		if (errx != nil) {
			log.Fatal(errx)
		}
		scanner.Scan()
		ystr := scanner.Text()
		y, erry := strconv.ParseFloat(ystr, 64)
		if (erry != nil) {
			log.Fatal(erry)
		}
		var p structs.Point
		structs.CreatePoint(&p, x, y)
		var val structs.MapValue
		structs.CreateMapValue(&val, name, p)
		(*datainfo)[i] = val
	}

	structs.CreateGraph(graph, numNodes)
	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			scanner.Scan()
			xstr := scanner.Text()
			x, errx := strconv.ParseFloat(xstr, 64)
			if (errx != nil) {
				log.Fatal(errx)
			}
			structs.AddEdge(graph, i, j, x)
		}
	}

	fmt.Println(datainfo)
	fmt.Println(structs.GetAdjMatrix(*graph))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}