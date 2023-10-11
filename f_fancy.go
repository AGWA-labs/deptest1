//go:build fancy
package deptest1

import "github.com/AGWA-labs/deptest2"

func F() {
	println("deptest1.F (fancy) called")
	deptest2.F()
}
