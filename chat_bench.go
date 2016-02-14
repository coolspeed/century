package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	targetAddr  = flag.String("a", "127.0.0.1:6666", "target chat server address")
	testMsgLen  = flag.Int("l", 8, "test message length")
	testConnNum = flag.Int("c", 3, "test connection number")
	testSeconds = flag.Int("t", 5, "test duration in seconds")
)

func main() {
	flag.Parse()

	var (
		outNum uint64
		inNum  uint64
		stop   uint64
	)

	msg := make([]byte, *testMsgLen)
    for i := 0; i < (*testMsgLen - 1); i++ {
        msg[i] = 'a'
    }
    msg[*testMsgLen - 1] = '\n'
    
    go func() {  // goroutine for timer and stoper
		time.Sleep(time.Second * time.Duration(*testSeconds))
		atomic.StoreUint64(&stop, 1)
	}()
    
    wg := new(sync.WaitGroup)

	for i := 0; i < *testConnNum; i++ {
		wg.Add(1)  // reader & writer.
		
		l := len(msg)
		recv := make([]byte, l)

		if conn, err := net.DialTimeout("tcp", *targetAddr, time.Minute*99999); err == nil {
			go func() {  // goroutine: reader
				for {
					for rest := l; rest > 0 ; {
						i, err := conn.Write(msg);
						rest -= i
						if err != nil {
							log.Println(err)
							break
						}
					}
	
					atomic.AddUint64(&outNum, 1)
	
					if atomic.LoadUint64(&stop) == 1 {
						break
					}
				}
				wg.Done()  // wait only reader goroutines
			}()
			
			// goroutine: writer
			go func() {
				for {
					for rest := l; rest > 0 ; {
						i, err := conn.Read(recv)
						rest -= i
						if err != nil {
							log.Println(err)
							break
						}
					}

					atomic.AddUint64(&inNum, 1)

					if atomic.LoadUint64(&stop) == 1 {
						break
					}
				}
//				wg.Done()
			}()
		} else {
			log.Println(err)
		}
	}
	
	wg.Wait()

	fmt.Println("Benchmarking:", *targetAddr)
	fmt.Println(*testConnNum, "clients, running", *testMsgLen, "bytes,", *testSeconds, "sec.")
	fmt.Println()
	fmt.Println("Speed:", outNum/uint64(*testSeconds), "request/sec,", inNum/uint64(*testSeconds), "response/sec")
	fmt.Println("Requests:", outNum)
	fmt.Println("Responses:", inNum)
}

