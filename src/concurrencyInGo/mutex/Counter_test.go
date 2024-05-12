package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter_Incr(t *testing.T) {
	type fields struct {
		CountType int
		Name      string
		mu        sync.Mutex
		count     uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var counter Counter
			var wg sync.WaitGroup
			wg.Add(10)

			for i := 0; i < 10; i++ {
				go func() {
					defer wg.Done()
					
					for j := 0; j < 10000; j++ {
						counter.Incr()
					}
				}()
			}

			wg.Wait()
			fmt.Println(counter.Count())
		})
	}
}
