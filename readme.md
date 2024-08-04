# HTTP Flood

## O que é?

Um ataque de inundação HTTP é um tipo de ataque volumétrico de negação de serviço distribuída (DDoS) desenvolvido para sobrecarregar um servidor visado com solicitações HTTP. Depois que o alvo ficar saturado com solicitações e não conseguir responder ao tráfego normal, ocorrerá a negação de serviço para as solicitações adicionais de usuários verdadeiros.

## Como ocorre?

Os ataques de inundação HTTP são um tipo de ataque DDoS de "camada 7". A camada 7 é a camada de aplicação do modelo OSI e se refere a protocolos da internet, tais como HTTP. O HTTP é a base das solicitações da internet feitas pelo navegador e é comumente usado para carregar páginas web ou para enviar conteúdos de formulários pela internet. Mitigar ataques na camada de aplicação é especialmente complexo, pois o tráfego malicioso é difícil de distinguir do tráfego normal.

A fim de alcançar a máxima eficiência, os atores maliciosos geralmente empregarão ou criarão botnets para maximizar o impacto de seu ataque. Utilizando vários dispositivos infectados com um malware, um invasor é capaz de potencializar seus esforços, lançando um volume maior de tráfego de ataque.

### Há duas variedades de ataques de inundação HTTP:

- **Ataque HTTP GET** - nessa forma de ataque, vários computadores ou outros dispositivos são coordenados para enviar várias solicitações de imagens, arquivos ou de algum outro ativo de um servidor visado. Quando o alvo é inundado com solicitações e respostas recebidas, ocorre a negação de serviço para solicitações adicionais de fontes de tráfego legítimas.
- **Ataque HTTP POST** - normalmente, quando um formulário em um site é enviado, o servidor deve tratar a solicitação recebida e enviar os dados para uma camada de persistência, na maioria das vezes um banco de dados. O processo de tratamento dos dados do formulário e execução dos comandos necessários do banco de dados é relativamente intensivo em comparação com a quantidade de capacidade de processamento e largura de banda necessária para enviar a solicitação POST. Esse ataque utiliza a disparidade no consumo relativo de recursos, enviando muitas solicitações post diretamente para um servidor visado até que sua capacidade esteja saturada e ocorra a negação de serviço.

## Como previnir?

## Ideias

- Criar ambiente de sevidor para receber o ataque
  - O servidor pode estar em um ambiente docker

## Referências

- https://github.com/Leeon123/CC-attack
-
