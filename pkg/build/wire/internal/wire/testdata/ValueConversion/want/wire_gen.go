// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectedMessage() Foo {
	foo := _wireFooValue
	return foo
}

var (
	_wireFooValue = Foo("Hello, World!")
)