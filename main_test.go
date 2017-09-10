package main

import (
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_serve(t *testing.T) {
	type args struct {
		r chan int64
		w *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serve(tt.args.r, tt.args.w)
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		j int64
		g chan int
		k *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process(tt.args.j, tt.args.g, tt.args.k)
		})
	}
}
