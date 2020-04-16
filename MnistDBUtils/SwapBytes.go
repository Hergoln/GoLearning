package MnistDBUtils

func swapBytes(byt []byte) []byte {
	bytR := make([]byte, len(byt))
	for i := range byt {
		bytR[i] = byt[len(byt) - i - 1]
	}
	return bytR
}