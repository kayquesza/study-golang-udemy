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