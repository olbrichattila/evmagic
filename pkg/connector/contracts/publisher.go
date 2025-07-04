package contracts

type Publisher interface {
	Publish(topic string, msg []byte) error
}
