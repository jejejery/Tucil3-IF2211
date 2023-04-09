package algorithm

import (
	structs "backend/lib/structs"
	util "backend/lib/util"
)

func UCS(datainfo map[int]structs.MapValue, graph structs.Graph, idxstart int, idxdest int) structs.NodeInfo{
	var nodes structs.PrioQueue
	structs.CreatePrioQueue(&nodes)
	visited := make([]int, 0, 5)

	var currentNodeInfo structs.NodeInfo 
	structs.CreateEmpty(&currentNodeInfo, idxstart)

	for structs.GetId(currentNodeInfo) != idxdest {
		if !util.IsMember(visited, structs.GetId(currentNodeInfo)) {
			for i := 0; i < structs.GetNumNodes(graph); i++ {
				if (!util.IsMember(visited, i) && structs.GetValue(graph, structs.GetId(currentNodeInfo), i) != 0) {
					var temp structs.NodeInfo
					tempPath := make([]int, len(structs.GetPath(currentNodeInfo)))
					copy(tempPath, structs.GetPath(currentNodeInfo))
					structs.CreateInfo(&temp, i, append(tempPath, structs.GetId(currentNodeInfo)), structs.GetPathCost(currentNodeInfo) + structs.GetValue(graph, structs.GetId(currentNodeInfo), i))
					var tempNode structs.Node
					structs.CreateNode(&tempNode, temp)
					
					structs.Enqueue(&nodes, &tempNode)
				}
			}
			visited = append(visited, structs.GetId(currentNodeInfo))
		}
		if (structs.LenQueue(nodes) == 0) {
			return structs.CreateInvalidInfo()
		}
		currentNodeInfo = structs.Dequeue(&nodes)
	}
	structs.CreateInfo(&currentNodeInfo, structs.GetId(currentNodeInfo), append(structs.GetPath(currentNodeInfo), structs.GetId(currentNodeInfo)), structs.GetPathCost(currentNodeInfo))
	return currentNodeInfo
}