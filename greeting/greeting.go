package greeting

type Hoge struct {
  name string
}

func NewHoge(name string) *Hoge {
  return &Hoge{
    name: name,
  }
}

func (h *Hoge) Greeting() {
  fmt.Println("hello " + h.name + "!")
}
