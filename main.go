package main

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"log"
	"time"
)

const (
	startIP = "10.100.207."
)

func main() {
	fmt.Println("Hello World")
	ends := [10]string{"40", "41", "42", "43", "30", "31", "32", "33", "20", "21"}
	for _, item := range ends {
		err := pingIp(startIP + item)
		if err != nil {
			log.Println("Ping IP "+startIP+item+" failed: ", err)
		}
	}
}
func pingIp(ip string) error {

	stop := make(chan struct{})
	pngr, err := probing.NewPinger(ip)
	if err != nil {
		return err
	}
	pngr.SetPrivileged(true)
	pngr.Interval = 2 * time.Second
	pngr.Size = 548
	pngr.OnRecv = func(pkt *probing.Packet) {
		logMessage := fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v", pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		log.Println(logMessage)
		pngr.Stop()
		stop <- struct{}{}
	}
	pngr.OnSend = func(pkt *probing.Packet) {
		logMessage := fmt.Sprintf("ping " + pkt.Addr + " ...")
		log.Println(logMessage)
	}

	go func() {
		err := pngr.Run()
		if err != nil {
			log.Fatal("pinger can't start" + err.Error() + ", exit status 1")
		}
	}()
	select {
	case <-stop:
		{
			log.Println("stop ping")
			return nil
		}
	}
	return nil
}
