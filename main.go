package main

import (
	"io"
	"os"

	"github.com/armon/go-socks5"
)

func main() {
	if os.Args[1] == "s" {
		lg.trace("listen kcp on port ", 64321)
		kl := kcpRawListen("0.0.0.0", 64321, "demo key", "demo salt")
		for kc := range kl.accept() {
			ss := smuxServerWrapper(kc)
			for sc := range ss.accept() {
				go func(sc *smuxServerSideConnection) {
					conf := &socks5.Config{}
					server, err := socks5.New(conf)
					if err != nil {
						lg.trace("error while creating socks5 server", err)
					}
					server.ServeConn(sc.stream)
					sc.close()
				}(sc)
			}
		}
	} else {
		lg.trace("listen tcp on port 9999")
		tl := tcpListen("0.0.0.0", 9999)
		kc := kcpRawConnect(os.Args[2], 64321, "demo key", "demo salt")
		ss := smuxClientWrapper(kc)
		for tc := range tl.accept() {
			sc := ss.connect()
			go func(tc *tcpServerSideConn) {
				io.Copy(sc.stream, tc.conn)
			}(tc)

			io.Copy(tc.conn, sc.stream)

			sc.close()
			tc.close()
		}
	}
}
