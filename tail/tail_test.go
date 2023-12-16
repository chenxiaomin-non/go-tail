package tail_test

import (
	"sync"
	"testing"

	"github.com/chenxiaomin-non/go-tail/tail"
)

func TestNewObserver(t *testing.T) {
	// new observer
	obs, err := tail.NewObserver("test.log")
	if err != nil {
		t.Fatal(err)
	}

	// set mode
	err = obs.SetReadFromTail()
	if err != nil {
		t.Fatal(err)
	}

	// get new content
	err = obs.Start()
	// read new content from channel

	if err != nil {
		t.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 3; i++ {
			t.Log(<-obs.Publisher)
		}
		wg.Done()
	}()

	if r := recover(); r != nil {
		t.Fatal(r)
	}

	wg.Wait()
	// test ok
	defer t.Log("test ok")
}
