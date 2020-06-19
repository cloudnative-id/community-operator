package dispatcher

import (
	"fmt"
)

type Dispatcher struct {
    A int
    B int
}

func (db *Dispatcher) DemoMethod(x int) {
    y:= db.A + db.B + x
    fmt.Println(y)
}
func (db *Dispatcher) Telegram() {
	fmt.Println("jhaha")
}
