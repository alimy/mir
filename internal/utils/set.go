package utils

type Set map[string]struct{}

func (s Set) Add(items ...string) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set) Exist(item string) bool {
	_, exist := s[item]
	return exist
}

func (s Set) List() []string {
	methods := make([]string, 0, len(s))
	for m := range s {
		methods = append(methods, m)
	}
	return methods
}
