package box

import "time"

var _ Record = (*Item)(nil)

type Item struct {
	key          string
	value        interface{}
	lastUpdated  time.Time
	creationTime time.Time
	setTTL       time.Duration
	metadata     map[string]interface{}
}

func (i *Item) Key() string {
	return i.key
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) LastUpdated() time.Time {
	return i.lastUpdated
}

func (i *Item) CreationTime() time.Time {
	return i.creationTime
}

func (i *Item) TTL() time.Duration {
	if i.setTTL < 0 {
		return i.setTTL
	}

	remainingTime := i.setTTL - time.Since(i.creationTime)
	if remainingTime < 0 {
		remainingTime = 0
	}

	return remainingTime
}

func (i *Item) Metadata() map[string]interface{} {
	return i.metadata
}