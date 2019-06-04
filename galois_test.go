package aes

import "testing"

func TestGaloisMultiply(t *testing.T) {
    result := GaloisMultiply(0xB1, 0x29)
    if result != 0x9F {
        t.Errorf("GF(2^8): 0xB1 * 0x29 = 0x%x, expected 0x9F", result)
    }
}
