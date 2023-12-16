package main

import (
	sync "sync"

	mnt "github.com/chenxiaomin-non/go-tail/mnt"
)

func main() {
	mng := mnt.GetObsManager()
	wg := sync.WaitGroup{}
	wg.Add(1)

	PrintLog(mng, &wg)

	defer wg.Wait()
}

func PrintLog(mng *mnt.ObsManager, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 5; i++ {
			msg := mng.PopMessage()
			println(msg)
		}
		wg.Done()
	}()
}
