## Desafio Técnico: Stress-Test

Sistema que realiza testes de carga enviando múltiplas requisições a uma URL e gera um **relatório com tempo total de execução, total de requisições e distribuição dos códigos de status HTTP (200, 404, 500, etc.).**

###  Tecnologias
- **Go**
- **Fiber** (API HTTP)
- **Docker** (Execução isolada)
- **Goroutines** (Concorrência)

---

### Clone o repositório (https):
 https://github.com/saraqueirozs/challenge-goexpert-stress-test

### Executando com Docker

```sh
docker build -t load-tester .
docker run --rm -p 3000:3000 load-tester  # Inicia API na porta 3000


 ### Testes Recomendados

Para testar diferentes códigos de resposta HTTP, utilize as URLs abaixo.

```sh
docker run --rm load-tester --url=https://httpstat.us/200 --requests=10 --concurrency=2

docker run --rm load-tester --url=https://httpstat.us/404 --requests=10 --concurrency=2

docker run --rm load-tester --url=https://httpstat.us/500 --requests=10 --concurrency=2

docker run --rm load-tester --url=https://httpstat.us/503 --requests=10 --concurrency=2

docker run --rm load-tester --url=https://httpstat.us/200?sleep=5000 --requests=10 --concurrency=2

docker run --rm load-tester --url=https://google.com --requests=10 --concurrency=2

