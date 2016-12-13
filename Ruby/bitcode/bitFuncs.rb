# Sets the bit at pos in the byteeger n.
def setBit(n, pos)
	mask = (1 << pos)
	return (n | mask)
end
# Clears the bit at pos in n.
def clearBit(n, pos)
	mask = ~(1 << pos)
	return (n & mask)
end
def hasBit(n, pos)
	val = n & (1 << pos)
	return (val > 0)
end