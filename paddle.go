package paddle
import "errors"
func Pad(block int, data []byte) ([]byte, error) {
	if block < 1 {
		return nil, errors.New("paddle.Pad: The block size is too low")
	}
	if data == nil {
		return nil, errors.New("paddle.Pad: The data is nil")
	}

	// Append terminator of padding.
	data = append(data, 1)

	// Append rest of padding as zero-ed bytes.
	if block > len(data) {
		data = append(data, make([]byte, block - len(data))...)
	} else if len(data) % block != 0 {
		data = append(data, make([]byte, block - (len(data) % block))...)
	}

	return data, nil
}
func Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("paddle.Unpad: The data has a length of zero")
	}
	for {
		if data[len(data) - 1] == 1 {
			data = data[:len(data) - 1]
			break
		} else if data[len(data) - 1] == 0 {
			data = data[:len(data) - 1]
		} else {
			return nil, errors.New("paddle.Unpad: The data is corrupted")
		}
	}
	return data, nil
}
