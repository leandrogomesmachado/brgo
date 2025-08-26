package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	buildCmd   = flag.Bool("build", false, "Compilar programa")
	runCmd     = flag.Bool("run", false, "Executar programa")
	outputFlag = flag.String("o", "", "Caminho do arquivo de saída")
	tempDir    = flag.String("temp", os.TempDir(), "Diretório temporário para arquivos intermediários")
)

func main() {
	flag.Parse()

	// Verificar se há argumentos suficientes
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Uso: brgo [opcoes] arquivo.brgo")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputPath := args[0]
	
	// Verificar se o arquivo existe
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Printf("Erro: arquivo %s não encontrado\n", inputPath)
		os.Exit(1)
	}

	// Determinar se o comando deve construir ou executar
	shouldBuild := *buildCmd
	shouldRun := *runCmd

	// Se nenhum for especificado, assumir ambos
	if !shouldBuild && !shouldRun {
		shouldBuild = true
		shouldRun = true
	}

	// Criar um novo preprocessador
	preprocessador := NovoPreprocessador()

	// Determinar caminho temporário para o arquivo Go
	tempGoFile := filepath.Join(*tempDir, "brgo_"+filepath.Base(strings.TrimSuffix(inputPath, ".brgo"))+".go")

	// Processar o arquivo .brgo para .go
	if err := preprocessador.ProcessarArquivo(inputPath, tempGoFile); err != nil {
		fmt.Printf("Erro ao pré-processar o arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo traduzido com sucesso: %s\n", tempGoFile)

	// Se apenas pré-processamento for solicitado, sair
	if !shouldBuild && !shouldRun {
		return
	}

	// Determinar caminho de saída
	outputPath := *outputFlag
	if outputPath == "" {
		if shouldRun {
			// Se estiver apenas executando, usar um binário temporário
			outputPath = filepath.Join(*tempDir, "brgo_"+filepath.Base(strings.TrimSuffix(inputPath, ".brgo")))
		} else {
			// Caso contrário, usar o nome do arquivo sem extensão
			outputPath = strings.TrimSuffix(inputPath, ".brgo")
		}
	}

	// Compilar o código Go
	if shouldBuild {
		fmt.Println("Compilando...")
		cmd := exec.Command("go", "build", "-o", outputPath, tempGoFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Erro ao compilar: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Compilado com sucesso: %s\n", outputPath)
	}

	// Executar o programa, se solicitado
	if shouldRun {
		fmt.Println("Executando...")
		cmd := exec.Command(outputPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				os.Exit(exitError.ExitCode())
			}
			fmt.Printf("Erro ao executar: %v\n", err)
			os.Exit(1)
		}

		// Se o binário for temporário e executado com sucesso, limpá-lo
		if *outputFlag == "" && shouldRun {
			os.Remove(outputPath)
		}
	}

	// Limpar o arquivo .go temporário se necessário
	if shouldBuild || shouldRun {
		os.Remove(tempGoFile)
	}
}
