package matrix

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	RotateImage(m)
	want := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("got %v", m)
	}
}

func TestSpiralOrder(t *testing.T) {
	got := SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestSetZeroes(t *testing.T) {
	m := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	SetZeroes(m)
	want := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("got %v", m)
	}
}

func TestTranspose(t *testing.T) {
	got := Transpose([][]int{{1, 2, 3}, {4, 5, 6}})
	want := [][]int{{1, 4}, {2, 5}, {3, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestSearchMatrix(t *testing.T) {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	if !SearchMatrix(m, 3) {
		t.Error("3 present")
	}
	if SearchMatrix(m, 13) {
		t.Error("13 absent")
	}
}

func TestSearchMatrixII(t *testing.T) {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if !SearchMatrixII(m, 5) {
		t.Error("5 present")
	}
	if SearchMatrixII(m, 20) {
		t.Error("20 absent")
	}
}

func TestFloodFill(t *testing.T) {
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	got := FloodFill(img, 1, 1, 2)
	want := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestGameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	GameOfLife(board)
	want := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
	if !reflect.DeepEqual(board, want) {
		t.Errorf("got %v", board)
	}
}

func TestDiagonalTraverse(t *testing.T) {
	got := DiagonalTraverse([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	want := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}
