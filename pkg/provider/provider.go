package provider

// Provider encapsulates logic for interactions between communication service provider and application.
type Provider interface {
	SendSMS(Message) (Confirmation, error)
	ReceiveSMS(count int) ([]Message, error)
}

type Message struct {
}

type Confirmation struct {
}

type Status int

const (
	Unknown Status = iota
	Sent
	Delivered
	Received
	Unread
	Read
)
