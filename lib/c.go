package c

import d "github.com/nobishino/gomodules-d"

type C struct {
	d.D
	d.D1_1
}

// gomodules-dのv1.2で追加された新しい関数を呼び出す
func (C) Method2() {
	d.D1_2()
}
