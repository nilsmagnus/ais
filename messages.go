package ais

type Packet interface {
	GetMessageId() int64
}

type ShipStaticAndVoyageRelatedData struct {
	MessageId             int64         `json:"messageId"`
	RepeatIndicator       int64         `json:"repeatIndicator"`
	UserID                int64         `json:"userId"`
	AisVersion            int64         `json:"aisVersion"`
	ImoNumber             int64         `json:"imoNumber"`
	CallSign              string        `json:"callSign"`
	Name                  string        `json:"name"`
	ShipType              int64         `json:"shipType"`
	Dimension             ShipDimension `json:"dimension"`
	FixingDeviceType      int64         `json:"fixingDeviceType"`
	ETA                   string        `json:"eta_MMDDHHMM_UTC"`
	MaximumPresentDraught int64         `json:"maximumPresentDraught"`
	Destination           string        `json:"destination"`
	DTE                   bool          `json:"dte"`
	Spare                 bool          `json:"spare"`
	NumberOfBits          int64         `json:"numberOfBits"`
}

type ShipDimension struct {
	A int64
	B int64
	C int64
	D int64
}

func (s ShipStaticAndVoyageRelatedData) GetMessageId() int64 {
	return s.MessageId
}
