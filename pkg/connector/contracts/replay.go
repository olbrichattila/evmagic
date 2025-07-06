package contracts

type Replay interface {
	Replay(actionId string) (bool, error)
	Store(parentActionIdentifier string, replayData []byte) error
}
