- interface (pointer receiver doesnt implement interface)?

```
type FooI interface {
  bar()
}

type Foo struct {

}
// Why doesn't this satisfy FooI interface?
func (foo *Foo) bar() {

}
```

- channels (what is buffered channels)
