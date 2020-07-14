package main

import "fmt"

// create sruct Node
type Node struct {
	val int
	next *Node
}

// create struct the list integers
type intList struct {
	name	string
	head    *Node
	cNode 	*Node
}
// to create list of integers
func createintList(value string) *intList {
	// return instance intList with name
	return &intList{
		name: value,
	}
}

func (p *intList) addNode(value int) error {
	// create instance Node named s
	s := &Node{
		val:	value,
		next:	nil,
	}
	// if instance p of struct intList = empty
	if p.head == nil {
		p.head = s
	} else {
		// current position Node is head
		cNode := p.head
		// go till the array end where cNode do not have link to Next
		for cNode.next != nil {
			cNode = cNode.next
		}
		// at the integers array end func out instance Node named s
		cNode.next = s
	}
	return nil
}

// method of struct intList
func (p *intList) showElementLinkedList() error {
	cNode := p.head
	if cNode == nil {
		fmt.Println("LinkedList is empty.")
		return nil
	}
	// fmt.Printf("%+v %d\n", *cNode, cNode.val)
	fmt.Printf("%d ", cNode.val)
	for cNode.next != nil {
		cNode = cNode.next
	// fmt.Printf("%+v %d\n", *cNode, cNode.val)

		fmt.Printf("%d ", cNode.val)
	}
	return nil
}

func (p *intList) nextStep() *Node {
	p.cNode = p.head.next
	return p.cNode
}

func removeDuplicate(head *Node) *Node {
	if head == nil {
		return nil
	}
	var cNode = head
	var p = head
	for cNode.next != nil {
		if cNode.val == cNode.next.val {
			cNode.next = cNode.next.next
		} else {
			cNode = cNode.next
		}
	}
	return p
}

func main() {
	nums := []int{1, -1,-1, -2, -3,-3}
	linkedList := "linkedlist"
	ListNode := createintList(linkedList)

	for _,v := range nums {
		ListNode.addNode(v)
	}
	
	removeDuplicate(ListNode.head)
	
	ListNode.showElementLinkedList()
}
