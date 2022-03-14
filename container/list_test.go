package container

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	L := NewList()
	L.AddHead(1)
	L.AddTail(2)
	L.AddNodeByNo(2, 3)
	L.DeleteAllByVal(1)
	for i := 0; i < L.Size(); i++ {
		fmt.Println(L.GetNodeByNo(i + 1))
	}

}
