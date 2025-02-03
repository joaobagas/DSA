package dstr_test

import (
	"DSA/dstr"
	"strconv"
	"testing"
)

func TestLinkedList(t *testing.T) {
	childEnd := dstr.CreateLinkedList[string]("Child2", nil)
	childMiddle := dstr.CreateLinkedList[string]("Child1", &childEnd)
	head := dstr.CreateLinkedList[string]("Head", &childMiddle)

	if dstr.GetData(head) != "Head" {
		t.Fatalf("There is a problem with creating nodes!")
	}

	childAfterEnd := dstr.CreateLinkedList[string]("Child3", nil)
	dstr.AddNode(&head, &childAfterEnd)

	count := 0
	for {
		curr := head
		count++
		if dstr.IsNextNode(curr) {
			curr = dstr.NextNode(curr)
			if dstr.GetData(curr) != "Child"+strconv.Itoa(count) {
				t.Fatalf("Error with adding children!1")
			}
		} else {
			if count != 3 {
				t.Fatalf("Error with adding children!2")
			}
			break
		}
	}
}
