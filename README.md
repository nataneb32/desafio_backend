# Desafio Backend
## Objetivo
Desenvolver uma aplicação (somente o backend) que possibilite realizar o cadastro de
hóspedes e o check in. Queremos ver como você resolve problemas no seu dia-a-dia. Não
há necessidade de desenvolver o frontend da aplicação, vamos utilizar o Postman para
testar sua aplicação.
## Requisitos funcionais
- CRUDL para o cadastro de hóspedes;
- Deve ser possível buscar hóspedes cadastrados pelo nome, documento ou telefone;
- Consultar hóspedes que já realizaram o check in e não estão mais no hotel;
- Consultar hóspedes que ainda estão no hotel;
- As consultas devem apresentar o valor (valor total e o valor da última
hospedagem) já gasto pelo hóspede no hotel.

## Regras de negócio
- Uma diária no hotel de segunda à sexta custa R$ 120,00;
- Uma diária no hotel em finais de semana custa R$ 150,00;
- Caso a pessoa precise de uma vaga na garagem do hotel há um acréscimo diário,
sendo R$ 15,00 de segunda à sexta e R$ 20,00 nos finais de semana;
- Caso o horário da saída seja após às 16h30 deve ser cobrada uma diária extra.
## Expectativas
- Desenvolva um sistema em Go, com banco de dados PostgreSQL e API REST.
- Desenvolva o problema utilizando frameworks, bibliotecas que você julgue
adequado para resolver o problema;
- Caso seja preciso realizar algum build ou algum passo extra para gerar a sua
solução, você deve detalhar o que deve ser feito no arquivo README.md de seu
projeto.

## Como dar build?
```/bin/sh
sudo docker-compose build
sudo docker-compose up
```

## API com exemplos.

### POST /guest - Cria um hospede.
```/bin/sh
curl --request POST \
  --url http://localhost:8080/guest \
  --header 'content-type: application/json' \
  --data '{
	"nome": "Foo",
	"telefone": "11111111",
	"documento": "12345/123"
}'
```

retorna

```json
{
  "id": 2,
  "nome": "Foo",
  "telefone": "11111111",
  "documento": "12345/123",
  "checkins": null
}
```

### GET /guest/:id - Consultar um hospede pelo id.
```
curl --request GET \
  --url http://localhost:8080/guest/2
```

Return

```json
{
	"guest":{
		"id":2,
		"nome":"Foo",
		"telefone":"2499999999",
		"documento":"12341-123",
		"checkins":[
			{
			"id":2,
			"hospede":2,
			"dataEntrada":"2020-08-29T08:00:00Z",
			"dataSaida":"2020-09-30T12:00:00Z",
			"adicionalVeiculo":false
			}
		]
		},"lastBill":414000,"totalBill":1242000}
```

### GET /guests - Faz uma pesquisa com querys pelo hospede.
```/bin/sh
curl --request GET \
  --url 'http://localhost:8080/guests?nome=Foo&limit=10&page=1&telefone=11111111&documento=12345%2F123&=' 
``` 
retorna
```json
{
  "guests": [
    {
      "id": 6,
      "nome": "Foo",
      "telefone": "11111111",
      "documento": "12345/123",
      "checkins": []
    }
  ],
  "totalPages": 1
}
```

### PUT /guest/:id - Atualiza os valores do hospede com id = :id
```/bin/sh
curl --request PUT \
  --url http://localhost:8080/guest/2 \
  --header 'content-type: application/json' \
  --data '{
	"nome": "Foo2",
	"telefone": "11111111",
	"documento": "12345/123"
}'
```

retorna 200 OK

### DELETE /guest/:id - Deleta o hospede com id = :id
```/bin/sh
curl --request DELETE \
  --url http://localhost:8080/guest/2
```

retorna 200 OK
