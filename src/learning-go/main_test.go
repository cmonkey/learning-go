package main

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOne(t *testing.T){

	x := 1
	addOne(&x)
	if x != 2{
		t.Fatal("Does not add one ")
	}
	assert.Equal(t, x, 2, "Does not add one")
}
