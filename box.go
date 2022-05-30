package golang_united_school_homework

import (
	"errors"
	"fmt"
	"strings"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape // 4
	shapesCapacity int     // 5 Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

var err1 = errors.New("the index is out of range")

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range. b.shapesCapacity > 0 &&
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	}
	return errors.New("out of the shapesCapacity range")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	// panic("implement me")
	if len(b.shapes)-1 >= i {
		for j := range b.shapes {
			if j == i {
				return b.shapes[j], nil
			}
		}
	}
	return nil, err1
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	// panic("implement me")
	var sh Shape
	if i <= len(b.shapes)-1 {
		sh = b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return sh, nil
	}
	return nil, err1
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	// panic("implement me")
	var sh Shape
	if i <= len(b.shapes)-1 {
		sh = b.shapes[i]
		b.shapes[i] = shape
		return sh, nil
	}
	return nil, err1
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var perimeter float64 = 0
	for _, v := range b.shapes {
		perimeter = perimeter + v.CalcPerimeter()
	}
	return perimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64
	for _, v := range b.shapes {
		area = area + v.CalcArea()
	}
	// panic("implement me")
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	// panic("implement me")
	var sh []Shape
	for i := range b.shapes {
		t := fmt.Sprintf("%T", b.shapes[i])
		if !strings.HasSuffix(t, "Circle") {
			sh = append(sh, b.shapes[i])
		}
	}
	if len(sh) == len(b.shapes) {
		return errors.New("circles are not exist in the list")
	}
	b.shapes = sh
	return nil
}
