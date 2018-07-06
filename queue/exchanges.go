package queue

const (
	// PinExchange is the name of the fanout exchange for regular ipfs pins
	PinExchange = "ipfs-pin"
	// ClusterPinExchange is the name of the fanout exchange for cluster ipfs pins
	ClusterPinExchange = "ipfs-cluster-pin"
)

// DeclareIPFSPinExchange is used to declare the exchange used to handle ipfs pins
func (qm *QueueManager) DeclareIPFSPinExchange() error {
	return qm.Channel.ExchangeDeclare(
		PinExchange, // name
		"fanout",    // type
		true,        // durable
		false,       // auto-delete
		false,       // internal
		false,       // no wait
		nil,         // args
	)
}

// DeclareIPFSClusterPinExchange is used to declare the exchange used to handle ipfs cluster pins
func (qm *QueueManager) DeclareIPFSClusterPinExchange() error {
	return qm.Channel.ExchangeDeclare(
		ClusterPinExchange, // name
		"fanout",           // type
		true,               // durable
		false,              // auto-delete
		false,              // internal
		false,              // no wait
		nil,                // args
	)
}