#### O que é o Apache kafka

- "O Apache Kafka é uma plataforma distribuída de streaming de eventos open-source, que é utilizada por milhares de empresas para uma alta performance em pipeline de dados, stream de analytics, integração de dados e aplivações de missão crítica" - https://kafka.apache.org


#### Prós

- Altíssimo throughput.
- Latência extremamente baixa (2ms).
- Escalável.
- Armazenamento.
- Alta disponibilidade.
- Se conecta com quase tudo.
- Bibliotecas prontas para as mais diversas tecnologias.
- Ferramentas open-source.


#### Obs

- Remendação de no mínimo 3 brokers (clusters).

#### Tópicos

- Tópico é o canal de comunicação responsável por receber e disponibilizar os dados enviados para o kafka (+- como se fosse uma fila, diferente de uma fila normal o consumidor não retira e nem bloqueia a mensagem ao ler ela).

#### Anatomia de um registro

- Offset 0 -> 
    - Headers: Meta data, opcional
    - Key:
    - Value: Conteúdo da mensagem.
    - Timestamp

#### Partições

- Cada tópico pode ter uma ou mais partições para conseguir garantir a distribuição e resiliência de seus dados.
- Uma key pode ser informada na mensagem, mensagens com a mesma key tendem a ir para a mesma partição, essa seria a forma de garantir sequecialidade entre duas mensagens.
- É possível definir uma propriedade "replicator factor", isso faz com que um número X de redundâncias sejam criadas para uma partição nos outros clusters.

#### Produtor

***Garantias de entrega***

- Ack 0: Sem confirmação (fire and forget).
- Ack 1: Líder confirma o recebimento.
- Ack -1: All, todos o líder e os followers (réplicas da partição) precisam confirmar o recebimento.
  
  ----

- At most once: Melhor performance, mas pode perder algumas mensagens.
- At least once: Perfomance moderada, mas pode duplicar mensagens.
- Exacly once: Pior performance, msa garante apenas um envio.

***Produtor Indepotente***

Em um cenário em que haja algum problema (de rede por exemplo), e o ack não seja realizado, então o produtor posta novamente a mensagem, mas por algum motivo o post anterior também é reprocessado, gerando uma duplicidade da mansagem.

- Off: O kafka não vai verificar duplicidade da mensagem, e ambas passarão.
- On: O kafka irá verificar duplicidade da mensagem (impactando em performance), e descartará uma das mensagens e também irá garantir que a mensagem entre na ordem correta.

#### Consumidor

***Consumer Groups***

- Caso não haja grupos, o consumer irá consumer de todas as partições de um tópico, caso um grupo de consumidores seja definido, eles irão se dividir e cada partição terá um consumidor do grupo "responsável" por consumi-la.
- No cenário de grupos, caso haja mais consumidores que partições, o excendente ficará parado (idle).
- No cenário de grupos, caso haja mais partições do que consumidores, um consumidor irá consumir mais de uma partição.