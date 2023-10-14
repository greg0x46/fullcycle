### Enviroment


##### BUILD IMAGE
~~~
docker build -t fc-graphql .
~~~

##### RUN SERVER
~~~
docker run --rm -p 9090:9090 -v ~/fullcycle/comunicao-entre-sistemas/graphql:/app fc-graphql
~~~


##### RUN BASH
~~~
docker run -it --rm -p 9090:9090 -v ~/fullcycle/comunicao-entre-sistemas/graphql:/app fc-graphql bash
~~~

##### GQL GENERATE
~~~
go run github.com/99designs/gqlgen generate
~~~

### Queries

~~~graphql
mutation createClass {
  createClass(
    input: {name: "Anfíbios", description: "Anfíbios são animais vertebrados com duas fases de vida (aquática e terrestre)."}
  ) {
    id
    name
    description
  }
}

query queryClasses {
  classes {
    id,
    name,
    description
  }
}

query queryClassesWithAnimals {
  classes {
    id,
    name,
    description,
    animals {
      id,
      name,
      description
    }
  }
}


mutation createAnimal {
  createAnimal(
    input: {
      name: "Perereca", 
      classId:"a3d2d411-8bab-4632-8eec-740897213b1a",
      description: "Perereca é uma designação genérica de anfíbios da ordem Anura da família Hylidae, Pelodryadidae ou Phyllomedusidae."
    }
  ) {
    id
    name
    description,
  }
}

query queryAnimals {
  animals {
    id,
    name,
    description,
  }
}

query queryAnimalsWithType {
  animals {
    id,
    name,
    description,
    type {
      id,
      name
    }
  }
}
~~~