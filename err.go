package jay

import "errors"

// ErrUnexpectedEOB indicates the end of the byte slice buffer was unexpectedly encountered
// while deserializing a fixed-size block, resulting in an incomplete result.
// Either the overall byte slice length is incorrect,
// or the length specified for one or more fields is incorrect.
var ErrUnexpectedEOB = errors.New("unexpected EOB")
