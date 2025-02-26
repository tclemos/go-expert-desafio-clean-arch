# go-expert-desafio-clean-arch

Este repositório possui 3 componentes, são eles:

- [GRPC](./cmd/grpc/main.go): Um servidor GRPC que utiliza banco de dados Sqlite3 para obter os dados de ORDERS.
- [REST](./cmd/rest/main.go): Um servidor REST que se conecta ao servidor GRPC para obter os dados de ORDERS.
- [GRAPHQL](./cmd/graphql/main.go): Um servidor GRAPHQL que se conecta ao servidor GRPC para obter os dados de ORDERS.

Para executar os componentes acima basta ter o `Make`, `Docker` e `Compose` instalados e executar o seguinte comando:

```bash
make run
```

Verifique o conteúdo do arquivo [Makefile](./Makefile) para mais comandos.

A configuração de cada componente está dentro do arquivo `.env` de cada pasta e os mesmos já estão configurados para funcionar com o comando acima.

- GRPC: <http://localhost:50051>
- REST: <http://localhost:3000>
- GRAPHQL: <http://localhost:8080>

As chamadas HTTP de teste estão no arquivo [order.http](./test/order.http)
As chamadas GraphQL de teste estão no arquivo [order.graphql](./test/order.graphql)
