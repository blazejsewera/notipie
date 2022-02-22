package model

type netNotification interface {
	ToJSON() string
	validate() bool
}

const NotEnoughInfoInNotificationErrorMessage = "not enough information in net notification"
