#### Confiabilidade

- Consumer acknowledgement (ack): Confirmação por parte do consumer.
  - Basic.Ack: Consumidor confirma que recebeu e conseguiu processar e a mensagem é removida da fila.
  - Basic.Reject: Consumer rejeita mensagem e ela volta pra fila (é possível fazer um requeue false, de modo que a mensagem será descartada).
  - Basic.Nack: Igual ao reject porém pode ser feito em lote.
- Publisher confirm: RabbitMq confirma que recebeu a mensagem do produtor.
  - Produtor deve atribuir um id para a mensagem.
  - Exchange recebe a mensagem e faz um Ack ou Nack para o produtor passando como referência o id atribuido pelo próprio.
- Filas e mensagens duráveis / persistidas (lazy queue)

#### Tipos de Exchange

- Direct: encaminha diretamente a uma queue a partir do routing key.
- Fanout: encaminha para todas as queues relacionadas à exchange.
- Topic: encaminha para as queues cuja o rounting key dê match (semelhante a expressão regular).
- Headers: encaminha para queue de acordo com o header específicado na mensagem (menos utilizado).


#### Queues

- FIFO: Padrão 
- Prioridade: É possível alterar o padrão FIFO e definir prioridades.
- Propriedades:
  - Durable: Se ela deve ser salva mesmo do restart do broker.
  - Auto-delete: Removida automaticamente quando o consumer se desconecta.
  - Expiry: Define o tempo que não há mensagens ou clientes consumindo, após esse tempo a filka é excluída.
  - Message TTL: Tempo de vida da mensagem.
  - Max length ou bytes: Quantidade de mensagens ou tamanho máximo permitido em bytes.
    - Overflow:
      - Drop head: remove a última.
      - Reject publish: rejeita novas mensagens.
  - Exclusive: Somente o channel que criou pode acessar.
  
#### Dead letter queues

- Algumas mensagens não conseguem ser entregues por qualquer motivo.
- São encaminhadas para uma Exchange específica que roteia as mensagens para uma dead letter queue.
- Tais mensagens podem ser consumidas e averiguadas posteriormente.


#### Lazy Queueues

- Mensagens são armazenadas em disco.
- Exige alto I/O.
- Quanto há milhões de mensagem em uma fila, por qualquer motivo, há a possibilidade de liberar a memória, armazenando especificamente as mensagens da fila em questão em disco.

#### Utilidades

- RabbitMQ Simulator: http://tryrabbitmq.com/