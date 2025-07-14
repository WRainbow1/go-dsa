
package main

type ListNode struct {
     Val int
     Next *ListNode
 }

func reorderList(head *ListNode)  {

    if head == nil || head.Next == nil {
        return
    }

    curr := head

    var stack []*ListNode
	visited := make(map[*ListNode]bool)

	for i := 1; true; i++{
		stack = append(stack, curr)
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}
    
	curr = head
    for i := 1; true; i++ {
		
		last := stack[len(stack)-1]   // Get the last element
		stack = stack[:len(stack)-1]

		if visited[last] {
			curr.Next = nil
			break
		}

		visited[curr] = true
		visited[last] = true

        last.Next = curr.Next
		curr.Next = last
		curr = last.Next
    }
}