# BRGo: Compilador Go com Palavras-Chave em Português

BRGo é um wrapper para Go que permite escrever código usando palavras-chave em português do Brasil, facilitando o aprendizado da linguagem para programadores brasileiros.

## Funcionalidades

- Traduz automaticamente palavras-chave em português para Go padrão
- Suporta todas as funcionalidades do Go
- Mantém a mesma performance e características do Go original
- Simples de usar, com comandos semelhantes ao Go

## Exemplos de Tradução

 Português (BRGo) | Go Original |
|-----------------|-------------|
| quebra          | break   	|
| caso            | case   		|
| canal           | chan   		|
| const           | const   	|
| continua        | continue    |
| padrao          | default   	|
| adia            | defer   	|
| senao           | else   		|
| atravessa       | fallthrough |
| para            | for   		|
| func            | func   		|
| vai             | go   		|
| vaipara         | goto   		|
| se              | if   		|
| importa         | import   	|
| interface       | interface   |
| mapa            | map   		|
| pacote          | package   	|
| intervalo       | range   	|
| retorna         | return   	|
| seleciona       | select   	|
| estrutura       | struct   	|
| escolhe         | switch   	|
| tipo            | type   		|
| var             | var   		|
| principal       | main   		|
| imprime         | print   	|
| imprimeln       | println   	|
| novo            | new   		|
| cria            | make   		|
| comprimento     | len   		|
| capacidade      | cap   		|
| anexa           | append   	|
| copia           | copy   		|
| deleta          | delete   	|
| panico          | panic   	|
| recupera        | recover   	|
| verdadeiro      | true   		|
| falso           | false   	|
| nulo            | nil   		|


## Requisitos

- Go 1.16 ou superior instalado no sistema
- Variáveis de ambiente Go configuradas corretamente

## Instalação

```bash
# Clone este repositório
git clone https://github.com/seunome/brgo.git

# Entre no diretório
cd brgo/traducao

# Compile o projeto
go build -o brgo .

# Mova para um diretório no seu PATH (opcional)
# Por exemplo:
# mv brgo /usr/local/bin/
```

## Uso

### Compilar um programa:

```bash
brgo -build arquivo.brgo
```

### Executar um programa:

```bash
brgo -run arquivo.brgo
```

### Compilar e executar (padrão):

```bash
brgo arquivo.brgo
```

### Especificar o arquivo de saída:

```bash
brgo -build -o meuprograma arquivo.brgo
```

## Estrutura do Projeto

- `main.go` - Ponto de entrada do compilador BRGo
- `mapeamento.go` - Dicionário de mapeamento de palavras-chave
- `preprocessador.go` - Implementação do pré-processador
- `exemplos/` - Exemplos de programas em BRGo

## Como Funciona

O BRGo funciona como um pré-processador que traduz código escrito com palavras-chave em português para código Go padrão:

1. Lê o arquivo `.brgo` com código em português
2. Substitui as palavras-chave em português por suas equivalentes em Go
3. Gera um arquivo Go temporário
4. Chama o compilador Go padrão para compilar o código
5. Opcionalmente executa o programa compilado

# Iniciando um Projeto BRGo do Zero

Este tutorial explica como criar um projeto BRGo do zero, configurar um repositório no GitHub, inicializar um módulo Go, fazer commit do código e compilar usando o transpiler BRGo.

## Pré-requisitos

- **Go**: Versão 1.18 ou superior instalada. Veja as [instruções de instalação](https://golang.org/doc/install).
- **Git**: Instalado e configurado. Veja as [instruções de instalação](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
- **GitHub CLI (opcional)**: Facilita a criação do repositório. Veja as [instruções de instalação](https://cli.github.com/manual/installation).
- **BRGo Transpiler**: O código fonte do transpiler deve ser compilado como um executável chamado `brgo`.

## Passo a Passo

### 1. Crie o Repositório no GitHub

1. Acesse [github.com](https://github.com) e faça login.
2. Clique em **New Repository** ou use o GitHub CLI para criar o repositório:
   ```bash
   gh repo create meu-projeto-brgo --public --description "Projeto em BRGo"
3. Clone o repositório localmente:
   ```bash
   git clone https://github.com/usuario/meu-projeto-brgo.git
4. Crie um diretório para seu projeto:
   ```bash
   mkdir meu-projeto-brgo
   cd meu-projeto-brgo
5. Inicialize um módulo Go:
   ```bash
   go mod init meu-projeto-brgo
6. Crie um arquivo README.md:
   ```bash
   echo "# meu-projeto-brgo" > README.md
7. Faça o primeiro commit:
   ```bash
   git add .
   git commit -m "Initial commit"
8. Adicione o repositório remoto:
   ```bash
   git remote add origin https://github.com/usuario/meu-projeto-brgo.git
9. Agora vamos compilar o projeto fazer um código e compilar ele:
    ```bash
    mkdir src
    cd src
10. Criaremos o arquivo principal.brgo:
    ```bash
    echo "pacote principal" > principal.brgo
11. Agora vamos editar o pacote com o editor preferido:
    ```bash
    pacote principal
    importa "github.com/usuario/repo"

    func principal() {
        imprime("Olá, mundo!")
    }

Nota: Substitua github.com/usuario/repo por uma dependência real ou remova a linha importa se não houver dependências externas.

12. Inicializamos o nosso módulo:
    ```bash
    go mod init github.com/seu-usuario/meu-projeto-brgo

13. Criamos o arquivo go.mod:
    ```bash
    module github.com/seu-usuario/meu-projeto-brgo

    go 1.18

14. Novamente adicione e de push no codigo:
    ```bash
    git add .
    git commit -m "Adicionando arquivo principal.brgo"
    git push origin main

15. Agora vamos compilar o código:
    ```bash
    brgo -build -o teste.exe principal.brgo

## Tranpilando

1. Use o comando para transpilar:
    ```bash
    ./brgo src saida

src: Diretório com os arquivos .brgo.
saida: Diretório onde os arquivos .go e o go.mod serão gerados.

Resultado:Arquivo gerado: saida/principal/main.go

2. Agora compilamos:
    ```bash
    cd saida
    go mod tidy
    go build

Resultado:Arquivo gerado: saida/principal/main.exe

## Contribuindo

Contribuições são bem-vindas! Você pode ajudar a:

1. Adicionar mais palavras-chave ao mapeamento
2. Melhorar o pré-processador
3. Criar exemplos e documentação

## Licença

Este projeto está licenciado sob a mesma licença do Go.
