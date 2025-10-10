package vpc

import "errors"

const (
	UUID = "uuid"
)

const (
	IPRangeKeys = "ip_range"
	NameKeys    = "name"

	DescriptionKeys = "description"
)

var errIncorrectOctet = errors.New("incorrect octet for IP range")
