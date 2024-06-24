package p2p

// Peer Интерфейс, который представляет удалённый узел
//
// В контексте сетевого программирования, “peer” обычно относится к участнику сети,
// который взаимодействует с другими участниками на равных,
// а “удаленный узел” — это компьютер или устройство,
// находящееся в другом месте, с которым можно установить соединение через сеть.
// Но здесь Peer (так как он участник) будет являться удалённым узлом.
type Peer interface{}

// Transport — это все, что связано с связью между узлами (в данном случае Peer) в сети
// Он может быть TCP, UDP, websockets
//
// В контексте сетевого программирования, “транспорт” относится к механизмам,
// которые обеспечивают передачу данных между различными узлами в сети.
// Это может включать в себя протоколы, алгоритмы и другие технические решения,
// которые позволяют эффективно обмениваться информацией.
type Transport interface {
	ListenAndAccept() error
}
