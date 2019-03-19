package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	//fmt.Println(runtime.GOARCH)
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitura amd64")
	}
	t.Fail()
}
