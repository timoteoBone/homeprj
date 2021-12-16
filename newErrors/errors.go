package newErrors

import "fmt"

type InternalError struct {
	err string
}

type ThirdPartyError struct {
	err string
}

type OtherError struct {
	err string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%v", "InternalError")
}

func (e *ThirdPartyError) Error() string {
	return fmt.Sprintf("%v", "ThirdPartyError")
}

func (e *OtherError) Error() string {
	return fmt.Sprintf("%v", "OtherError")
}

func KnowType(err error) string {
	switch err.(type) {
	case *OtherError:
		return "Type of error is Other Type"
	case *ThirdPartyError:
		return "Type of error is Third Party"
	case *InternalError:
		return "Type of error is Internal"

	default:
		return "Unknown type"
	}
}
