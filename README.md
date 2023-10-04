# Hangman-Game

## Visão Geral

O "hangman-game" é um jogo da forca simples escrito em Go, que desafia os jogadores a adivinhar uma palavra oculta. O jogo utiliza um arquivo JSON chamado "words.json" para escolher aleatoriamente palavras que os jogadores devem adivinhar.

## Dependências

O "hangman-game" não possui dependências externas além da biblioteca padrão do Go.

## Como Jogar

Para jogar o "hangman-game", siga as etapas abaixo:

1. Clone ou faça o download do repositório do projeto para o seu computador.

2. Execute o programa com o seguinte comando:

   ```shell
   go run main.go
   ```

3. O jogo escolherá aleatoriamente uma palavra do arquivo "words.json" e exibirá uma série de astericos representando as letras da palavra oculta.

4. Você deve adivinhar as letras uma por uma ou digitando a palavra inteira no console. O jogo informará se a letra está presente na palavra ou não.

5. Continue adivinhando letras até adivinhar todas as letras da palavra ou até esgotar suas tentativas. O jogo informará se você ganhou ou perdeu.

## Arquivo "words.json"

O arquivo "words.json" contém uma matriz de palavras que o jogo utiliza para escolher aleatoriamente a palavra oculta. Aqui está um exemplo de seu conteúdo:

```json
{
  "words": [
  "gato",
  "cachorro",
  "elefante",
  "leão",
  "tigre",
  "girafa",
  "rinoceronte"
  ]
}

```

Você pode adicionar ou remover palavras conforme necessário, desde que mantenha o formato JSON correto.
