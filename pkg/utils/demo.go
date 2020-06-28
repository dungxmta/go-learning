package utils

// type UtilsIn interface {
// 	Demo() string
// }

type UtilsObj struct {
	Demo func() string
}

var Utils *UtilsObj

func GetUtils() *UtilsObj {
	if Utils == nil {
		Utils = &UtilsObj{
			Demo: Demo,
		}
	}
	return Utils
}

func Demo() string {
	return "1"
}
