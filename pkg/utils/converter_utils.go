package utils

func ToByteArray(src []uint8) []byte {
	var rawData []byte
	for _, bit := range src {
		rawData = append(rawData, bit)
	}
	return rawData
}
