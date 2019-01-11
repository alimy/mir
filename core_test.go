package mir

type site struct {
	v1       Group `mir:"v1"`
	index    Get   `mir:"/index/"`
	articles Get   `mir:"//{subdomain}.domain.com/articles/{category}/{id:[0-9]+}?{filter}&{pages}#GetArticles"`
}

func (h *site) Index() string {
	return "Index"
}

func (h *site) GetArticles() string {
	return "GetArticles"
}

type handlerFunc func() string

type simpleEngine struct {
	pathHandler map[string]handlerFunc
}

func (e *simpleEngine) Register(entries ...interface{}) error {
	for _, entry := range entries {
		if tagFields, err := TagFieldsFrom(entry); err == nil {
			for _, tagField := range tagFields {
				if tagField.Handler != nil {
					if h, ok := tagField.Handler.(func() string); ok {
						e.pathHandler[tagField.Path] = h
					}
				}
			}
		} else {
			return err
		}
	}
	return nil
}
