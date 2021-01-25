module github.com/torvim/matechef

go 1.16

replace github.com/notnil/chess v1.5.0 => ./chess

replace github.com/torvim/matechef/engine v1.0.0 => ./engine

require (
	github.com/abiosoft/ishell v2.0.0+incompatible
	github.com/torvim/matechef/engine v1.0.0
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/notnil/chess v1.5.0
	github.com/stretchr/testify v1.7.0 // indirect
)
