package unbounded

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMakeInfiniteNoPause(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range out {
			vi := v.(int)
			if lastVal+1 != vi {
				t.Errorf("Unexpected value; expected %d, got %d", lastVal+1, vi)
			}
			lastVal = vi
		}
		wg.Done()
		fmt.Println("finished reading")
	}()
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
	fmt.Println("finished writing")
	wg.Wait()

	if lastVal != 99 {
		t.Errorf("Didn't get all values, last one received was %d", lastVal)
	}
}

func TestMakeInfiniteSlowRead(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range out {
			time.Sleep(50*time.Millisecond)
			vi := v.(int)
			if lastVal+1 != vi {
				t.Errorf("Unexpected value; expected %d, got %d", lastVal+1, vi)
			}
			lastVal = vi
		}
		wg.Done()
		fmt.Println("finished reading")
	}()
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
	fmt.Println("finished writing")
	wg.Wait()

	if lastVal != 99 {
		t.Errorf("Didn't get all values, last one received was %d", lastVal)
	}
}

func TestMakeInfiniteSlowWrite(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range out {
			vi := v.(int)
			if lastVal+1 != vi {
				t.Errorf("Unexpected value; expected %d, got %d", lastVal+1, vi)
			}
			lastVal = vi
		}
		wg.Done()
		fmt.Println("finished reading")
	}()
	for i := 0; i < 100; i++ {
		time.Sleep(50*time.Millisecond)
		in <- i
	}
	close(in)
	fmt.Println("finished writing")
	wg.Wait()

	if lastVal != 99 {
		t.Errorf("Didn't get all values, last one received was %d", lastVal)
	}
}
