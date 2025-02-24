# go-expert-desafio-clean-arch

Este repositório possui 2 componentes, são eles:

- [GRPC](./cmd/grpc/main.go): Um servidor GRPC que utiliza banco de dados Sqlite3
- [REST](./cmd/rest/main.go): Um servidor REST que utiliza se conecta ao servidor GRPC

Para executar os componentes acima basta ter o Make, Docker e Compose instalados e executar o seguinte comando:

```bash
make run
```

A configuração de cada componente está dentro do arquivo `.env` de cada pasta e os mesmos já estão configurados para funcionar com o comando acima.

As URLs de teste estão no arquivo [order.http](./test/order.http)
