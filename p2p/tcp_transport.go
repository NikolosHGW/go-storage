package p2p

import (
	"fmt"
	"net"
)

// TCPPeer представляет собой удалённый узел в сети,
// связанный с другим узлом через установленное TCP-соединение.
type TCPPeer struct {
	conn net.Conn

	// если мы устанавливаем и получаем соединение => outbound == true (соединение инициировано исходящим)
	// если мы принимаем и получаем соединение => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listener      net.Listener
	listenAddress string

	// mu    sync.RWMutex
	// peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddress)

	if err != nil {
		return fmt.Errorf("net.Listen(tcp, ...) doesnt work: %w", err)
	}

	t.listener = ln

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.Printf("new incoming connection %+v\n", peer)
}
