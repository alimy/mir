package container

// OnceFunc used after OnceSet add it
type OnceFunc func(string) error

// OnceSet once set
type OnceSet interface {
	Add(string) error
	Exist(string) bool
}
type onceSet struct {
	inSet    map[string]struct{}
	onceFunc OnceFunc
}

// Add add a item to set
func (s *onceSet) Add(it string) error {
	if _, exist := s.inSet[it]; !exist {
		err := s.onceFunc(it)
		if err != nil {
			return err
		}
		s.inSet[it] = struct{}{}
	}
	return nil
}

// Exist whether it exist
func (s *onceSet) Exist(it string) bool {
	_, exist := s.inSet[it]
	return exist
}

// New return an OnceSet instance
func NewOnceSet(onceFunc OnceFunc) OnceSet {
	return &onceSet{
		inSet:    make(map[string]struct{}),
		onceFunc: onceFunc,
	}
}
