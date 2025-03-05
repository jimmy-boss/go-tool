// Package go_tool
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-3-5 10:50
//
// --------------------------------------------
package go_tool

import (
	"github.com/sasha-s/go-deadlock"
	"testing"
)

func TestDealLock(t *testing.T) {
	var mu deadlock.RWMutex
	mu.Lock()
	defer mu.Unlock()

	mu.RLock()
	defer mu.RUnlock()
}
