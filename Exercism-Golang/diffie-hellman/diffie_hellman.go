package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	res, err := rand.Int(rand.Reader, big.NewInt(0).Sub(p, two))
	if err != nil {
		panic(err)
	}
	return big.NewInt(0).Add(res, two)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
