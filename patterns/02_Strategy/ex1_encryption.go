package encryption

import "error"

type AsymEncryptionStrategy interface {
	Encrypt(data interface{}) (byte[] cipher, error)
}

type EllipticCurvestrategy struct {} 
type RSA struct {}

func (strat *EllipticCurvestrategy) Encrypt(data interface{}) (byte[] cipher, error) {
	// some complex math
	...
	return cipher, err 
}

func (strat *RSAstrategy) Encrypt(data interface{}) (byte[] cipher, error) {
	// some complex math
	...
	return cipher, err 
} 

func encryptMessage(msg string, strat AsymEncryptionStrategy) (byte[] cipher, error) {
	return strat.Encrypt(msg)
}

// usage
msg := "this is a confidential message"
cipher, err := encryptMessage(msg, encryption.EllipticCurvestrategy)
cipher, err := encryptMessage(msg, encryption.RSAstrategy)