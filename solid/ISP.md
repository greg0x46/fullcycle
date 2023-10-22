### Interface Segregation Principle

O Interface Segregation Principle (ISP) ou Princípio da Segregação de Interface trata das desvantagens da implementação de interfaces “gordas”.

Ao ISP atribuímos a seguinte premissa:

>Crie interfaces granulares a específicas para os seus clientes.

Um outro mantra que simboliza muito bem a ideia do ISP é:

>Clientes não devem ser forçados a depender de interfaces que eles não usam.

Aqui “Clientes” não são os usuários finais do sistema, e sim as classes e algoritmos que dependem de uma interface dentro do sistema.

De forma mais clara, podemos dizer que o principio prega que uma interface não deve forçar uma classe a implementar coisas que ela não irá utilizar.

Interfaces que tem muitos comportamentos (Interfaces gordas) geralmente se espalham pelo sistema trazendo complexidade e dificuldade de manutenção ao código.

**Vamos ver um exemplo fácil de quebra desse princípio?**

~~~
public interface ITelefone{
  void Tocar();
  void Discar();
  void TirarFoto();
}

public class TelefoneCelular : ITelefone{
  public void Tocar() { ... }
  public void Discar() { ... }
  public void TiraFoto() { ... }
}

public class TelefoneComum : ITelefone{
  public void Tocar() { ... }
  public void Discar() { ... }
  public void TiraFoto() { 
    throw new NotImplementedException();
  }
}
~~~

Perceba que a classe TelefoneCelular implementou a interface corretamente e todos os métodos eram usuais a classe.

Já para a classe TelefoneComum tivemos um método que lançou uma Exception, pois aquele metódo não tinha utilidade para a classe.

Percebemos que criarmos uma Interface genérica e nada específica às nossas classes isso pode gerar complexidade e difícil manutenção posterior ao código.

**E como resolver e respeitar a premissa do principio?**

Simples, criando interfaces específicas para as classes.

~~~
public interface ITelefoneCelular
{
  void Tocar();
  void Discar();
  void TirarFoto();
  void Conectar3G();
}

public interface ITelefoneComum
{
  void Tocar();
  void Discar();
}

public class TelefoneCelular : ITelefoneCelular{
  public void Tocar() { ... }
  public void Discar() { ... }
  public void TiraFoto() { ... }
  public void Conectar3G() { ... }
}

public class TelefoneComum : ITelefoneComum{
  public void Tocar() { ... }
  public void Discar() { ... }
}
~~~

**Concluindo a analise da premissa do ISP**

A premissa do ISP nos alerta quanto a dependência e utilização de interfaces “gordas” que forçam os clientes a implementar métodos desnecessários.

Respeitando a premissa do ISP geramos facilidade de manutenção, pois temos especificidade nas classes clientes, quebramos o acoplamento entre as classes que a implementação de interfaces “gordas” traz e ainda ganhamos coesão e eficiência no código.

---

Fonte: [Solid — L.S.P — Liskov Substitution Principle: Thiago Arahão](https://medium.com/@tbaragao/solid-l-s-p-liskov-substitution-principle-3a31c3a7b49e)