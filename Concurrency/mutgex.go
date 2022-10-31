// Program with race condition fixed by mutex
package main
import (
	"fmt"
	"sync" // to import sync later on
)
var GFG = 0


func worker(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	GFG = GFG + 1
	m.Unlock()

	wg.Done()
}
func main() {
	var w sync.WaitGroup
    var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)		
		go worker(&w,&mu)
	}

	w.Wait()
	fmt.Println("Value of x", GFG)
}

