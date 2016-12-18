package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(T *testing.T) {
	fmt.Println("Graph testing start")
	g := New(3)
	assert.Equal(T, 3, g.Nodes())
	assert.Equal(T, 0, g.Edges())
	g.print()
	var edge Edge
	edge.SetSrc(0)
	edge.SetDest(1)
	fmt.Println("After adding (0,1)")
	err := g.AddEdge(edge)
	g.print()
	assert.Nil(T, err)
	assert.Equal(T, 1, g.Edges())
	assert.Equal(T, true, g.Present(edge))
	err = g.AddEdge(edge)
	assert.NotNil(T, err)
	edge.SetDest(2)
	assert.Equal(T, false, g.Present(edge))
	edge.SetDest(-1)
	err = g.AddEdge(edge)
	assert.NotNil(T, err)
	edge.SetSrc(1)
	edge.SetDest(-1)
	err = g.DeleteEdge(edge)
	assert.NotNil(T, err)
	edge.SetDest(2)
	err = g.DeleteEdge(edge)
	assert.NotNil(T, err)
	fmt.Println("After deleting (0,1)")
	edge.SetSrc(0)
	edge.SetDest(1)
	err = g.DeleteEdge(edge)
	assert.Nil(T, err)
	assert.Equal(T, 0, g.Edges())
	g.print()
	fmt.Println("After againg adding (0,1)")
	edge.SetSrc(0)
	edge.SetDest(1)
	err = g.AddEdge(edge)
	g.print()
	assert.Nil(T, err)
	fmt.Println("After adding (2,1)")
	edge.SetSrc(2)
	edge.SetDest(1)
	err = g.AddEdge(edge)
	g.print()
	assert.Nil(T, err)
	fmt.Println("After adding (2,2)")
	edge.SetSrc(2)
	edge.SetDest(2)
	err = g.AddEdge(edge)
	g.print()
	assert.Nil(T, err)
	fmt.Println("Graph Testing end")
}
