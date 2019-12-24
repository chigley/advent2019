package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
	"github.com/chigley/advent2019/vector"
)

const (
	interfaces = 50
	queueSize  = 256
)

type packet vector.XY

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(program))
}

func Part1(program []int) int {
	packets := make([]chan packet, interfaces)
	for i := 0; i < interfaces; i++ {
		packets[i] = make(chan packet, queueSize)
	}

	var (
		natPacket packet
		natMutex  sync.RWMutex
	)
	go func() {
		ySeen := make(map[int]struct{})

		for {
			time.Sleep(1 * time.Second)
			for i := 0; i < interfaces; i++ {
				if len(packets[i]) > 0 {
					continue
				}
			}

			// idle
			natMutex.RLock()
			packetToSend := natPacket
			natMutex.RUnlock()
			if _, ok := ySeen[packetToSend.Y]; ok {
				log.Fatal(packetToSend.Y)
			}
			ySeen[packetToSend.Y] = struct{}{}
			log.Printf("Sending NAT packet with %#v", packetToSend)
			packets[0] <- packetToSend
		}
	}()

	ret := make(chan int)
	for i := 0; i < interfaces; i++ {
		inputs := make(chan int)
		out := intcode.New(program).RunInteractive(inputs, nil, intcode.NonBlockingInputPairs())
		inputs <- i

		// Pass packets we receive onto our machine as two separate inputs
		go func(i int) {
			for p := range packets[i] {
				inputs <- p.X
				inputs <- p.Y
			}
		}(i)

		// Pass packets we send onto the destination machine's queue
		go func() {
			for dest := range out {
				packet := packet{<-out, <-out}
				if 0 <= dest && dest < interfaces {
					packets[dest] <- packet
				} else if dest == 255 {
					natMutex.Lock()
					natPacket = packet
					natMutex.Unlock()
				}
			}
		}()
	}
	return <-ret
}
