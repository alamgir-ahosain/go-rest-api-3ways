package products

type MemStore struct{
	list map[string] Product
}

func NewMemStore() *MemStore{
	list:=make(map[string] Product)
	return & MemStore{
		list ,
	}
}
