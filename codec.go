package ais

func extractNumber(payload []byte, isSigned bool, offset uint, width uint) int64 {
	result := uint64(0)

	for i := offset; i < offset+width; i++ {
		result <<= 1
		if i < uint(len(payload)) {
			result |= uint64(payload[i])
		}
	}

	if isSigned {
		return makeSigned(result, width)
	}

	return int64(result)
}

func makeSigned(input uint64, length uint) int64 {
	result := int64(input)
	maxValue := int64(1) << length

	if result >= maxValue/2 {
		result = result - maxValue
	}

	return result
}

func extractString(payload []byte, offset uint, width uint, dropSpace bool) string {
	numChars := width / 6

	result := make([]byte, numChars)

	for i := uint(0); i < numChars; i++ {
		number := extractNumber(payload, false, offset, 6)
		offset += 6
		if number < 32 {
			number = number + 64
		}

		result[i] = byte(number)
	}

	/* The string is closed by @ */
	stripSpace := len(result)
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != '@' {
			if !dropSpace || result[i] != ' ' {
				break
			}
		}
		stripSpace--
	}

	result = result[:stripSpace]

	return string(result)
}
