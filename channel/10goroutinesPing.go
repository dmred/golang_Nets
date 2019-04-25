package main

import (
    "github.com/sparrc/go-ping"
	"fmt"
)


func kek(c chan string) {
	
	pinger, err := ping.NewPinger("nigma.eu")
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}

		pinger.OnFinish = func(stats *ping.Statistics) {
			fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
			fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}

		fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
		
		pinger.Run()
}


func main (){
	var c chan string = make(chan string)
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);
	go kek(c);

	var input string
    fmt.Scanln(&input)
}
