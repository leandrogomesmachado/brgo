// Mapeamento de palavras-chave do Go para Português do Brasil
package main

// MapeamentoPtBrParaGo associa palavras-chave em português para suas equivalentes em Go
var MapeamentoPtBrParaGo = map[string]string{
	// Palavras-chave
	"quebra":       "break",
	"caso":         "case",
	"canal":        "chan",
	"const":        "const",
	"continua":     "continue",
	"padrao":       "default",
	"adia":         "defer",
	"senao":        "else",
	"atravessa":    "fallthrough",
	"para":         "for",
	"func":         "func",
	"vai":          "go",
	"vaipara":      "goto",
	"se":           "if",
	"importa":      "import",
	"interface":    "interface",
	"mapa":         "map",
	"pacote":       "package",
	"intervalo":    "range",
	"retorna":      "return",
	"seleciona":    "select",
	"estrutura":    "struct",
	"escolhe":      "switch",
	"tipo":         "type",
	"var":          "var",

	// Funções builtin importantes
	"principal":    "main",
	"imprime":      "print",
	"imprimeln":    "println",
	"novo":         "new",
	"cria":         "make",
	"comprimento": "len",
	"capacidade":   "cap",
	"anexa":        "append",
	"copia":        "copy",
	"deleta":       "delete",
	"panico":       "panic",
	"recupera":     "recover",
	
	// Valores especiais
	"verdadeiro":   "true",
	"falso":        "false",
	"nulo":         "nil",
}

// GoParaPtBr inverte o mapa para traduzir de Go para Português
func gerarGoParaPtBr() map[string]string {
	inverso := make(map[string]string, len(MapeamentoPtBrParaGo))
	for pt, go_kw := range MapeamentoPtBrParaGo {
		inverso[go_kw] = pt
	}
	return inverso
}
