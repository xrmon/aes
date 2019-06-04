package aes

// m(x) = x^8 + x^4 + x^3 + x + 1 = 0x11B
const m uint = 0x11B

func getOrder(p uint) uint {
    // Check for edge case where p = 0
    if p == 0 {
        return 0
    }
    // Returns the order of polynomial p
    order := uint(0)
    // Keep increasing order until we reach order of p
    for (p >> order) != 0 {
        order++
    }
    return order - 1
}

func polyMultiply(a uint, b uint) uint {
    // Polynomial multiplication
    var c uint = 0
    var a_bit, b_bit uint
    a_limit, b_limit := getOrder(a)+1, getOrder(b)+1
    // Multiply polynomials a with b to produce c
    for a_pow := uint(0); a_pow < a_limit; a_pow++ {
        for b_pow := uint(0); b_pow < b_limit; b_pow++ {
            // Extract current bits for a and b
            a_bit = (a >> a_pow) & 1
            b_bit = (b >> b_pow) & 1
            // If non-zero term, multiply the polynomial terms
            if a_bit == 1 && b_bit == 1 {
                var order uint = a_pow + b_pow
                var bit uint = 1 << order
                // Add to c
                c = c ^ bit
            }
        }
    }
    return c
}

func polyDivide(p uint, m uint) (uint, uint) {
    // Divides polynomial p by polynomial m
    // Returns both remainder and quotient
    var order uint = 15
    var quotient uint = 0
    // Keep reducing until order of p is smaller than order of m
    for {
        order = getOrder(p)

        // If order is less than 8, we have finished
        if order < 8 {
            return p, quotient
        }
        // Find what we need to multiply m by to get p
        var q uint = order - 8
        // Find the remainder for this round
        p = p ^ (m << q)
        // Update the quotient
        quotient ^= 1 << q
    }
}

func galoisReduce(p uint) (uint, uint) {
    // Reduces p modulo the irreducible polynomial m(x)
    return polyDivide(p, m)
}

func GaloisMultiply(a byte, b byte) byte {
    // Multiply a by b to get c
    c := polyMultiply(uint(a), uint(b))

    // Reduce modulo m(x) and return
    result, _ := galoisReduce(c)
    return byte(result)
}
