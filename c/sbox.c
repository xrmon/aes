#include <stdint.h>
#include <stdio.h>

// Adapted from code from Wikipedia: https://en.wikipedia.org/wiki/Rijndael_S-box
// Prints out sbox in Go syntax

#define ROTL8(x,shift) ((uint8_t) ((x) << (shift)) | ((x) >> (8 - (shift))))

void initialize_aes_sbox(uint8_t sbox[256]) {
	uint8_t p = 1, q = 1;

	/* loop invariant: p * q == 1 in the Galois field */
	do {
		/* multiply p by 3 */
		p = p ^ (p << 1) ^ (p & 0x80 ? 0x1B : 0);

		/* divide q by 3 (equals multiplication by 0xf6) */
		q ^= q << 1;
		q ^= q << 2;
		q ^= q << 4;
		q ^= q & 0x80 ? 0x09 : 0;

		/* compute the affine transformation */
		uint8_t xformed = q ^ ROTL8(q, 1) ^ ROTL8(q, 2) ^ ROTL8(q, 3) ^ ROTL8(q, 4);

		sbox[p] = xformed ^ 0x63;
	} while (p != 1);

	/* 0 is a special case since it has no inverse */
	sbox[0] = 0x63;
}

int main() {
    uint8_t sbox[256];
    initialize_aes_sbox(sbox);

    printf("const sbox := [256]byte{");
    for (int i = 0; i < 255; i++) {
        printf("0x%x, ", sbox[i]);
    }
    printf("0x%x}\n", sbox[255]);
}
