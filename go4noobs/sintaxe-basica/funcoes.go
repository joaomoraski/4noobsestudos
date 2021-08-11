// uma função simples
func functionName() {}

// função com parâmetros (novamente, os tipos vêm depois dos identificadores)
func functionName(param1 string, param2 int) {}

// vários parâmetros do mesmo tipo
func functionName(param1, param2 int) {}

// retorno do tipo de declaração
func functionName() int {
	return 42
}

// Pode retornar vários valores de uma vez
func returnMulti() (int, string) {
	return 42, "foobar"
}

var x, str = returnMulti()

// Retorne vários resultados nomeados simplesmente usando return
func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"

	return // n e s serão retornados
}

var x, str = returnMulti2()

