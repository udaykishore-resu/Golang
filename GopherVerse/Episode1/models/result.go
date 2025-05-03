package models

type Result struct {
	ProducerID int
	URL        string
	Bytes      int
	Err        error
}
