package ds

type linkedListElement struct {
	next  *linkedListElement
	value interface{}
}

type LinkedList struct {
	size uint
	head *linkedListElement
	tail *linkedListElement
}

func (ll *LinkedList) Size() uint {
	return ll.size
}

func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList) ForEach(f func(interface{})) {
	for p := ll.head; p != nil; p = p.next {
		f(p.value)
	}
}

func (ll *LinkedList) Filter(f func(interface{}) bool) List {
	result := &LinkedList{}
	for p := ll.head; p != nil; p = p.next {
		if f(p.value) {
			result.Add(p.value)
		}
	}
	return result
}

func (ll *LinkedList) Add(value interface{}) {
	if ll.head == nil {
		ll.head = &linkedListElement{
			value: value,
		}
		ll.tail = ll.head
	} else {
		ll.tail.next = &linkedListElement{
			value: value,
		}
		ll.tail = ll.tail.next
	}
	ll.size++
}

func (ll *LinkedList) Insert(value interface{}) {
	ll.InsertIn(value, 0)
}

func (ll *LinkedList) InsertIn(value interface{}, index uint) {
	var previous *linkedListElement
	pivot := ll.head
	for i := uint(0); i < index; i++ {
		previous = pivot
		pivot = pivot.next
	}
	newElement := &linkedListElement{value: value}
	if ll.head == nil { // to reach here it means index is 0 and pivot is nil
		ll.head = newElement
		ll.tail = newElement
	} else if previous == nil {
		newElement.next = ll.head
		ll.head = newElement
	} else if pivot == nil { // to reach here means the head is not nil and it had gone through the list
		ll.tail.next = newElement
		ll.tail = newElement
	} else {
		newElement.next = pivot.next
		pivot.next = newElement
	}
	ll.size++
}

func (ll *LinkedList) ReplaceIn(value interface{}, index uint) interface{} {
	pivot := ll.head
	for i := uint(0); i < index; i++ {
		pivot = pivot.next
	}
	oldValue := pivot.value
	pivot.value = value
	return oldValue
}

func (ll *LinkedList) Get(index uint) interface{} {
	pivot := ll.head
	for i := uint(0); i < index; i++ {
		pivot = pivot.next
	}
	return pivot.value
}

func (ll *LinkedList) Head() interface{} {
	if ll.head != nil {
		return ll.head.value
	}
	return nil
}

func (ll *LinkedList) Tail() interface{} {
	if ll.tail != nil {
		return ll.tail.value
	}
	return nil
}

func (ll *LinkedList) RemoveFrom(index uint) interface{} {
	var previous *linkedListElement
	pivot := ll.head
	for i := uint(0); i < index; i++ {
		previous = pivot
		pivot = pivot.next
	}

	if previous == nil { // to reach here, index is 0
		ll.head = ll.head.next
		if ll.head == nil { // removed last element, clearing the tail too
			ll.tail = nil
		}
	} else if pivot == ll.tail {
		ll.tail = previous
		ll.tail.next = nil
	} else {
		previous.next = pivot.next
		pivot.next = nil
	}
	ll.size--
	return pivot.value
}
