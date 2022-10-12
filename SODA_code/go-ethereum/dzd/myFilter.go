package dzd

type MyFilter struct {
	filter map[string]bool
}

func NewMyFilter() *MyFilter {
	mf := contracts
	return &MyFilter{
		filter: mf,
	}
}

func (mf *MyFilter) AddToFilter(addr string) {
	mf.filter[addr] = true
}

func (mf *MyFilter) Check(addr string) bool {
	return mf.filter[addr]
}

func (mf *MyFilter) Del(addr string) {
	delete(mf.filter, addr)
}
