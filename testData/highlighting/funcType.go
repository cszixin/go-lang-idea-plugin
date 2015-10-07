package main

type demoStr struct {
    Field string
}

type f func() *demoStr

func DemoFunc() f {
    return func () *demoStr {
        return &demoStr{Field: "dmeo"}
    }
}

func main() {
    fun := DemoFunc()
    field := fun()
    _ = field.Field
}

type A struct {
}

func (self *A) chain() *A{
    return self
}

func getAnA() (result *A) {
    return &A{}
}

func _() {
    getAnA().chain()
}

func _(callback func() (interface{}, error)) {
	if _, err := callback(); err != nil {
		err.Error()
	}
}

func _(callback func() error) {
	if err := callback(); err != nil {
		err.Error()
	}
}

func _(fn func() (i int, e error)) {
	i, e :=  fn()
	print(i)
	print(e.Error())
}

func _(fn func() (myErr error)) {
	err := fn()
	err.Error()
}

func _(fn func() error) {
	err := fn()
	err.Error()
}