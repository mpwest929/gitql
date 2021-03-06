package sql

import "errors"

type Nameable interface {
	Name() string
}

type Resolvable interface {
	Resolved() bool
}

type Transformable interface {
	TransformUp(func(Node) Node) Node
	TransformExpressionsUp(func(Expression) Expression) Node
}

type Expression interface {
	Resolvable
	Type() Type
	Name() string
	Eval(Row) interface{}
	TransformUp(func(Expression) Expression) Expression
}

type Node interface {
	Resolvable
	Transformable
	Schema() Schema
	Children() []Node
	RowIter() (RowIter, error)
}

type Table interface {
	Nameable
	Node
}

type Database interface {
	Nameable
	Tables() map[string]Table
}

var ErrInvalidType = errors.New("invalid type")
