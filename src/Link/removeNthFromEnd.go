package main

//删除链表的倒数第n个节点

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	head := new(ListNode)
	head.Val = 4

	node1 := new(ListNode)
	node1.Val = 5
	head.Next = node1

	node2 := new(ListNode)
	node2.Val = 1
	node1.Next = node2

	node3 := new(ListNode)
	node3.Val = 9
	node2.Next = node3

	removeNthFromEnd1(head,2)

}
//**得先找到要删除节点的前一个元素
func removeNthFromEnd1(head *ListNode, n int) *ListNode {

	l := getLength(head)//链表长度
	dump := &ListNode{0,head}
	cur:= dump
	for i:=0;i<l-n ;i++{
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dump.Next
}
//获取节点的长度
func getLength(head *ListNode)  int {
	length:=0
	for ;head!=nil;head= head.Next {
		length ++
	}
	return length
}