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
	if ll.head == nil {
		ll.head = &linkedListElement{
			value: value,
		}
		ll.tail = ll.head
	} else {
		ll.head = &linkedListElement{
			next:  ll.head,
			value: value,
		}
	}
	ll.size++
}

func (ll *LinkedList) Get(index int) interface{} {
	pivot := ll.head
	for i := 0; i < index; i++ {
		pivot = pivot.next
	}
	return pivot.value
}