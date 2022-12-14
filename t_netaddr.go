package arbitrary

import (
	"fmt"
	"net"
)

// NetAddr returns an arbitrary net.Addr.
func (arb T) NetAddr() net.Addr {

	var network string
	{
		if arb.Bool() {
			network = "tcp"
		} else {
			network = "udp"
		}
	}

	var port uint16
	{
		port = uint16(arb.randomness.Intn(65536))

		if 0 == arb.randomness.Intn(5) {
			port = uint16(arb.randomness.Intn(1024))
		}

		if 0 == arb.randomness.Intn(4) {
			port = uint16(arb.randomness.Intn(100))
		}

		if 0 == arb.randomness.Intn(30) {
			port = 80
		}
		if 0 == arb.randomness.Intn(40) {
			port = 79
		}
	}

	var value string
	func(){

		value = fmt.Sprintf("%d.%d.%d.%d:%d", arb.randomness.Intn(256), arb.randomness.Intn(256), arb.randomness.Intn(256), arb.randomness.Intn(256), port)

		if 0 == arb.randomness.Intn(16) {
			value = fmt.Sprintf("127.0.0.1:%d", port)
			return
		}

		if 0 == arb.randomness.Intn(16) {
			value = fmt.Sprintf("192.168.0.1:%d", port)
			return
		}

		if 0 == arb.randomness.Intn(7) {
			value = fmt.Sprintf("%d.%d:%d", arb.randomness.Intn(256), arb.randomness.Intn(16777216), port)
			return
		}

		if 0 == arb.randomness.Intn(7) {
			value = fmt.Sprintf("%d.%d.%d:%d", arb.randomness.Intn(256), arb.randomness.Intn(256), arb.randomness.Intn(65536), port)
			return
		}
	}()

	return internalNetAddr{
		network: network,
		value: value,
	}
}

type internalNetAddr struct {
	network string
	value string
}

var _ net.Addr = internalNetAddr{}

func (receiver internalNetAddr) Network() string {
	return receiver.network
}

func (receiver internalNetAddr) String() string {
	return receiver.value
}
