package notifier

import "github.com/nerodesu017/lambdalabs-sniper/src/constants"

// use this for webhooks with discord (or any other way to notify yourself)
type Notifier interface {
	Notify(gpus []*constants.GPU) error
}