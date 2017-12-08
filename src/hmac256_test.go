package main

import "testing"

func BenchmarkSmall(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("Lorem ipsum")

	for n := 0; n < b.N; n++ {
		Digest(hmacInput, hmacTestKey)
	}
}

func BenchmarkLarge(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("00123014764700968325001230147647009683250012301476470096832500123014764700968325")

	for n := 0; n < b.N; n++ {
		Digest(hmacInput, hmacTestKey)
	}
}

func BenchmarkSmallParallel(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("Lorem ipsum")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digest(hmacInput, hmacTestKey)
		}
	})
}

func BenchmarkLargeParallel(b *testing.B) {
	hmacTestKey := []byte("xxxx")
	hmacInput := []byte("00123014764700968325001230147647009683250012301476470096832500123014764700968325")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digest(hmacInput, hmacTestKey)
		}
	})
}
