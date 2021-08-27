package ds

type List interface {
	Size() uint
	Clear()
	ForEach(func(uint, interface{}))
	Map(func(uint, interface{}) interface{})
	Filter(func(uint, interface{}) bool) List
	Insert(v interface{})
	InsertIn(v interface{}, index uint)
	ReplaceIn(v interface{}, index uint) interface{}
	Add(v interface{})
	RemoveFrom(index uint) interface{}
	Get(index uint) interface{}
	Head() interface{}
	Tail() interface{}
}
