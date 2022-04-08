package netutil

import "net"

func FindFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "[::]:0")
	if err != nil {
		return -1, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return -1, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	_ = l.Close()
	return port, nil
}
