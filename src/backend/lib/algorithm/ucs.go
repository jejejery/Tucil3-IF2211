package algorithm

import (
	structs "backend/lib/structs"
	util "backend/lib/util"
)

func UCS(datainfo map[int]structs.MapValue, graph structs.Graph, idxstart int, idxdest int) {
	nodes := make([]structs.NodeInfo, 0, 5)
	visited := make([]int, 0, 5)

	var currentNode int = idxstart
	var currentNodeInfo structs.NodeInfo

	for currentNode != idxdest {
		for i := 0; i < structs.GetNumNodes(graph); i++ {
			if (!util.IsMember(visited, i) && structs.GetValue(graph, currentNode, i) != 0) {
				
			}
		}
	}
}