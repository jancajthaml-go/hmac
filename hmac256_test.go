package main

import "testing"

func BenchmarkHmac256Parallel(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("Lorem ipsum")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Hmac256(hmacInput, hmacTestKey)
		}
	})
}

func BenchmarkHmac25665Serial(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("Lorem ipsum")

	for n := 0; n < b.N; n++ {
		Hmac256(hmacInput, hmacTestKey)
	}
}
