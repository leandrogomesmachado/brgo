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
    // Expressão para identificar a declaração de pacote
    rePacote := regexp.MustCompile(`^pacote\s+([a-zA-Z][a-zA-Z0-9_]*)`)

    for scanner.Scan() {
        linha := scanner.Text()

        // Identificar o nome do pacote
        if matches := rePacote.FindStringSubmatch(linha); len(matches) > 1 {
            linha = strings.Replace(linha, "pacote", "package", 1)
        }
        
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
    pacotes := make(map[string][]string) // Mapa de nome do pacote para lista de arquivos
    importsExternos := make(map[string]struct{}) // Conjunto de importações externas

    // Primeira passada: coletar nomes de pacotes e importações
    err := filepath.Walk(dirEntrada, func(caminho string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if info.IsDir() || !strings.HasSuffix(caminho, ".brgo") {
            return nil
        }

        // Ler o arquivo para identificar o nome do pacote
        arquivo, err := os.Open(caminho)
        if err != nil {
            return fmt.Errorf("erro ao abrir arquivo %s: %w", caminho, err)
        }
        defer arquivo.Close()

        scanner := bufio.NewScanner(arquivo)
        rePacote := regexp.MustCompile(`^pacote\s+([a-zA-Z][a-zA-Z0-9_]*)`)
        reImport := regexp.MustCompile(`importa\s+"([^"]+)"`)

        var pacoteNome string
        for scanner.Scan() {
            linha := scanner.Text()
            if matches := rePacote.FindStringSubmatch(linha); len(matches) > 1 {
                pacoteNome = matches[1]
            }
            if matches := reImport.FindStringSubmatch(linha); len(matches) > 1 {
                importPath := matches[1]
                if strings.HasPrefix(importPath, "github.com/") {
                    importsExternos[importPath] = struct{}{}
                }
            }
        }

        if pacoteNome != "" {
            pacotes[pacoteNome] = append(pacotes[pacoteNome], caminho)
        }

        return scanner.Err()
    })

    if err != nil {
        return err
    }

    // Segunda passada: processar arquivos e organizar por pacote
    for pacote, arquivos := range pacotes {
        for _, caminho := range arquivos {
            relativo, err := filepath.Rel(dirEntrada, caminho)
            if err != nil {
                return fmt.Errorf("erro ao calcular caminho relativo: %w", err)
            }

            // Criar diretório com o nome do pacote
            caminhoSaida := filepath.Join(dirSaida, pacote, strings.TrimSuffix(relativo, ".brgo")+".go")
            if err := p.ProcessarArquivo(caminho, caminhoSaida); err != nil {
                return err
            }
        }
    }

    // Gerar arquivo go.mod
    if err := p.gerarGoMod(dirSaida, importsExternos); err != nil {
        return fmt.Errorf("erro ao gerar go.mod: %w", err)
    }

    return nil
}

// gerarGoMod cria um arquivo go.mod com as dependências
func (p *Preprocessador) gerarGoMod(dirSaida string, importsExternos map[string]struct{}) error {
    goModPath := filepath.Join(dirSaida, "go.mod")
    f, err := os.Create(goModPath)
    if err != nil {
        return err
    }
    defer f.Close()

    // Nome do módulo (pode ser personalizado ou derivado do diretório)
    nomeModulo := filepath.Base(dirSaida)
    fmt.Fprintf(f, "module %s\n\n", nomeModulo)
    fmt.Fprintln(f, "go 1.18") // Use a versão mínima do Go necessária

    if len(importsExternos) > 0 {
        fmt.Fprintln(f, "\nrequire (")
        for importPath := range importsExternos {
            fmt.Fprintf(f, "\t%s v0.0.0\n", importPath)
        }
        fmt.Fprintln(f, ")")
    }

    return nil
}