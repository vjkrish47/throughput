import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	clients   = 50
	replicas  = 5
	requests  = 2000
)

type Replica struct {
	capacity int64
}

type LoadBalancer struct {
	replicas []Replica
}

func (lb *LoadBalancer) selectReplica() int {
	var best int
	var maxCap int64 = -1
	for i, r := range lb.replicas {
		c := atomic.LoadInt64(&r.capacity)
		if c > maxCap {
			maxCap = c
			best = i
		}
	}
	return best
}

func processRequest(r *Replica) {
	atomic.AddInt64(&r.capacity, -1)
	time.Sleep(time.Millisecond)
	atomic.AddInt64(&r.capacity, 1)
}

func client(lb *LoadBalancer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		idx := lb.selectReplica()
		processRequest(&lb.replicas[idx])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	lb := &LoadBalancer{
		replicas: make([]Replica, replicas),
	}

	for i := 0; i < replicas; i++ {
		lb.replicas[i].capacity = int64(rand.Intn(5) + 5)
	}

	start := time.Now()

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go client(lb, &wg)
	}

	wg.Wait()
	fmt.Println("Execution Time:", time.Since(start))
}
