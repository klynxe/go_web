package dicorator

import (
	"net/http"
)

type decorHandle func(w http.ResponseWriter, r *http.Request)

type DecoratorFunc func(decorHandle) decorHandle

type decorator struct {
	decorators []DecoratorFunc
}

func NewDecorator(args ...DecoratorFunc) decorator {
	d := decorator{
		make([]DecoratorFunc, 0, len(args)),
	}
	for _, df := range args {
		d.decorators = append(d.decorators, df)
	}
	return d
}

//change original decorator
func (d *decorator) Decor(handle decorHandle) decorHandle {
	for _, arg := range d.decorators {
		handle = arg(handle)
	}
	return handle
}

//change original decorator
func (d *decorator) addDecoratorFunc(args ...DecoratorFunc) {
	for _, df := range args {
		d.decorators = append(d.decorators, df)
	}
}

//not change previous decoratoe
func (d decorator) appendDecoratorFunc(args ...DecoratorFunc) decorator {
	for _, df := range args {
		d.decorators = append(d.decorators, df)
	}
	return d
}

func Decorate(handle decorHandle, args ...DecoratorFunc) decorHandle {
	for _, arg := range args {
		handle = arg(handle)
	}
	return handle
}

func AddHeaderFabric(headerKey, headerValue string) DecoratorFunc {
	return func(dec decorHandle) decorHandle {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(headerKey, headerValue)
			dec(w, r)
		}
	}
}

func AddHeader(handle decorHandle) decorHandle {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handle(w, r)
	}
}

/*func test() {
	_ = Decorate(handle.SignUp, AddHeader, AddHeaderFabric("Content-Type", "application/json"))
}*/
