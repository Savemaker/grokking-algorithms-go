package main

import (
	"strconv"
	"testing"
)

func TestListAdd(t *testing.T) {
	list := &LinkedList{}

	list.Add(10)
	list.Add(5)

	expected := "105"
	got := ""

	head := list.Head
	for head != nil {
		got += strconv.Itoa(head.Val)
		head = head.Next
	}
	if got != expected {
		t.Errorf("got %s wanted %s", got, expected)
	}
}
func TestRemove(t *testing.T) {
	list := &LinkedList{}

	list.Add(10)
	list.Add(5)
	list.Remove(10)

	expected := "5"
	got := ""

	head := list.Head
	for head != nil {
		got += strconv.Itoa(head.Val)
		head = head.Next
	}
	if got != expected {
		t.Errorf("got %s wanted %s", got, expected)
	}
}
