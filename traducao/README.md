# BRGo: Compilador Go com Palavras-Chave em Português

BRGo é um wrapper para Go que permite escrever código usando palavras-chave em português do Brasil, facilitando o aprendizado da linguagem para programadores brasileiros.

## Funcionalidades

- Traduz automaticamente palavras-chave em português para Go padrão
- Suporta todas as funcionalidades do Go
- Mantém a mesma performance e características do Go original
- Simples de usar, com comandos semelhantes ao Go

## Exemplos de Tradução

| Português (BRGo) | Go Original |
|------------------|------------|
| se               | if         |
| senao            | else       |
| para             | for        |
| escolhe          | switch     |
| caso             | case       |
| estrutura        | struct     |
| interface        | interface  |
| pacote           | package    |
| importa          | import     |
| func             | func       |
| retorna          | return     |
| principal        | main       |

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

## Contribuindo

Contribuições são bem-vindas! Você pode ajudar a:

1. Adicionar mais palavras-chave ao mapeamento
2. Melhorar o pré-processador
3. Criar exemplos e documentação

## Licença

Este projeto está licenciado sob a mesma licença do Go.
