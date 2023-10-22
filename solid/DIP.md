### Dependency Inversion Principle

O Dependency Inversion Principle (DIP) ou Princípio da Inversão de Dependência é a base para termos um projeto com um excelente design orientado a objetos, focado no domínio e com um arquitetura flexível.

O DIP é regido pela seguinte premissa:

> Dependa de abstrações e não de implementações.

De uma forma objetiva o princípio nos faz entender que sempre devemos depender de abstrações e não das implementações, afinal de contas, as abstrações mudam menos e facilitam a mudança de comportamento e as evoluções do código.

**Vamos observar um exemplo de violação desse princípio?**

~~~
public class Interruptor
{
  private Ventilador _ventilador;
  
  public void Acionar()
  {
    if(!_ventilador.Ligado)
      _ventilador.Ligar();
    else
      _ventilador.Desligar();
  }
}

public class Ventilador
{  
  public bool Ligado {get; set; }
  
  public void Ligar() { ... }
  
  public void Desligar() { ... }
}
~~~

No exemplo, podemos perceber que além de quebrar outros princípios do SOLID, a classe concreta Interruptor depende de uma outra classe concreta (Ventilador).

O Interruptor deveria ser capaz de acionar qualquer dispositivo independente de ser um ventilador uma lâmpada ou até mesmo um carro.

**Vamos corrigir o exemplo aplicando o DIP?**

~~~
interface IDispositivo
{
  bool Ligado { get; set; }
  void Acionar();
  void Ligar();
  void Desligar();
}

public class Ventilador : IDispositivo
{  
  public bool Ligado { get; set; }
  
  public void Acionar ()
  {
    if (!this.Ligado)
      this.Ligar();
    else
      this.Desligar();
  }
  
  public void Ligar() { ... }
  
  public void Desligar() { ... }
}

public class Lampada : IDispositivo
{  
  public bool Ligado { get; set; }
  
  public void Acionar ()
  {
    if (!this.Ligado)
      this.Ligar();
    else
      this.Desligar();
  }
  
  public void Ligar() { ... }
  
  public void Desligar() { ... }
}

public class Interruptor
{
  private readonly IDispositivo _dispositivo;
  
  public void AcionarDispositivo()
  {
    _dispositivo.Acionar();
  }
}
~~~

Percebam que agora a classe concreta Interruptor depende da abstração de um IDispositivo e não mais de uma classe concreta.

**Concluindo a analise…**

O DIP trás uma série de benefícios, principalmente em relação a arquitetura de software.

O principio torna o aplicação focada na resolução dos problemas, fazendo da implementação um mero detalhe.

Tendo como base a afirmativa acima, podemos perceber que a abstração IDispositivo está diretamente vinculado com o cliente (Interruptor), tornando sua implementação (Ventilador, Lampada) um detalhe. : )

Identificar as abstrações é importante para que mantenhamos o projeto flexível, robusto e preparado para que as futuras implementações não sejam difíceis e complexas.

---

Fonte: [Solid — D.I.P — Dependency Inversion Principle: Thiago Arahão](https://medium.com/@tbaragao/solid-d-i-p-dependency-inversion-principle-e87527f8d0be)