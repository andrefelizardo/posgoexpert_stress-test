# posgoexpert_stress-test

## Build da Imagem Docker

Utilize o comando abaixo para buildar a imagem Docker. Este comando irá utilizar o Dockerfile presente para compilar sua aplicação Go em uma imagem standalone.

```bash
docker build -t stress-test .
```

## Rodando a Aplicação com Docker

Após buildar a imagem, você pode executar a aplicação com o Docker. Use o seguinte comando para rodar a aplicação, passando os parâmetros necessários para os testes:

```bash
docker run --rm stress-test --url=http://google.com --requests=1000 --concurrency=10
```

- `--rm` remove o container assim que ele finaliza sua execução.
- Os parâmetros `--url`, `--requests` e `--concurrency` são passados para a aplicação e definidos conforme suas necessidades.
