## HTTP (HyperText Transfer Protocol)

O HTTP é a base de toda a comunicação na web. Ele funciona através de um modelo de **Requisição (Request)** e **Resposta (Response)** entre um Cliente e um Servidor.

### Componentes Básicos
- **URI/URL**: O endereço do recurso que queremos acessar (ex: `/usuarios`).
- **Métodos (Verbos)**: Definem a intenção da mensagem.
  - **GET**: Busca dados (não deve alterar o estado do servidor).
  - **POST**: Cria novos registros.
  - **PUT**: Atualiza registros existentes.
  - **DELETE**: Remove dados.

### Implementação em Go
Utilizamos o pacote nativo `net/http` para subir servidores robustos de forma simples.

- **HandleFunc**: Mapeia uma rota (URI) para uma função de processamento.
- **ResponseWriter**: Interface para construir a resposta enviada ao cliente.
- **Request**: Struct que contém todos os dados da requisição recebida.

### Execução do Servidor
O servidor é iniciado com `http.ListenAndServe(":PORTA", nil)`. Em ambiente de desenvolvimento, é comum usarmos portas como `5000`, `8080` ou `3000`.