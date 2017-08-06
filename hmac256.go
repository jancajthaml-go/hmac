package main

import (
	"crypto"
	_ "crypto/sha256"
	"hash"
)

/*
type digest struct {
  	h     [8]uint32
  	x     [64]byte
  	nx    int
  	len   uint64
}
*/

func Hmac256(input []byte, secret []byte) (checksum []byte) {
	var (
		opad  []byte
		ipad  []byte
		outer hash.Hash
		inner hash.Hash
		i     int
	)

	outer = crypto.SHA256.New()
	inner = crypto.SHA256.New()

	ipad = make([]byte, 64)
	opad = make([]byte, 64)

	if len(secret) > 64 {
		outer.Write(secret)
		secret = outer.Sum(nil)
		/*
			hash := outer.checkSum()
			secret = append(nil, hash[:]...)
		*/
	}
	copy(ipad, secret)
	copy(opad, secret)

	i = len(ipad) - 1
	for i > 0 {
		ipad[i] ^= 0x36
		i--
	}

	i = len(opad) - 1
	for i > 0 {
		opad[i] ^= 0x5c
		i--
	}

	inner.Write(ipad)
	inner.Write(input)

	in := inner.Sum(nil)
	/*
		hash := inner.checkSum()
		in = append(nil, hash[:]...)
	*/
	outer.Reset()
	/*
				d.h[0] = 0x6A09E667
		  		d.h[1] = 0xBB67AE85
		  		d.h[2] = 0x3C6EF372
		  		d.h[3] = 0xA54FF53A
		  		d.h[4] = 0x510E527F
		  		d.h[5] = 0x9B05688C
		  		d.h[6] = 0x1F83D9AB
		  		d.h[7] = 0x5BE0CD19
			  	d.nx = 0
			  	d.len = 0
	*/
	outer.Write(opad) //outer.Write(in[0:])
	outer.Write(in)

	return outer.Sum(in) //return outer.Sum(in[:0])
	/*

	 */
}
