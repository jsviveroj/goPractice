package variables

import (
	"fmt"
)

func ShowInts() {
	var commonInt int
	int32 := int32(10)
	int64 := int64(100)

	fmt.Println("commonInt = ", commonInt)
	fmt.Println("Int32 = ", int32)
	fmt.Println("Int64 = ", int64)
}
