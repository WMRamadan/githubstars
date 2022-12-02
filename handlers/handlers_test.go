package handlers

import (
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	type args struct {
		location string
		bearer   string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsers(tt.args.location, tt.args.bearer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepos(t *testing.T) {
	type args struct {
		language string
		bearer   string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRepos(tt.args.language, tt.args.bearer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}