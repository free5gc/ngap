package ngapSctp_test

import (
	"encoding/binary"
	"net"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"git.cs.nctu.edu.tw/calee/sctp"

	"free5gc/lib/ngap/logger"
	"free5gc/lib/ngap/ngapSctp"
)

var testClientNum = 2

func TestSCTP(t *testing.T) {
	runtime.GOMAXPROCS(10)

	listenConn := ngapSctp.Server("127.0.0.1")
	go func() {
		for {
			readChan := make(chan ngapSctp.ConnData, 1024)
			conn, err := ngapSctp.Accept(listenConn)
			assert.True(t, err == nil)
			logger.NgapLog.Printf("SCTP Accept from: %s", conn.RemoteAddr().String())
			go ngapSctp.Start(conn, readChan)
			go sendSCTPMsg(conn, readChan)
		}
	}()
	time.Sleep(10 * time.Millisecond)
	for i := 0; i < testClientNum; i++ {
		go func() {
			testClient()
		}()
		time.Sleep(100 * time.Millisecond)

	}
	// time.Sleep(5000 * time.Microsecond)

	if err := ngapSctp.Destroy(listenConn); err != nil {
		logger.NgapLog.Fatal("Cannot close Listener listenConn")
	}

	s2 := ngapSctp.Server("127.0.0.1")
	if err := ngapSctp.Destroy(s2); err != nil {
		logger.NgapLog.Fatal("Cannot close Listener s2")
	}
}

func sendSCTPMsg(conn net.Conn, readChan chan ngapSctp.ConnData) {
	for {
		tbuff := <-readChan
		buff := []byte("I'm fine. Thank you!")
		if len(tbuff.GetData()) <= 0 {
			continue
		}
		// ngapSctp.SendMsg(conn, append(tbuff[:32], buff...))
		err := ngapSctp.SendMsg(conn, buff)
		if err != nil {
			break
		}
	}
}

func testClient() {
	ipStr := "127.0.0.1"
	ips := []net.IPAddr{}
	if ip, err := net.ResolveIPAddr("ip", ipStr); err != nil {
		logger.NgapLog.Errorf("Error resolving address '%s': %v", ipStr, err)
	} else {
		ips = append(ips, *ip)
	}
	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    38412,
	}
	logger.NgapLog.Printf("raw addr: %+v\n", addr.ToRawSockAddrBuf())

	var laddr *sctp.SCTPAddr
	conn, err := sctp.DialSCTP("sctp", laddr, addr)
	if err != nil {
		logger.NgapLog.Fatalf("failed to dial: %v", err)
	}
	logger.NgapLog.Printf("Dail LocalAddr: %s; RemoteAddr: %s", conn.LocalAddr(), conn.RemoteAddr())
	for {
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, 60)
		ppid := binary.LittleEndian.Uint32(bs)
		info := &sctp.SndRcvInfo{
			Stream: uint16(ppid),
			PPID:   uint32(ppid),
		}
		err := conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)
		if err != nil {
			logger.NgapLog.Fatalf("Connection Error %v", err)
		}
		msg := "Hello, how are you?"
		n, err := conn.SCTPWrite([]byte(msg), info)
		if err != nil {
			logger.NgapLog.Fatalf("failed to write: %v", err)
		}
		logger.NgapLog.Printf("write: %d, %s", n, msg)
		buf := make([]byte, 254)
		_, info, err = conn.SCTPRead(buf)
		if err != nil {
			logger.NgapLog.Fatalf("failed to read: %v", err)
			return
		}
		logger.NgapLog.Printf("read: info: %+v, %s", info, buf)
		time.Sleep(300 * time.Microsecond)
	}
}
