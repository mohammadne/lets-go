package basic

import "testing"

func TestSimpleChannel(t *testing.T) {
	simpleChannel()
}

func TestSynchChannel(t *testing.T) {
	synchChannel()
}

func TestNonBlocking(t *testing.T) {
	nonBlocking()
}

func TestSimulatePingPong(t *testing.T) {
	simulatePingPong()
}
