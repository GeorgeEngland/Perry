package data

type AlertData struct {
	Id       int
	Priority int
}

type NotificationData struct {
	Id    int
	Alert AlertData
}
