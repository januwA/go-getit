package getit_test

import (
	"testing"

	"github.com/januwA/go-getit"
)

type iService interface {
	New() iService
}

type foo struct{}

func (my *foo) New() iService {
	return my
}
func (my *foo) ById() string {
	return Get(&bar{}).ByName()
}

type bar struct{}

func (my *bar) New() iService {
	return my
}

func (my *bar) ByName() string {
	return "getit"
}

var service *getit.Getit[iService]

func Get[T any](serv T) T {
	return service.Get(serv).(T)
}

func TestGetit(t *testing.T) {
	service = new(getit.Getit[iService]).New(false)
	service.Register(new(foo), new(bar))

	foo_serv := new(foo)
	bar_serv := new(bar)

	if got := Get(foo_serv); got != foo_serv {
		t.Errorf("got %T want %T", got, foo_serv)
	}

	if got := Get(bar_serv); got != bar_serv {
		t.Errorf("got %T want %T", got, bar_serv)
	}

	if got := Get(foo_serv).ById(); got != "getit" {
		t.Errorf("got %T want %T", got, "getit")
	}

	// size := 10000
	// wg := new(sync.WaitGroup)
	// wg.Add(size)
	// for i := 0; i < size; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		Get(serv)
	// 	}()
	// }
	// wg.Wait()
}
