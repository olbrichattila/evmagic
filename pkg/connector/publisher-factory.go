package connector

import "github.com/olbrichattila/evmagic/pkg/connector/contracts"

type publishers struct {
	registered map[QueueType]contracts.Publisher
}

var publisherCache *publishers

func Publisher(qt QueueType) (contracts.Publisher, error) {
	if publisherCache == nil {
		publisherCache = &publishers{
			registered: make(map[QueueType]contracts.Publisher, 0),
		}
	}

	if p, ok := publisherCache.registered[qt]; ok {
		return p, nil
	}

	pb, err := getPublisher(qt)
	if err != nil {
		return nil, err
	}

	publisherCache.registered[qt] = pb

	return pb, nil
}
