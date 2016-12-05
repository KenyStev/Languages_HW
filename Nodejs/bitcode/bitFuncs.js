// Sets the bit at pos in the byteeger n.
exports.setBit = function(n, pos) {
    n |= (1 << pos);
    return n;
}

// Clears the bit at pos in n.
exports.clearBit = function(n, pos) {
    let mask = ~(1 << pos);
    n &= mask;
    return n;
}

exports.hasBit = function(n, pos) {
    let val = n & (1 << pos);
    return (val > 0);
}