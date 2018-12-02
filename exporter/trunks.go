package exporter

import "time"

type Trunk struct {
	ID                 string     `json:"Id"`
	Number             string     `json:"Number"`
	Name               string     `json:"Name"`
	Host               string     `json:"Host"`
	Type               string     `json:"Type"`
	SimCalls           string     `json:"SimCalls"`
	ExternalNumber     string     `json:"ExternalNumber"`
	IsRegistered       bool       `json:"IsRegistered"`
	RegisterOkTime     *time.Time `json:"RegisterOkTime"`
	RegisterSentTime   *time.Time `json:"RegisterSentTime"`
	RegisterFailedTime *time.Time `json:"RegisterFailedTime"`
	CanBeDeleted       bool       `json:"CanBeDeleted"`
}