package structs

import "fmt"

type PrioQueue struct {
	node *Node
}

func CreatePrioQueue(p *PrioQueue) {
	p.node = nil
}

func Enqueue(queue *PrioQueue, node *Node) {
	if queue.node == nil {
		queue.node = node
	} else {
		//fmt.Println("-----------------------")
		var prev *Node = nil
		p := queue.node
		for p != nil && GetPathCostNode(*p) <= GetPathCostNode(*node) {
			//fmt.Println(prev, p, node)
			if (node.info.pathCost == 7) {
				//fmt.Println(p)
			}
			prev = p
			p = p.next
			if (node.info.pathCost == 7) {
				//fmt.Println(prev)
			}
		} // GetPathCost(*p) > GetPathCost(*node)
		if prev == nil {
			node.next = p
			queue.node = node
		} else if p == nil {
			//fmt.Println("hehe", prev, node)
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

func LenQueue(queue PrioQueue) int {
	var i int = 0
	p := queue.node
	for p != nil {
		i++;
		p = p.next
	}
	return i
}

func PrintQueue(queue PrioQueue) {
	p := queue.node
	for p != nil {
		fmt.Print(*p, " ")
		p = p.next
	}
	fmt.Println()
}