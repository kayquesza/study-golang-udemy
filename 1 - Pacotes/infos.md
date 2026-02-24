## Packages

Ao lidar com Pacotes em Go, se cria um *Módulo*.
**Módulo:** É um conjunto de pacotes que compõem o projeto, sejam pacotes internos (criados) dentro do projeto ou pacotes externos já existentes. 
*Usado quando:* Se tem mais de 1 pacote no mesmo projeto. 

*Para criar o Módulo (**no terminal**):*
go mod init <nomeDoModulo> 
    *exemplo:* go mod init modulo

**Nota:** Responsável por centralizar todas as dependências do projeto. 



### Comandos *importantes* via terminal

#### Para rodar/compilar o projeto:
go run <nome> 
    *exemplo:* go run main.go


#### Contendo o módulo, é possível rodar/compilar de outra forma. É possível utilizar o: *go build* que criará um arquivo com o nome do módulo já criado (será um arquivo binário do projeto compilado, será criado na raiz) e, para roda-lo no terminal, se utiliza essa forma no terminal: *./<nomeDoArquivoGerado>*.
    *exemplo:* 
        1 - go build
        2 - *é criado um arquivo binário executável chamado 'modulo'.*
        3 - ./modulo
        4 - *a saída será o conteúdo compilado do estado atual do projeto.*
**Nota:** O arquivo não se atualiza sozinho, é necessário utilizar o *go build* novamente para recompilar o arquivo.


#### Comando `go install` fará o mesmo que o *go build*, mas o arquivo gerado ficará na raiz onde o **Go** foi instalado. 


#### Para utilizar um pacote externo 
go get <pacote>
    *exemplo:* go get github.com/badoux/checkmail
**Nota:** Será instalado também qualquer dependência que o próprio pacote contiver; atualizando o arquivo `go.mod`. 


#### Comando `go mod tidy` é responsável por retirar toda e qualquer dependência que não esteja sendo utilizada e baixar qualquer depedência que esteja sendo solicitada e ainda não foi instalada.



### Informações relevantes

- **Funções**
Funções com *letras minúsculas* não são exportadas a outros pacotes.
Funções com *letras maiúsculas* são exportadas para outros pacotes e, além disso, é de boa prática inserir um comentário acima da mesma para explicar o que ela faz. 