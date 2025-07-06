package replay

import (
	"github.com/olbrichattila/evmagic/pkg/connector/contracts"
	memoryReplay "github.com/olbrichattila/evmagic/pkg/replay/memory-replay"
)

func New(publisher contracts.Publisher) contracts.Replay {
	// TODO create factory to create replay by config or env (switch)

	return memoryReplay.New(publisher)
}
