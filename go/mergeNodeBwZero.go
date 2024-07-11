package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
	  return nil
	}
	temp := &ListNode{Val: 0}
	current := temp
	sum := 0
	for head != nil {
	  if head.Val == 0 {
		if sum != 0 {
		  current.Next = &ListNode{Val: sum}
		  current = current.Next
		  sum = 0
		}
	  } else {
		sum += head.Val
	  }
	  head = head.Next
	}
	if sum != 0 {
	  current.Next = &ListNode{Val: sum}
	}
	return temp.Next
  }
