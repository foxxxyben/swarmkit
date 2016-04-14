package state

import (
	"errors"

	"github.com/docker/go-events"
)

// Apply takes an item from the event stream of one Store and applies it to
// a second Store.
func Apply(store Store, item events.Event) (err error) {
	return store.Update(func(tx Tx) error {
		switch v := item.(type) {
		case EventCreateTask:
			return tx.Tasks().Create(v.Task)
		case EventUpdateTask:
			return tx.Tasks().Update(v.Task)
		case EventDeleteTask:
			return tx.Tasks().Delete(v.Task.ID)

		case EventCreateService:
			return tx.Services().Create(v.Service)
		case EventUpdateService:
			return tx.Services().Update(v.Service)
		case EventDeleteService:
			return tx.Services().Delete(v.Service.ID)

		case EventCreateNetwork:
			return tx.Networks().Create(v.Network)
		case EventUpdateNetwork:
			return tx.Networks().Update(v.Network)
		case EventDeleteNetwork:
			return tx.Networks().Delete(v.Network.ID)

		case EventCreateNode:
			return tx.Nodes().Create(v.Node)
		case EventUpdateNode:
			return tx.Nodes().Update(v.Node)
		case EventDeleteNode:
			return tx.Nodes().Delete(v.Node.ID)

		case EventCreateVolume:
			return tx.Volumes().Create(v.Volume)
		case EventUpdateVolume:
			return tx.Volumes().Update(v.Volume)
		case EventDeleteVolume:
			return tx.Volumes().Delete(v.Volume.ID)
		case EventCommit:
			return nil
		}
		return errors.New("unrecognized event type")
	})
}
