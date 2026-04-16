## HTML (HyperText Markup Language)

O HTML é a linguagem de marcação utilizada para estruturar o conteúdo das páginas web. Ele utiliza **Tags** para informar ao navegador como os elementos devem ser exibidos.

### Conceitos Fundamentais
- **Marcação vs Programação**: O HTML não possui lógica (if, loops, variáveis). Ele apenas descreve a estrutura.
- **Tags**: Elementos delimitados por `< >`. A maioria possui uma abertura e um fechamento (ex: `<h2>Conteúdo</h2>`).
- **Semântica**: O uso correto das tags (`h1` para títulos principais, `p` para parágrafos) ajuda na acessibilidade e no rankeamento.

### Integração com Go
No ecossistema Go, utilizamos o pacote nativo `html/template` para:
1. Carregar arquivos HTML externos.
2. Renderizar esses arquivos como resposta a uma requisição HTTP.
3. Injetar dados dinâmicos do backend diretamente no frontend de forma segura.

### Visão de Segurança (Infosec)
Para um profissional de segurança, entender HTML é o primeiro passo para compreender ataques como **XSS (Cross-Site Scripting)**, onde scripts maliciosos são injetados dentro de tags HTML para roubar cookies ou sessões de usuários.

### Renderização Dinâmica com Templates

O pacote de templates do Go permite que transformemos ficheiros HTML estáticos em páginas dinâmicas através da injeção de estruturas de dados.

#### Fluxo de Trabalho
1. **Carregamento**: Usamos `template.ParseGlob` ou `template.ParseFiles` para ler os ficheiros do disco.
2. **Definição de Dados**: Criamos uma `struct` ou `map` com as informações que queremos exibir.
3. **Execução**: O método `ExecuteTemplate` funde os dados com o HTML e envia o resultado para o `http.ResponseWriter`.

#### Segurança: text/template vs html/template
Embora ambos funcionem de forma semelhante, para páginas web **deve-se sempre usar `html/template`**.
- **Context-Aware Escaping**: O pacote `html/template` limpa automaticamente os dados inseridos. Se um utilizador tentar enviar um script malicioso (XSS) num campo de nome, o Go irá transformá-lo em texto inofensivo antes de renderizar.

#### Sintaxe Básica nos Ficheiros .html
- `{{ . }}`: Refere-se ao objeto de dados completo passado para o template.
- `{{ .Campo }}`: Acede a um campo específico de uma struct (deve ser exportado/letra maiúscula em Go).