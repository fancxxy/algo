package list

// Node 双链表结点
type Node struct {
	Value interface{}
	list  *List
	next  *Node
	prev  *Node
}

// Next 返回node的后继结点
func (node *Node) Next() *Node {
	if next := node.next; node.list != nil && node.list.head != next {
		return next
	}
	return nil
}

// Prev 返回node的前驱结点
func (node *Node) Prev() *Node {
	if prev := node.prev; node.list != nil && node.list.head != prev {
		return prev
	}
	return nil
}

// List 双链表
type List struct {
	head *Node
	len  int
}

// Len 返回链表长度
func (list *List) Len() int {
	return list.len
}

// Empty 判断链表是否为空
func (list *List) Empty() bool {
	return list.len == 0
}

// Clear 清空链表
func (list *List) Clear() {
	list.head.next = list.head
	list.head.prev = list.head
	list.len = 0
}

// Front 返回首元素
func (list *List) Front() *Node {
	if list.len == 0 {
		return nil
	}
	return list.head.next
}

// Back 返回尾元素
func (list *List) Back() *Node {
	if list.len == 0 {
		return nil
	}
	return list.head.prev
}

// Find 查找value结点，没找到返回nil
func (list *List) Find(value interface{}) *Node {
	node := list.head.next
	for node != list.head && node.Value != value {
		node = node.next
	}

	if node == list.head {
		return nil
	}
	return node
}

// FindLast 反向查找value结点，没找到返回nil
func (list *List) FindLast(value interface{}) *Node {
	node := list.head.prev
	for node != list.head && node.Value != value {
		node = node.prev
	}

	if node == list.head {
		return nil
	}
	return node
}

// Append 添加到尾元素
func (list *List) Append(values ...interface{}) {
	for _, value := range values {
		list.insert(value, list.head.prev)
	}
}

// PushFront 插入到首元素
func (list *List) PushFront(value interface{}) *Node {
	return list.insert(value, list.head)
}

// PushBack 插入到尾元素
func (list *List) PushBack(value interface{}) *Node {
	return list.insert(value, list.head.prev)
}

// InsertAfter 把value插入到at结点之后，返回新插入的结点
func (list *List) InsertAfter(value interface{}, at *Node) *Node {
	if at == nil || at.list != list {
		return nil
	}

	return list.insert(value, at)
}

// InsertBefore 把value插入到at结点之前，返回新插入的结点
func (list *List) InsertBefore(value interface{}, at *Node) *Node {
	if at == nil || at.list != list {
		return nil
	}

	return list.insert(value, at.prev)
}

func (list *List) insert(value interface{}, at *Node) *Node {
	node := &Node{Value: value, list: list}
	next := at.next

	node.next = next
	next.prev = node
	at.next = node
	node.prev = at
	list.len++
	return node
}

// PopFront 移除首元素
func (list *List) PopFront() interface{} {
	if list.len == 0 {
		return nil
	}
	return list.remove(list.head.next)
}

// PopBack 移除尾元素
func (list *List) PopBack() interface{} {
	if list.len == 0 {
		return nil
	}
	return list.remove(list.head.prev)
}

// Remove 移除node结点，返回结点元素
func (list *List) Remove(node *Node) interface{} {
	if node == nil || node == list.head || node.list != list {
		return nil
	}

	return list.remove(node)
}

func (list *List) remove(node *Node) interface{} {
	next, prev := node.next, node.prev
	next.prev = prev
	prev.next = next
	list.len--
	node.next = nil
	node.prev = nil
	node.list = nil
	return node.Value
}

// MoveToFront 移动结点到链表头
func (list *List) MoveToFront(node *Node) *Node {
	if node == nil || node.list != list {
		return nil
	}

	return list.move(node, list.head)
}

// MoveToBack 移动结点到链表尾
func (list *List) MoveToBack(node *Node) *Node {
	if node == nil || node.list != list {
		return nil
	}

	return list.move(node, list.head.prev)
}

// MoveAfter 移动node到at之后
func (list *List) MoveAfter(node, at *Node) *Node {
	if node == nil || at == nil || node.list != list || at.list != list {
		return nil
	}

	return list.move(node, at)
}

// MoveBefore 移动node到at之前
func (list *List) MoveBefore(node, at *Node) *Node {
	if node == nil || at == nil || node.list != list || at.list != list {
		return nil
	}

	return list.move(node, at.prev)
}

func (list *List) move(node, at *Node) *Node {
	if node == at {
		return node
	}

	// 删除node结点
	next, prev := node.next, node.prev
	next.prev = prev
	prev.next = next

	// 插入node结点到at之后
	next = at.next
	at.next = node
	node.prev = at
	node.next = next
	next.prev = node

	return node
}

// Values 获取全部元素的值，返回slice
func (list *List) Values() []interface{} {
	node := list.head.next
	slice := make([]interface{}, list.len)
	for i := 0; i < len(slice); i++ {
		slice[i] = node.Value
		node = node.next
	}
	return slice
}

// New 创建双链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = &Node{list: list}
	list.head.next = list.head
	list.head.prev = list.head
	if len(args) != 0 {
		for _, arg := range args {
			list.PushBack(arg)
		}
	}
	return list
}
