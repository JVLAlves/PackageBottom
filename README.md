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
method (l methodi) FirstMethodfromTheCodeTheFi
```

## SHELL SCRIPTS

### NEW GO (NEWGO)
O script newgo é um bash script feito para facilitar a criação de arquivos .go. Através de um único comando no terminal é possível criar um novo diretório e um novo arquivo Go de uma só vez. Para isso é necessário seguir os seguintes passos:

1. Baixar o script e lhe-dar as permissões necessárias.
```bash
chmod +x newgo
```

2. Copiar ou mover o script para dentro da pasta /usr/bin/ ou /usr/local/bin/
```bash
sudo cp newgo /usr/local/bin/
```
ou
```bash
sudo mv newgo /usr/bin/
```

3. Com o comando script copiado (ou movido) chame-o, seguido do nome desejado para o arquivo
```bash
newgo exemple
```
Se desejar colocar esse novo conjunto dirétorio/arquivo dentro de outro diretório basta dizer o caminho:
```bash
newgo exempledir/exemple.go
```
Ainda, caso não seja informado o nome do arquivo, ele irá criar um conjunto diretório/arquivo com uma sequência de números aleatórios.
```bash
joes@DNA-83203:~/plans$ newgo
198402/198402.go
```

### GIT HERE (GITH)
O script gith é um bash script que facilita criação de repositórios locais. Com um único comando - e um link - encurta-se todo o processo de "init", "remote add", "pull" e afins. Para que o comando funcione de acordo com o planejado, siga os seguintes passos:

1. Baixar o script e lhe-dar as permissões necessárias.
```bash
chmod +x gith
```

2. Copiar ou mover o script para dentro da pasta /usr/bin/ ou /usr/local/bin/
```bash
sudo cp gith /usr/local/bin/
```
ou
```bash
sudo mv gith /usr/bin/
```
3.  Com o comando script copiado (ou movido) chame-o e espere-o perguntar qual o link do repositório.
```git
Type the repository link: git@github.com:JVLAlves/PackageBottom.git
```
Ao realizar esses passos, o comando realiza o "init", "remote add" e "pull" automaticamente. 



