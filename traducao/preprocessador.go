package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Preprocessador converte código BRGo para Go padrão
type Preprocessador struct {
	mapeamento map[string]string
}

// NovoPreprocessador cria uma nova instância do preprocessador
func NovoPreprocessador() *Preprocessador {
	return &Preprocessador{
		mapeamento: MapeamentoPtBrParaGo,
	}
}

// ProcessarArquivo converte um arquivo .brgo para .go
func (p *Preprocessador) ProcessarArquivo(caminhoEntrada, caminhoSaida string) error {
	entrada, err := os.Open(caminhoEntrada)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo de entrada: %w", err)
	}
	defer entrada.Close()

	// Cria diretórios para o arquivo de saída se necessário
	dir := filepath.Dir(caminhoSaida)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de saída: %w", err)
	}

	saida, err := os.Create(caminhoSaida)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %w", err)
	}
	defer saida.Close()

	return p.Processar(entrada, saida)
}

// Processar converte o código de BRGo para Go
func (p *Preprocessador) Processar(entrada io.Reader, saida io.Writer) error {
	scanner := bufio.NewScanner(entrada)
	
	// Expressão regular para identificar palavras (identifiers)
	reIdentificador := regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*`)
	
	// Expressão para identificar strings (para não substituir dentro delas)
	reString := regexp.MustCompile(`"[^"]*"`)
	reCharacter := regexp.MustCompile(`'[^']*'`)
	
	for scanner.Scan() {
		linha := scanner.Text()
		
		// Preservar strings e caracteres
		stringsPreservadas := map[string]string{}
		linha = reString.ReplaceAllStringFunc(linha, func(s string) string {
			placeholder := fmt.Sprintf("__STRING_%d__", len(stringsPreservadas))
			stringsPreservadas[placeholder] = s
			return placeholder
		})
		
		charsPreservados := map[string]string{}
		linha = reCharacter.ReplaceAllStringFunc(linha, func(s string) string {
			placeholder := fmt.Sprintf("__CHAR_%d__", len(charsPreservados))
			charsPreservados[placeholder] = s
			return placeholder
		})
		
		// Traduzir palavras-chave
		linha = reIdentificador.ReplaceAllStringFunc(linha, func(palavra string) string {
			if goWord, ok := p.mapeamento[palavra]; ok {
				return goWord
			}
			return palavra
		})
		
		// Restaurar strings e caracteres
		for placeholder, original := range stringsPreservadas {
			linha = strings.Replace(linha, placeholder, original, 1)
		}
		for placeholder, original := range charsPreservados {
			linha = strings.Replace(linha, placeholder, original, 1)
		}
		
		fmt.Fprintln(saida, linha)
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erro ao ler arquivo de entrada: %w", err)
	}
	
	return nil
}

// ProcessarDiretorio processa todos os arquivos .brgo em um diretório
func (p *Preprocessador) ProcessarDiretorio(dirEntrada, dirSaida string) error {
	return filepath.Walk(dirEntrada, func(caminho string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Pular diretórios
		if info.IsDir() {
			return nil
		}
		
		// Processar apenas arquivos .brgo
		if !strings.HasSuffix(caminho, ".brgo") {
			return nil
		}
		
		// Calcular caminho relativo para manter a estrutura
		relativo, err := filepath.Rel(dirEntrada, caminho)
		if err != nil {
			return fmt.Errorf("erro ao calcular caminho relativo: %w", err)
		}
		
		// Criar caminho de saída com extensão .go
		caminhoSaida := filepath.Join(dirSaida, strings.TrimSuffix(relativo, ".brgo")+".go")
		
		// Processar o arquivo
		return p.ProcessarArquivo(caminho, caminhoSaida)
	})
}
