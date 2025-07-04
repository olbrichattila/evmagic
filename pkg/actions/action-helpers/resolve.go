package actionHelper

import "github.com/olbrichattila/evmagic/pkg/helpers"

type SnsAction struct {
	TopicArn string
	Subject  string
	Message  string
}

func ToSNSAction(data []byte) (SnsAction, error) {
	return helpers.ToStruct[SnsAction](data)
}
