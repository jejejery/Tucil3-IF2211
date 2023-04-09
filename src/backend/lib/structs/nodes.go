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

func GetIdNode(node Node) int {
	return node.info.id
}

func GetPathNode(node Node) []int {
	return node.info.path
}

func GetPathCost(node Node) float64 {
	return node.info.pathCost
}

