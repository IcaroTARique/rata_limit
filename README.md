# Rate Limit
O Rate Limit é uma técnica utilizada para limitar 
a quantidade de requisições que um cliente pode
fazer a um serviço em um determinado período de
tempo. O Rate Limit é uma técnica muito utilizada 
para proteger serviços de abusos, como ataques de
negação de serviço (DDoS) e para garantir a 
qualidade do serviço.

## Rate Limit usando Redis
O Redis é uma escolha popular para implementar
rate limiting (limitação de taxa) devido às suas
características de alta performance, baixa latência
e suporte a operações atômicas.

## Papel do Redis
O Redis é usado para armazenar contadores que 
mantêm o número de requisições feitas por um 
usuário ou IP em um intervalo de tempo específico.
Esses contadores são incrementados a cada requisição.


### Operações Atômicas
O Redis suporta operações atômicas, garantindo que 
incrementos e verificações de contadores sejam feitos
de maneira consistente sem condições de corrida 
(race conditions). Comandos como INCR, EXPIRE, GET, 
e SET são frequentemente usados.

### Expiration
Redis permite definir um tempo de expiração para chaves,
que é útil para implementar janelas de tempo deslizantes
(sliding windows) ou fixas (fixed windows) no rate 
limiting. Por exemplo, você pode usar o comando EXPIRE 
para garantir que um contador reseta após um intervalo
de tempo específico.

### Alta performance e Baixa Latência
Redis é extremamente rápido e pode lidar com um grande
número de operações por segundo, tornando-o ideal para
sistemas de rate limiting que precisam ser escaláveis
e responsivos.

## Utilização
Para rodar a aplicação, basta rodar o seguinte comando
na raiz do projeto:
```bash
docker-compose up -d
```
e veja a mágica acontecer.

## Chamadas de teste
Você pode testar a aplicação fazendo chamadas para o:
```bash
curl -X GET "http://localhost:8082/status" -H "API_KEY: any-key"
```
