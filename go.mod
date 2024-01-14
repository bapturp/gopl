module github.com/bapturp/gopl

go 1.21.6

replace (
	github.com/bapturp/gopl/ch02/ex2.01 => ./ch02/ex2.01
	github.com/bapturp/gopl/ch02/ex2.02/lenconv => ./ch02/ex2.02/lenconv
	github.com/bapturp/gopl/ch02/ex2.02/weiconv => ./ch02/ex2.02/weiconv
	github.com/bapturp/gopl/ch04/ex4.10/github => ./ch04/ex4.10/github
	github.com/bapturp/gopl/ch04/ex4.11/client => ./ch04/ex4.10/client
	github.com/bapturp/gopl/ch04/ex4.11/model => ./ch04/ex4.10/model
)

require github.com/tj/go-editor v1.0.0

require github.com/pkg/errors v0.8.1 // indirect
