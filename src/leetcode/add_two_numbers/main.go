package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (resultNode *ListNode) {
	var addValue int = 0
	var curNode *ListNode
	for {
		if l1 == nil && l2 == nil && addValue == 0{
			return
		}

		var (
			l1Value int = 0
			l2Value int = 0
		)

		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}

		addValue = addValue + l1Value + l2Value
		val := addValue % 10
		addValue = addValue / 10

		newNode := &ListNode{
			Val: val,
		}

		if curNode != nil {
			curNode.Next = newNode
			curNode = newNode

		} else {
			resultNode = newNode
			curNode = newNode
		}

	}

	return 
}

func main()  {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}

	resultNode := AddTwoNumbers(l1, l2)
	for resultNode != nil {
		fmt.Println(resultNode.Val)
		resultNode = resultNode.Next
	}
}