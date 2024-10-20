package writer

import (
	"fmt"
	"io"
	"net"
)

var _writer io.Writer = (*TBLWriter)(nil)

type TBLWriter struct {
	conn *net.UDPConn
}

func NewTblWriter(addr string) (*TBLWriter, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		return nil, err
	}

	return &TBLWriter{
		conn: conn,
	}, nil

	//
	//data, err := bufio.NewReader(conn).ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Print("> ", string(data))

}

func (w *TBLWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("write to buff: %v", string(p))
	return w.conn.Write(p)
}
