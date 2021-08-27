package ds

type List interface {
	Size() uint
	Clear()
	ForEach(func(interface{}))
	Filter(func(interface{}) bool) List
	Insert(v interface{})
	Add(v interface{})
	Get(index int) interface{}
	Head() interface{}
	Tail() interface{}
}
