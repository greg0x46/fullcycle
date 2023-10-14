## Geral 

#### Comandos básico
- `docker`: lista comandos e suas descrições.
- `docker ps`: lista containers ativos.
- `docker ps -a`: lista todos containers.
- `docker start <container>`: inicia um container.
- `docker stop <container>`: para um container.
- `docker rm <container>`: deleta um container.
- `docker rm <container> -f`: força a exclusão de um container
- `docker ps -a -q`: lista o id de todos containers
- `docker build <container-registry-user>/<image>:<version> <docker-file-path>`: constroi um imagem a partir de um dockerfile padrão.
- `- <container-registry-user>/<image>:<version> <docker-file-path> -f Dockerfile.prd`: constroi um imagem a partir de um dockerfile específico.
- `docker logs <container>`: exibe logs de um container.
- `docker login`: loga no dockerhub.
- `docker logout`: desloga do dockerhub.

#### Docker run
- `docker run <image>`: sobe um container a partir da imagem (tag default latest).
- `docker run <image>:<tag>`: sobe um container a partir de uma imagem e de uma tag.
- `docker run -i <...>`: -i significa interativo, "anexa" o terminal ao container.
- `docker run -t <...>`: -t significa tty, permite rodar comandos no container.
- `docker run -it <...> bash`: executa um container com as opções -i e -t e abre o bash do container no terminal.
- `docker run --rm <...>`: remove automaticamente um container após sua execução.
- `docker run -p <host-port>:<container-port> <...>`: cria um link entre uma porta do docker host e docker container.
- `docker run -d <...>`: -d segnifica desanexar, libera o terminal (roda em background).
- `docker run --name <name>`: define um nome para o container
- `docker run -v <docker-host-path>:<docker-container-path> <...>`: mapeia uma pasta do docker host para o container, cria a pasta caso não exista.
- `docker run --mount type=bind,source=<docker-host-path>,target=<docker-container-path>`: melhor forma de mapeiar uma pasta do host para o container (não cria pastas).
- `docker run --mount type=volume,source=<volume-name>,target=<docker-container-path>`: mapeia um volume no container.

## Volume

#### Basic commands
- `docker volume`: lista comandos relativos à volumes
- `docker volume ls`: lista volumes.
- `docker volume create <name>`: cria um volume.
- `docker volume inspect <name>`: exibe detalhes do volume.
- `docker volume prune`: delete volumes que não estão em uso.

## Images

####Basic commands
- `docker pull <image>:<tag>`: baixa uma imagem.
- `docker push <container-registry-user>/<image>:<version>`: faz upload de uma imagem para o dockerhub.
- `docker rmi <image>:<tag>`: deleta uma  imagem


## Dockerfile

#### File sintaxe
- `FROM <image>:<version>`: define uma imagem base.
- `ENV <var> <value>`: define uma variável de ambiente.
- `USER <user>`: define o usuário do container (ele precisa existir)
- `WORKDIR <container-path>`: define "pasta principal" do container.
- `RUN <command>`: executa um comando.
- `COPY <docker-host-path> <docker-container-path>`: copia um arquivo do host para o container, "\<docker-host-path\>" é relativo ao path do Dockerfile.
- `EXPOSE <port>`: expõe uma porta do container
- `ENTRYPOINT ["echo", "Hello World"]`: é executado no final do script do Dockerfile, antes do CMD, esse comando sempre será executado e pode ser utilizado para manter o container rodando.
- `CMD ["echo", "Hello World"]`: é executa no final do script do Dockerfile, podde ser sobrescrito pelo comando run, exemplo `docker run <...> bash` nesse caso o "bash" iria sobrescrever o `["echo", "Hello World"]`.


## Network 

#### Types of network
- **bridge** (default): compartilha rede entre containers
- **host**: mescla rede dos containers e do host.
- **overlay**: mescla rede entre diferente containers.
- **maclan**:  define um MAC address para cada container, simulando uma rede fisica.
- **none**: sem network.

#### Comandos de network    
- `docker network`: exibe uma lista de comandos relativo à network.
- `docker network create --driver bridge mynetwork`: exemplo do comando utilizado para criar uma network.
- `docker run <...> --network mynetwork`: examplo de comando para subir um conetainer anexando uma network.
- `docker network connect <network> <container>`: conecta um container a uma network.

## Hacks
- `docker rm $(docker ps -a -q) -f` : remove todos containers
- `curl http://host.docker.internal:8000`: "http://host.docker.internal" é o endereço do docker host, ele pode ser utilziado para acessar um serviço do host a partir de um container.
- `docker run --mount type=bind,source="$(pwd)"/html,target=<docker-container-path>`: "$(pwd)" pode ser usado para obter o diretório atual.
- `RUN apt-get update && apt-get install -y`: No contexto de um Dockerfile "&&" pode ser utilizado para concatenar comandos que devem ser executados sequencialmente.
-`\`: No contexto de um Dockerfile "\\" pode ser utilizado para concatenar linhas, de forma a executar comandos em multipas linhas como se estivessem na mesma.
`````
RUN apt-get update && \ 
    apt-get install -y
`````
