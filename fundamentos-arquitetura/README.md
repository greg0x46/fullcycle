# Fundamentos Arquitetura de Software

### Definição 
"Organização de um istema, conteomplando seus componentes, os relacionamentos entre estes e com o ambiente, e os princípios que governam seu projeto e evolução." (ISO 42010:2011)

### Pilares
- Organização de um sistema
- Componentização
- Relacionamento entre sistemas
- Governança
- Ambiente
- Projeto
- Projeção
- Cultura

---- 

### Frameworks 
São ferramentas e métodos que nos ajudam a focar essencialmente no objetivo final. Frameworks nos ajudam a definir um padrão de trabalho.
- The TOGAF Standard
  - Framework conceitual
  - Definição dos processos de arquitetura
  - Conceitos e nomenclaturas
  - Visão geral de tipos de arquiteturas
    - Negócio
    - Sistema de informação
    - Tecnologia
    - Planos de migração
- ISO 420210
  - Lançada em 2011
  - Simplificada em relação ao TOGAF
  - Formaliza funcamentos da área de arquitetura de software

----

### Sistemas Monolíticos

- Quando não vale apena
  - Times grandes
  - Necessidade de escalar todo sistema pelo de um área em específico esteja com pico de utilização
  - Risco de um deploy completo começa a se elevar
  - Necessidade de utilizar tecnologias diferentes

- Contras
  - Tudo dentro do mesmo sistema
  - Alto acoplamento
  - Processo de Deploy "completo" a cada mudança
  - Normalmente usa uma tecnologia
  - Um problema afeta todo o sistema
  - Maior complexidade para times
- Prós
  - Menos complexidade na maioria dos casos

----

### O que é um serviço

- Disponibiliza informação
- Realiza transações
- Resolve problemas do negócio
- Independente de tecnologia ou produto
- Pode estabalecer comunicação com diversos "clientes"

### SOA: Arquitetura orientada a serviços

- Serviços normalmente maiores baseados em funcionalidades
- Necessidade do ESB (Enterprise Service Bus)
- Single point of failure (se o esb cair, cai tudo)
- Compartilhamento de banco de dados é comum
- Muitas vezes também podem ter sistemas monolíticos sendo utilizados como serviços

#### Arquitetura baseada em microserviços

- Prós
  - Serviços pequenos com poucas responsabilidades
  - Maior tolerância a falhas
  - Totalmente indepedente
  - Cada serviço possui seu próprio banco de dados
  - Comunicação síncrona ou assíncrona
- Contras
  - Arquitetura complexa
  - Custo mais elevado
  - Necessidades me mais equipes para manter
  - Sistema precisa ser grtande o suficiente para justificar
  - Gera problemas que normalmente você não tinha antes
  - Monitoramento complexo

###### Principais características (Martin Fowler - http://bit.ley/fowler-microservicos)

*Componentização via serviços*
- Services dos microserviços != services da OO
- Componente é uma unidade de software independente que pode ser substituída ou atualizada
- Organização em torno do negócio
    - Cada produto pode ter um ou mais serviços
    - Desvantagens:
        - Chamadas externas são mais custosas do que chamadas locais
        - Cruzamento entre componentes pode se tornar complexo
        - Transação entre serviços são "grandes desafios"
        - Mudanças bruscas em regras de negócio podem afertar diversos serviços tornando o processo difícil de ser refeito    

*Organização em tono do negócio*
- ...

*Estrutura baseada em produtos, não em projetos*
- ...

*Smart endpoints & dumb pipes*
  - Exposição de APIs (ex: REST)
  - Comunicação entre serviços
  - Comunicação síncrona ou assíncrona
  - Utilização de mensageria (ex: RabitMQ)
  - Garantia de que um serviço foi executado baseado na execução das filas
 
 *Governança descentralizada*
  - Ferramenta certa para o trabalho certo, tecnologias podem ser definidas baseadas na necessidade do produto
  - Diferentes padrões entre squads
  - Contratos de interface de forma independente
  - Automação de infraestrutura
  
*Descentralização no gerenciamento de dados*
  - ...

*Automação de infraestrutura*
  - Cloud computing
  - Testes automatizados
  - Continous delivery
  - Contionous integration
  - Load balancer / Autoscaling

*Desenhado para falhar*
  - Tolerância a falha
  - Serviçõs que se comunicam precisam de fallback
  - Logging
  - Monitoramento em tempo real
  - Alarmes

*Design evolutivo*
- Produtos bem definidos podem evoluir ou serem extintos por razações de negócio
- Gerenciamento de versões
- Replacement and upgradeability
  
----


#### Escalando software
- Escala vertical (aumento de recursos computacionais)
- Escala horizontal (aumento de maquinas (clusters) atrás de um load balancer)
  - Disco efêmero
  - Servidor de aplicação separado do servidor de assets, Upload / Gravação de arquivos
  - Cache centralizado
  - Sessões centralizadas
  - Tudo pode ser destruído e criado facilmente ?
----

#### API gateways
- Uma API gateway recebe todas as chamadas de APIs dos clientes e então as roteia para os microserviços correspondentes...
- Em alguns casos ela também é responsável por realizar processos de verificação de segurança, como autenticação e autorização.

#### Service discovery
- Processo de service discovery é responsável por provcer mecanismos de identificação dos serviços disponíveis e suas instâncias
- Client side
  - Service registry
    - Health check
- Server side
  - Load balancer
    - Service registry
      - Health check
- Ferramentas populares
  - Netflix eureka
  - Consul
  - Etcd
  - ZooKeeper