package datastructures

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Add(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			if err := q.Add(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
		})
	}
}

func TestStack_GetAndRemove(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			got, err := q.GetAndRemove()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAndRemove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAndRemove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			got, err := q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type fields struct {
		head *node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Stack{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}