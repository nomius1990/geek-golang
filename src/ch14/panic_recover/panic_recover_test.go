package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	defer func() {

		//这里相当于 使用recover 达到了 try catch 的效果
		if err := recover(); err != nil {

			fmt.Println("Recover from!", err)
		}
		// fmt.Println("Finall!")

	}()

	fmt.Println("Start")
	// os.Exit(-1)
	panic(errors.New("Something is wrong!"))
}
