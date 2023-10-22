### Open Closed Principle

O Open Closed Principle ou Principio do aberto/fechado (OCP) é também um dos mais importantes da OOP (princípios da orientação a objetos) e certamente um dos mais complicados para se implementar, por isso dificilmente programadores o praticam.

O principio tem como objetivo trabalhar diretamente com a manutenção das classes e tem como premissa a seguinte afirmativa:

>Você deve ser capaz de estender um comportamento de uma classe sem a necessidade de modificá-lo.

Ou seja, as entidades de software como classes, módulos, funções, etc, devem estar abertas para extensão, porém fechadas para modificação.

Em outras palavras significa que esta classe pode ter seu comportamento alterado com facilidade quando necessário, sem a alteração do seu código fonte. Essa extensão pode ser feita através de herança, interface e composição.

**Vamos ver um exemplo?**

~~~
public enum TipoEmail {
	Texto,
	Html,
	Criptografado
}

public void class EnviarEmail(string mensagem, string assunto, TipoEmail tipo){
	if(tipo == TipoEmail.Texto)
	{
		mensagem = this.RemoverFormatacao(mensagem);
	}
	else if(tipo == TipoEmail.Html)
	{
		mensagem = this.InserirHtml(mensagem); 
	} 
	else if(tipo == TipoEmail.Criptografado)
	{
		mensagem = this.CriptografarMensagem(mensagem);
	}

	this.EnviarMensagem();
}
~~~

Acima temos o exemplo de uma classe que valida o tipo de e-mail para tratar a mensagem de acordo com o seu tipo, mas, quando uma nova forma de mensagem for criada, a classe deverá ser editada e um novo if deverá ser acrescentado a ela.

**Bem simples a solução não? Mas, com isso não quebraríamos o princípio do OCP (Open closed Principle)?**

O conceito de aberto/fechado (OCP) tem a premissa de criarmos novas classes para funcionalidades de tipos semelhantes, e caso tenhamos novas funcionalidades, novas classes sejam criadas.

**Legal, e o que ganhamos aplicando o OCP?**

##### Extensibilidade

Como citei, ao termos uma nova funcionalidade ou comportamento, não precisaremos alterar a classe já existente, e sim estendê-la. Com isso mantemos o código original confiável e intacto, e criamos um código com design duradouro, de qualidade e manutenibilidade altas.

##### Abstração

Com toda certeza já ouvimos sobre abstração, ela permite que toda a mágica aconteça. Se temos abstrações bem definidas, conseguimos de forma fácil estender os métodos da nossa aplicação.
O conceito OCP indica principalmente o uso da herança para praticarmos o extensão dos métodos.

**Vamos ver como a classe deveria ficar?**

~~~
public abstract class Email
{
	public abstract void Enviar(string assunto, string mensagem);
}

public class TextoEmail : Email
{
	public override void Enviar(string assunto, string mensagem)
	{
		Util.RemoverFormatacao(mensagem);
	}
}

public class HtmlEmail : Email
{
	public override void Enviar(string assunto, string mensagem)
	{
  		Util.InserirHtml(mensagem);
	}
}

public class CriptografadoEmail : Email
{
	public override void Enviar(string assunto, string mensagem)
	{
		Util.CriptografarMensagem(mensagem);;
	}
}
~~~

Veja que na solução, criamos várias classes, cada uma com uma responsabilidade definida, suas próprias regras de negócios e sem a necessidade de alterarmos a funcionalidade padrão devido a criação de uma nova regra.

Este design é a base de muitos dos padrões de projetos, um exemplo é o Factory Method, onde as classes de negócio estão implementadas sem uma especificação da classe concreta.

---

Fonte: [SOLID — OCP — Open Closed Principle: Thiago Arahão](https://medium.com/@tbaragao/solid-ocp-open-closed-principle-600be0382244)