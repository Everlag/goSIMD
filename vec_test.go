package simd

import (
	"testing"

	"fmt"
)

func TestMul(t *testing.T) {
	v1 := Vec4{1, 2, 3, 4}
	v2 := Vec4{5, 6, 7, 8}

	res := v1.Mul(v2)
	res2 := mul(v1, v2)

	fmt.Println(res, res2)

	// Placeholder until I get it to compile
	if res != res2 {
		t.Fatalf("Expected %v; got %v", res, res2)
	}
}

func TestAdd(t *testing.T) {
	v1 := Vec4{1, 2, 3, 4}
	v2 := Vec4{5, 6, 7, 8}

	res := v1.Add(v2)
	res2 := add(v1, v2)

	fmt.Println(res, res2)

	// Placeholder until I get it to compile
	if res != res2 {
		t.Fatalf("Expected %v; got %v", res, res2)
	}
}
