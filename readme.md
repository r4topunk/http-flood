# HTTP Flood

> **Nota:** Este repositório foi criado para fins educacionais. O autor não se responsabiliza pelo uso das ferramentas disponibilizadas aqui.

## O que é?

Um ataque de inundação HTTP é um tipo de ataque volumétrico de negação de serviço distribuída (DDoS) projetado para sobrecarregar um servidor alvo com solicitações HTTP. Quando o servidor fica saturado e incapaz de responder ao tráfego normal, ocorre uma negação de serviço para as solicitações adicionais de usuários legítimos.

## Como ocorre?

Os ataques de inundação HTTP são um tipo de ataque DDoS de "camada 7". A camada 7 é a camada de aplicação do modelo OSI e se refere a protocolos da internet, como o HTTP. O HTTP é a base das solicitações feitas pelos navegadores para carregar páginas web ou enviar conteúdos de formulários. Mitigar ataques na camada de aplicação é especialmente desafiador, pois o tráfego malicioso é difícil de distinguir do tráfego legítimo.

Para maximizar a eficiência, atores maliciosos geralmente utilizam botnets para amplificar o impacto do ataque. Com vários dispositivos infectados por malware, um invasor pode lançar um volume maior de tráfego de ataque.

### Tipos de ataques de inundação HTTP:

- **Ataque HTTP GET**: Vários computadores ou dispositivos são coordenados para enviar múltiplas solicitações de imagens, arquivos ou outros ativos de um servidor alvo. Quando o servidor fica inundado com solicitações e respostas, ocorre a negação de serviço para fontes de tráfego legítimas.
- **Ataque HTTP POST**: Quando um formulário é enviado em um site, o servidor deve processar a solicitação e enviar os dados para uma camada de persistência, geralmente um banco de dados. O processamento dos dados e a execução dos comandos do banco de dados consomem mais recursos do que o envio da solicitação POST. Este ataque explora essa disparidade, enviando múltiplas solicitações POST até que a capacidade do servidor seja saturada, resultando em negação de serviço.

## Desenvolvimento

### Parte 1

Como primeira etapa dos estudos, organizamos o ambiente de testes. Enquanto eu criava um script para explorar a vulnerabilidade, o [S4mplin](https://github.com/jacksoncastilho) configurou um servidor em Docker para termos um ambiente controlado de testes.

Para o script, iniciamos com dois modelos:

1. **Modelo em Go**: Escolhemos Go para explorar a linguagem e ter um primeiro contato. Como Go é popular na área de segurança da informação, acreditamos ser interessante aprender mais sobre ela.
2. **Modelo em JavaScript**: Diferente do modelo em Go, este script é executado no navegador de um usuário que acessa um site infectado, transformando o dispositivo do usuário em uma ferramenta de requisição para o alvo.

Para o servidor, configuramos um serviço Apache em um container. Este container serve como um ambiente controlado de testes, onde mediremos o comportamento do servidor sob ataque.

Nosso objetivo é entender como podemos melhorar ambos os lados do desenvolvimento, aprimorando a proteção do servidor à medida que o script se torna mais sofisticado.
