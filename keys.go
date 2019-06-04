package aes

// Rijndael key schedule implementation

func rotWord(in [4]byte) (out [4]byte) {
    // Circular shift left of one byte
    out[0] = in[1]
    out[1] = in[2]
    out[2] = in[3]
    out[3] = in[0]
    return
}

func subWord(in [4]byte) (out [4]byte) {
    // Substitutes bytes using sbox in aes.go
    for i, v := range in {
        out[i] = sbox[v]
    }
    return
}

func xor(inputs ...[4]byte) (out [4]byte) {
    // Xor all provided arguments
    for _, input := range inputs {
        for i, _ := range out {
            out[i] ^= input[i]
        }
    }
    return
}

func roundConstants(rounds int) [][4]byte {
    // Create the output slice
    var out = make([][4]byte, rounds)
    // Gets the round constants for all rounds
    var rc = make([]byte, rounds)
    for i := 1; i < rounds; i++ {
        if i == 1 {
            rc[i] = 1
        } else if rc[i-1] < 0x80 {
            rc[i] = 2*rc[i-1]
        } else {
            rc[i] = 2*rc[i-1] ^ 0x1B
        }
        // Expand to the full round constant
        out[i][0] = rc[i]
    }
    return out
}

func KeyExpansion(key []byte, keyLen int, rounds int) [][16]byte {
    // Key length in 32-bit words
    n := keyLen
    // Get slice of round constants
    rcon := roundConstants(rounds)
    // Create a slice to store the columns of the round keys
    var W = make([][4]byte, 4*rounds)
    // Expand the original key into the first n columns
    for i := 0; i < n; i++ {
        copy(W[i][:], key[i*4:i*4+4])
    }
    // Calculate the round keys
    for i := n; i < 4*rounds; i++ {
        if i % n == 0 {
            W[i] = xor(W[i-n], subWord(rotWord(W[i-1])), rcon[i/n])
        } else if (n > 6) && (i % n == 4) {
            W[i] = xor(W[i-n], subWord(W[i-1]))
        } else {
            W[i] = xor(W[i-n], W[i-1])
        }
    }
    // Create the output slice
    var result = make([][16]byte, rounds)
    for round := 0; round < rounds; round++ {
        copy(result[round][0:4][:],   W[4*round][:])
        copy(result[round][4:8][:],   W[4*round+1][:])
        copy(result[round][8:12][:],  W[4*round+2][:])
        copy(result[round][12:16][:], W[4*round+3][:])
    }
    return result
}

func ExpandKey128(key [16]byte) (roundKeys [11][16]byte) {
    expansion := KeyExpansion(key[:], 4, 11)
    copy(roundKeys[:], expansion)
    return
}

func ExpandKey192(key [24]byte) (roundKeys [13][16]byte) {
    expansion := KeyExpansion(key[:], 6, 13)
    copy(roundKeys[:], expansion)
    return
}

func ExpandKey256(key [32]byte) (roundKeys [15][16]byte) {
    expansion := KeyExpansion(key[:], 8, 15)
    copy(roundKeys[:], expansion)
    return
}
