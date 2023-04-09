package structs

type PrioQueue struct {
	node *Node
}

func CreatePrioQueue(p PrioQueue) {
	p.node = nil
}

func Enqueue(queue *PrioQueue, node *Node) {
	if queue.node == nil {
		queue.node = node
	} else {
		var prev *Node = nil
		p := queue.node
		for GetPathCost(*p) <= GetPathCost(*node) && p != nil {
			prev = p
			p = p.next
		} // GetPathCost(*p) > GetPathCost(*node)
		if prev == nil {
			node.next = p
			queue.node = node
		} else if p == nil {
			prev.next = node
		} else {
			node.next = p
			prev.next = node
		}
	}
}

func Dequeue(queue *PrioQueue) NodeInfo {
	temp := queue.node.info
	queue.node = queue.node.next
	return temp
}