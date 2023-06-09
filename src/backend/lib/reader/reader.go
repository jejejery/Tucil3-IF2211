package reader

import (
	structs "backend/lib/structs"
	"bufio"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReadInputLocal(datainfo *map[int]structs.MapValue, graph *structs.Graph, dir string) {
	f, err := os.Open(dir)
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

	if errNumNodes != nil {
		log.Fatal(errNumNodes)
	}

	for i := 0; i < numNodes; i++ {
		scanner.Scan()
		name := scanner.Text()
		scanner.Scan()
		xstr := scanner.Text()
		x, errx := strconv.ParseFloat(xstr, 64)
		if errx != nil {
			log.Fatal(errx)
		}
		scanner.Scan()
		ystr := scanner.Text()
		y, erry := strconv.ParseFloat(ystr, 64)
		if erry != nil {
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
			if errx != nil {
				log.Fatal(errx)
			}
			structs.AddEdge(graph, i, j, x)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadInput(datainfo *map[int]structs.MapValue, graph *structs.Graph, f multipart.File, c *gin.Context, her *float64) {

	//preprocess the input.txt
	buf := make([]byte, 1024)
	n, err := f.Read(buf)

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
		return
	}

	buffer := []string{}
	lines := strings.Split(string(buf[:n]), "\n")

	for _, line := range lines {
		lineWords := strings.Split(line, " ")
		buffer = append(buffer, lineWords...)
	}

	// How many vertices in graph?
	iterator := 0
	numNodes, errNumNodes := strconv.Atoi(string(buffer[iterator]))
	iterator++
	if errNumNodes != nil {
		fmt.Println("error ngab")
		c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
		return
	}
	for {
		if string(buffer[iterator]) != "" {
			break
		} else {
			iterator++
		}

	}


	hc, errHc := strconv.ParseFloat(buffer[iterator], 64)
	
	iterator++
	if errHc != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
		return
	}
	*her = hc

	//Add heruistic information
	for i := 0; i < numNodes; i++ {
		name := buffer[iterator]
		iterator++

		xstr := buffer[iterator]
		iterator++
		x, errx := strconv.ParseFloat(xstr, 64)
		if errx != nil {
			log.Fatal(errx)
			c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
			return
		}

		ystr := buffer[iterator]
		iterator++
		y, erry := strconv.ParseFloat(ystr, 64)
		if erry != nil {
			log.Fatal(erry)
			c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
			return
		}
		
		var p structs.Point
		structs.CreatePoint(&p, x, y)
		var val structs.MapValue
		structs.CreateMapValue(&val, name, p)
		(*datainfo)[i] = val

		for {
			if string(buffer[iterator]) != "" {
				break
			} else {
				iterator++
			}
		}
	}

	structs.CreateGraph(graph, numNodes)

	for {
		if string(buffer[iterator]) != "" {
			break
		} else {
			iterator++
		}
	}

	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			xstr := buffer[iterator]
			x, errx := strconv.ParseFloat(xstr, 64)
			if errx != nil {
				log.Fatal(errx)
				c.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
				return
			}
			structs.AddEdge(graph, i, j, x)
			iterator++
		}
		for {
			if i != numNodes-1{
				if string(buffer[iterator]) != "" {
					break
				} else {
					iterator++
				}
			}
		}
	}


}
