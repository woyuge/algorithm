package main
//单链表删除节点
import (
	"fmt"
)

func main()  {
// head = [4,5,1,9]
	 head := new(listNode)
	 head.val = 4
	 
	 node1 := new(listNode)
	 node1.val = 5
	 head.Next = node1
	 
	 node2 := new(listNode)
	 node2.val = 1
	 node1.Next = node2
	 
	 node3 := new(listNode)
	 node3.val = 9
	 node2.Next = node3
	// showNodes(head)
	 fmt.Println("-----------------")
	// showNodes(head)
}
//反转链表
func reverseList(head *ListNode) *ListNode {

	

}