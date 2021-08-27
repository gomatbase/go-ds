package ds

type List interface {
	Size() uint
	Clear()
	ForEach(func(interface{}))
	Filter(func(interface{}) bool) List
	Insert(v interface{})
	InsertIn(v interface{}, index uint)
	ReplaceIn(v interface{}, index uint) interface{}
	Add(v interface{})
	RemoveFrom(index uint) interface{}
	Get(index uint) interface{}
	Head() interface{}
	Tail() interface{}
}
