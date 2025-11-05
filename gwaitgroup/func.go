// Package gwaitgroup
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-11-5 11:57
//
// --------------------------------------------
package gwaitgroup

import "sync"

// Go is a wrapper for sync.WaitGroup, providing a more convenient way to use sync.WaitGroup.Add and sync.WaitGroup.Done.
func Go(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
