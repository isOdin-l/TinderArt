package entities

import "github.com/google/uuid"

type Swipe struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	TargetId  uuid.UUID
	Decision1 *bool
	Decision2 *bool
}

type BrokerMatchMessage struct {
	User1ID uuid.UUID
	User2ID uuid.UUID
}
