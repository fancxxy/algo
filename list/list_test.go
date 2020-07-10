package list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	expected := []interface{}{"a", "b", "c"}
	list := New(expected...)

	if actual := list.Len(); actual != len(expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}

	if actual := list.Values(); !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}
}

func TestAppend(t *testing.T) {
	expected := []interface{}{0, 1, 2, 3, 4, 5}
	list := New()
	list.Append(1, 3)
	list.PushFront(0)
	list.PushBack(5)
	list.InsertAfter(4, list.Find(3))
	list.InsertBefore(2, list.Find(3))

	if actual := list.Len(); actual != len(expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}

	if actual := list.Values(); !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}

	if actual := list.Front(); actual.Value.(int) != 0 {
		t.Errorf("actual %v expected %v", actual.Value.(int), 0)
	}

	if actual := list.Back(); actual.Value.(int) != 5 {
		t.Errorf("actual %v expected %v", actual.Value.(int), 5)
	}
}

func TestRemove(t *testing.T) {
	expected := []interface{}{1, 2, 3, 4}
	list := New([]interface{}{0, 1, 3, 2, 3, 4, 3, 5}...)

	if actual := list.Remove(list.FindLast(3)); actual != 3 {
		t.Errorf("actual %v expected %v", actual, 3)
	}
	if actual := list.Remove(list.Find(3)); actual != 3 {
		t.Errorf("actual %v expected %v", actual, 3)
	}
	if actual := list.PopFront(); actual != 0 {
		t.Errorf("actual %v expected %v", actual, 0)
	}
	if actual := list.PopBack(); actual != 5 {
		t.Errorf("actual %v expected %v", actual, 5)
	}

	if actual := list.Values(); !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}
}

func TestMove(t *testing.T) {
	expected := []interface{}{1, 3, 2}
	list := New(1, 2, 3)
	if actual := list.MoveToFront(nil); actual != nil {
		t.Errorf("actual %v, expected %v", actual, nil)
	}

	if actual := list.MoveToFront(list.Back()); actual.Value.(int) != 3 {
		t.Errorf("actual %v, expected %v", actual.Value.(int), 3)
	}

	if actual := list.MoveToBack(nil); actual != nil {
		t.Errorf("actual %v, expected %v", actual, nil)
	}

	if actual := list.MoveToBack(list.Back()); actual.Value.(int) != 2 {
		t.Errorf("actual %v, expected %v", actual.Value.(int), 2)
	}

	if actual := list.MoveToBack(list.Find(1)); actual.Value.(int) != 1 {
		t.Errorf("actual %v, expected %v", actual.Value.(int), 1)
	}

	if actual := list.MoveAfter(nil, list.Back()); actual != nil {
		t.Errorf("actual %v, expected %v", actual, nil)
	}

	if actual := list.MoveAfter(list.Find(2), list.Back()); actual.Value.(int) != 2 {
		t.Errorf("actual %v, expected %v", actual.Value.(int), 2)
	}

	if actual := list.MoveBefore(nil, list.Front()); actual != nil {
		t.Errorf("actual %v, expected %v", actual, nil)
	}

	if actual := list.MoveBefore(list.Find(1), list.Front()); actual.Value.(int) != 1 {
		t.Errorf("actual %v, expected %v", actual.Value.(int), 1)
	}

	if actual := list.Values(); !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}
}

func TestClear(t *testing.T) {
	list := New(1, 2, 3)
	list.Clear()
	if actual := list.Len(); actual != 0 {
		t.Errorf("actual %v expected %v", actual, 0)
	}

	if actual := list.Empty(); actual != true {
		t.Errorf("actual %v expected %v", actual, true)
	}

	if actual := list.Front(); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.Back(); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.Find(0); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.FindLast(0); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.InsertAfter(0, nil); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.InsertBefore(0, nil); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.PopFront(); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.PopBack(); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	if actual := list.Remove(nil); actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}
}

func TestRange(t *testing.T) {
	expected := []interface{}{1, 2, 3, 4, 5}
	list := New(expected...)

	var actual []interface{}
	for node := list.Front(); node != nil; node = node.Next() {
		actual = append(actual, node.Value)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}

	actual = nil
	for node := list.Back(); node != nil; node = node.Prev() {
		actual = append(actual, node.Value)
	}
	expected = []interface{}{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v expected %v", actual, expected)
	}
}
