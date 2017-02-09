package providers

import (
	"time"

	EventDb "../db"
)

/*
EventProvider type
*/
type DictionaryProvider struct {
	dbp EventDb.EventDbProvider
}

/*
insert event into mongo db
*/
func (dp *DictionaryProvider) GetDictionaryList(event EventTypes.Event) bool {
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now().UTC()
	}
	return ep.dbp.DbInsertEvent(event)
}

/*
/get full event list
*/
func (dp *DictionaryProvider) GetDictionaryDescription(dicCode string, ver string) ([]EventTypes.Event, error) {
	return ep.dbp.DbGetEventList()
}

/*
get event by id
*/
func (dp *DictionaryProvider) GetDictionary(dicCode string, ver string) (EventTypes.Event, error) {
	return ep.dbp.DbGetEvent(id)
}

/*
get event by id
*/
func (dp *DictionaryProvider) GetDictionaryItem(dicCode string, ver string, id string) (EventTypes.Event, error) {
	return ep.dbp.DbGetEvent(id)
}
