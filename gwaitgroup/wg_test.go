// Package gwaitgroup
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-11-5 11:38
//
// --------------------------------------------
package gwaitgroup

import (
	"sync"
	"testing"
	"time"
)

func TestWgGo(t *testing.T) {
	var wg WaitGroup

	start := time.Now()
	wg.Go(func() {
		t.Log("wg go")
	})

	wg.Go(func() {
		time.Sleep(time.Second * 10)
		t.Log("wg go")
	})
	wg.Wait()

	now := time.Now()
	if !now.Add(-time.Second * 10).After(start) {
		t.Error("wg go error wait can not wait goroutine finish")
	}
}

func TestGo(t *testing.T) {
	var wg sync.WaitGroup

	start := time.Now()
	Go(&wg, func() {
		t.Log("sync.WaitGroup go")
	})

	Go(&wg, func() {
		time.Sleep(time.Second * 10)
		t.Log("sync.WaitGroup go")
	})
	wg.Wait()

	now := time.Now()
	if !now.Add(-time.Second * 10).After(start) {
		t.Error("sync.WaitGroup go error wait can not wait goroutine finish")
	}
}
