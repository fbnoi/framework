package handler

func New[T any]() *Handler[T] {
	return &Handler[T]{
		idx: -1,
	}
}

type MD[T any] func(T, HandleFunc[T])
type HandleFunc[T any] func(T)

type Handler[T any] struct {
	mds      []MD[T]
	endpoint HandleFunc[T]

	idx int
}

func (h *Handler[T]) Then(mds ...MD[T]) *Handler[T] {
	h.mds = append(h.mds, mds...)
	return h
}

func (h *Handler[T]) Final(endpoint HandleFunc[T]) *Handler[T] {
	if h.endpoint != nil {
		panic("endpoint already set")
	}
	h.endpoint = endpoint
	return h
}

func (h *Handler[T]) Handle(t T) {
	h.nextHandleFunc()(t)
}

func (h *Handler[T]) nextHandleFunc() HandleFunc[T] {
	nextMd := h.nextMiddleware()
	return func(t T) {
		nextMd(t, h.nextHandleFunc())
	}
}

func (h *Handler[T]) nextMiddleware() MD[T] {
	h.idx++
	var md MD[T]
	if h.idx >= len(h.mds) {
		md = func(_t T, fn HandleFunc[T]) {
			if h.endpoint != nil {
				h.endpoint(_t)
			}
		}
	} else {
		md = h.mds[h.idx]
	}
	return md
}
