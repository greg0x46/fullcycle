
## Para o serviço funcionar é necessário criar um tabela no banco
- `docker-compose up -d`: subir containers
- `docker exec -it db bash`: entrar no container do mysql
- `mysql -u root -p`: logar no mysql cli.
- `use nodedb`: seleciona banco de dados.
- `create table people(id int not null auto_increment, name varchar(255), primary key (id));`

## Exemplos para subir container utilizando o comando run
**RUN WITHOUT DOCKER FILE**: `docker run --rm -it -v $(pwd)/app/:/usr/src/app -p 3000:3000 node:15 bash`
**BUILD**: `docker build -t greg0x46/hello-express .`
**BUILD Dockerfile.prod**: `docker build -t greg0x46/hello-express . -f Dockerfile.prod`
**RUN BUILDED IMAGE**: `docker run -p 3000:3000 --name hello-express greg0x46/hello-express`