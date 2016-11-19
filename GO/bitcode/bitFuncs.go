package bitcode

// Sets the bit at pos in the byteeger n.
func setBit(n byte, pos uint) byte {
    n |= (1 << pos)
    return n
}

// Clears the bit at pos in n.
func clearBit(n byte, pos uint) byte {
    mask := ^(1 << pos)
    n &= byte(mask)
    return n
}

func hasBit(n byte, pos uint) bool {
    val := n & (1 << pos)
    return (val > 0)
}