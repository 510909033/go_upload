package bo

import "fmt"

func Test1() {
	m := &menuBO{Id: 12}
	a, err := NewmenuBO(m, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a.UpdateTs)
	}
}
