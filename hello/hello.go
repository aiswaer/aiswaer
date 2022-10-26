package hello

type Hello struct{}

func (h *Hello) Say() string {
	return "hello asiwaer!"
}

func (h *Hello) Sayv2() string {
	return "v2 hello asiwaer!"
}
