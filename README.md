# PackageBottom

## BEM-VINDO A ÁREA PARA TESTE DE PROGRAMAS NUCLEARES :bomb::boom: 
Este repositório é destinado para o armazenamento de funções para utilição geral, isto é:

* Facilitar Debugging;
* Auxiliar na criação de arquivos e diretórios;
* Agilizar ou aprimorar processos no interior de programas;
* etc.

## FUNÇÕES

### FUNCTION PARSERER
Este programa utiliza Expressões Regulares específicas para buscar funções e métodos em um código Go. útil para Debugging e análise.

Ao executar, dois campos input aparecerâo a vocë. A primeira diz respeito ao nome do diretório que contém o arquivo .go, por exemplo:
```go
Type the directory name: /home/joes/plans/golang
```

O segundo é relativo ao nome do arquivo. Por exemplo:
```go
Type the file name: Exemple.go
```
Ao completar a execução você irá receber um arquivo com o mesmo nome, mas com a extensão .txt. As funções e métodos estarão listadas no seguinte formato:

```txt
func FirstFunctionFromTheCode
method (l methodi) FirstMethodfromTheCode
method (l methodi) FirstMethodfromTheCodeTheSecond
func FirstFunctionFromTheCodeTwo
func FirstFunctionFromTheCodeThree
func FirstFunctionFromTheCodeFour
func FirstFunctionFromTheCodeFive
method (l methodi) FirstMethodfromTheCodeTheFif
```


