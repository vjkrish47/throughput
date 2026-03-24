import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	clients   = 50
	replicas  = 5
	requests  = 2000
)

type LoadBalancer struct {
	index int
	mu    sync.Mutex
}

func (lb *LoadBalancer) selectReplica() int {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	r := lb.index
	lb.index = (lb.index + 1) % replicas
	return r
}

func processRequest() {
	time.Sleep(time.Millisecond)
}

func client(lb *LoadBalancer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		lb.selectReplica()
		processRequest()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	lb := &LoadBalancer{}
	start := time.Now()

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go client(lb, &wg)
	}

	wg.Wait()
	fmt.Println("Execution Time:", time.Since(start))
}

