package structs

type Node struct {
	info NodeInfo
	next *Node
}

type NodeInfo struct {
	id int
	path []int
	pathCost float64
}

func CreateNode(node *Node, info NodeInfo) {
	node.info = info
	node.next = nil
}

func CreateInfo(node *NodeInfo, id int, path []int, pathCost float64) {
	node.id = id
	node.path = path
	node.pathCost = pathCost
}

func CreateInvalidInfo() NodeInfo {
	var temp NodeInfo
	temp.id = -1
	temp.pathCost = 0
	return temp
}

func CreateEmpty(node *NodeInfo, id int) {
	node.id = id
	node.path = make([]int, 0, 5)
	node.pathCost = 0
}

func GetIdNode(node Node) int {
	return node.info.id
}

func GetPathNode(node Node) []int {
	return node.info.path
}

func GetPathCostNode(node Node) float64 {
	return node.info.pathCost
}

func GetId(node NodeInfo) int {
	return node.id
}

func GetPath(node NodeInfo) []int {
	return node.path
}

func GetPathCost(node NodeInfo) float64 {
	return node.pathCost
}


