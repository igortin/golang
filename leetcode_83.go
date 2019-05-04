/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var cNode = head
    var p = head
	for cNode.Next != nil {
        if cNode.Val ==  cNode.Next.Val {
			cNode.Next = cNode.Next.Next
		} else {
		    cNode = cNode.Next
		}
    }
    return p
}
