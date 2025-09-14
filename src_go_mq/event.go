package main

type MessageQueueEvent struct {
	MessagesQueue []MessageQueue `json:"messages"`
}

type MessageQueue struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       Details       `json:"details"`
}

type EventMetadata struct {
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	CreatedAt string `json:"created_at"`
}

type Details struct {
	QueueID string  `json:"queue_id"`
	Message Message `json:"message"`
}

type Message struct {
	MessageID              string            `json:"message_id"`
	MD5OfBody              string            `json:"md5_of_body"`
	Body                   string            `json:"body"`
	Attributes             Attributes        `json:"attributes"`
	MessageAttributes      MessageAttributes `json:"message_attributes"`
	MD5OfMessageAttributes string            `json:"md5_of_message_attributes"`
}

type Attributes struct {
	SentTimestamp string `json:"SentTimestamp"`
}

type MessageAttributes struct {
	MessageAttributeKey MessageAttribute `json:"messageAttributeKey"`
}

type MessageAttribute struct {
	DataType    string `json:"dataType"`
	StringValue string `json:"stringValue"`
}
