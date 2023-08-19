package ais

import "fmt"

func Decode(payload []byte) (Packet, error) {
	if len(payload)%8 != 0 {
		return nil, fmt.Errorf("Payload length must be a multiple of 8, was %d", len(payload))
	}

	messageId := extractNumber(payload, false, 0, 6)

	switch messageId {

	case 5:
		return parseShipAndVoyageRelatedData(messageId, payload)
	default:
		return nil, fmt.Errorf("Unknown message id %d", messageId)

	}
}

func parseShipAndVoyageRelatedData(messageId int64, paylaod []byte) (Packet, error) {
	return ShipStaticAndVoyageRelatedData{
		MessageId:       messageId,
		RepeatIndicator: extractNumber(paylaod, false, 6, 2),
		UserID:          extractNumber(paylaod, false, 8, 30),
		AisVersion:      extractNumber(paylaod, false, 38, 2),
		ImoNumber:       extractNumber(paylaod, false, 40, 30),
		CallSign:        extractString(paylaod, 70, 42, true),
		Name:            extractString(paylaod, 112, 120, true),
		ShipType:        extractNumber(paylaod, false, 232, 8),
		Dimension: ShipDimension{
			A: extractNumber(paylaod, false, 240, 9),
			B: extractNumber(paylaod, false, 249, 9),
			C: extractNumber(paylaod, false, 258, 6),
			D: extractNumber(paylaod, false, 264, 6),
		},
		FixingDeviceType:      extractNumber(paylaod, false, 270, 4),
		ETA:                   extractString(paylaod, 274, 20, true),
		MaximumPresentDraught: extractNumber(paylaod, false, 294, 8),
		Destination:           extractString(paylaod, 302, 120, true),
	}, nil
}
