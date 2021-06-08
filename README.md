# grettings-go
O intuito desse projeto é estudar #golang, documentei passo a passo do meu aprendizado aqui.

# Aprendendo #golang
## Dia 1
Segui o tutorial da documentação oficial: https://golang.org/doc/, instalei o Go e resolvi buscar uma IDE para me ajudar, como eu vim do mundo Java e estou muito acostumado com o #Intellij, baixei o #goland e estou no Free Trial de 30 dias.
A primeira impressão da linguagem é claramente a proximidade com C e C++ devido a palavras reservadas como struct e métodos com assinatura Printf, Scanf, etc..
Criei um arquivo go.mod e é nele que ficam as dependências de outras libs. Percebi que o Go armazena o código fonte das libs baixadas no home do usuário na pasta go, e assim como outras linguagens, existe uma casa para essas bibliotecas: https://go.dev/

Estou aprendendo agora uma nova linguagem, escrevi algo de errado? Por favor, comente.

## Dia 2
Seguindo o tutorial da documentação oficial: https://golang.org/doc/tutorial/handle-errors e https://golang.org/doc/effective_go#multiple-returns, comecei a lançar e tratar erros, e descobri que #golang permite que uma função retorne mais de um dado e isso é utilizado para retornar erros.

Também fui apresentado ao slice em https://golang.org/doc/tutorial/random-greeting e https://blog.golang.org/slices-intro. É um array "dinâmico". Confesso que achei um pouco complexo, mas entendi o mecanismo. Novamente a linguagem me lembra muito C e C++, mas agora não em sintaxe, mas sim na maneira como os dados são manipulados.

Achei muito interessante a regra de que funções que iniciam com letras maiúsculas são public e iniciadas em minúsculo são privadas.

Estou aprendendo agora uma nova linguagem, escrevi algo de errado? Por favor, comente.

## Dia 3
Fui apresentado ao map, similar a outras linguagens, até aqui nada de mais: https://golang.org/doc/tutorial/greetings-multiple-people
Porém ao ler o blog sobre map: https://blog.golang.org/maps, também fui apresentado ao mecanismo de Mutex do #golang, algo que fica normalmente escondido em outras linguagens com palavras reservadas, aqui está claramente exposto no código, prefiro as abordagens de outras linguagens.
Estou aprendendo agora uma nova linguagem, escrevi algo de errado? Por favor, comente.

## Dia 4
Um tutorial de linguagem que ensina testes unitários ganha meus respeitos: https://golang.org/doc/tutorial/add-a-test

Também aprendi a compilar e instalar o módulo do projeto: https://golang.org/doc/tutorial/compile-install

Final do tutorial, resolvi então brincar um pouco com o código escrito até aqui, e pesquisei um pouco sobre leitura de dados no console e encontrei um guia interessante: https://zetcode.com/golang/readinput/

Estou aprendendo agora uma nova linguagem, escrevi algo de errado? Por favor, comente.

## Dia 5
Com o tutorial oficial finalizado resolvi seguir o conselho o [Alex Rios](https://twitter.com/alextrending) e seguir o gitbook: https://quii.gitbook.io/learn-go-with-tests/ Do meu ponto de vista, "nenhum" código deveria ir para a branch master sem testes, então me identifiquei muito com a metodologia de TDD do livro.

Já no início, na instalação do Go, o guia recomenda instalar um linter para Go, sou muito fã de linters pois eles garantem que estamos utilizando boas práticas, e para quem está aprendendo uma nova linguagem, nada melhor do que aprender da maneira correta. O linter pode ser encontrado em: https://golangci-lint.run/

Segui lendo o tutorial e refatorando alguns pontos do código já escrito até aqui, adicionando constantes e organizando melhor os testes. Também descobri que #golang possui um mecanismo nativo para benchmarks. Muito bom!
