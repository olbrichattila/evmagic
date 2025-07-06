package contracts

type Replay interface {
	Replay(parentActionIdentifier string) (bool, error)
	Register(actionId string) error
	Store(parentActionIdentifier string, replayData []byte) error
}
