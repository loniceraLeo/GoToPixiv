//local server for listening requests
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

func start() {

	list, er := tls.Listen("tcp", "127.0.0.1:443", config)
	log.Println("GoToPixiv is launched.Listening on 127.0.0.1:443")
	if er != nil {
		log.Fatal(er)
		return
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go func() {
			remote, err := tls.Dial("tcp", "example.com:443", config)
			if err != nil {
				log.Fatal(err)
				return
			}

			defer remote.Close()
			defer conn.Close()

			e := tunnel(conn, remote)
			if e != nil {
				return
			}
		}()
	}
}

//pipe the data through tunnel
func tunnel(l, r net.Conn) error {
	ch := make(chan error)

	go func() {
		_, err := io.Copy(l, r)
		l.SetDeadline(time.Now())
		r.SetDeadline(time.Now())
		ch <- err
	}()

	_, err := io.Copy(r, l)
	l.SetDeadline(time.Now())
	r.SetDeadline(time.Now())
	t := <-ch
	if t == nil {
		t = err
	}

	return t
}
