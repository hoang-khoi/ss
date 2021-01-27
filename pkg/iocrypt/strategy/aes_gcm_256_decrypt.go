package strategy

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// AesGcm256Decrypt parses and decrypts a block of data by AES-GCM-256.
type AesGcm256Decrypt struct {
	Key [32]byte
}

// Process uses chunk.Deserialize to parse and decrypt blocks of data.
func (s *AesGcm256Decrypt) Process(dest io.Writer, src io.Reader) (n int, err error) {
	chunk := EncryptedChunk{}

	n, err = chunk.Deserialize(src)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(s.Key[:])
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	decrypted, err := gcm.Open(nil, chunk.Nonce, chunk.Block, nil)
	if err != nil {
		return
	}

	_, err = dest.Write(decrypted)
	return
}
