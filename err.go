package jay

import "errors"

// ErrUnexpectedEOB indicates the end of the byte slice buffer was unexpectedly encountered
// while deserialising a fixed-size block, resulting in an incomplete result.
var ErrUnexpectedEOB = errors.New("unexpected EOB")
