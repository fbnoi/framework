package handler

func New[T any]() *Handler[T] {
	return &Handler[T]{
		idx: -1,
	}
}

type Handler[T any] struct {
	mds      []func(T, func(T))
	endpoint func(T)

	idx int
}

func (h *Handler[T]) Then(mds ...func(T, func(T))) *Handler[T] {
	h.mds = append(h.mds, mds...)
	return h
}

func (h *Handler[T]) Final(endpoint func(T)) *Handler[T] {
	if h.endpoint != nil {
		panic("endpoint already set")
	}
	h.endpoint = endpoint
	return h
}

func (h *Handler[T]) Handle(t T) {
	h.nextHandleFunc()(t)
}

func (h *Handler[T]) nextHandleFunc() func(T) {
	nextMd := h.nextMiddleware()
	return func(t T) {
		nextMd(t, h.nextHandleFunc())
	}
}

func (h *Handler[T]) nextMiddleware() func(T, func(T)) {
	h.idx++
	var md func(T, func(T))
	if h.idx >= len(h.mds) {
		md = func(_t T, fn func(T)) {
			if h.endpoint != nil {
				h.endpoint(_t)
			}
		}
	} else {
		md = h.mds[h.idx]
	}
	return md
}
