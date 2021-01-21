# Descrição

Projeto de microservico baseado em vendas. Adaptação da video aula [Resiliência e regras de retry em Microsserviços](https://www.youtube.com/watch?v=PsNiFT3E0u8). Muito boa, assistam

# Arquitetura

* Order Service - Recebe os dados do usuário e envia para o microserviço Payment Service para processar o pagamento da order e responder

* Payment Service - Recebe uma request de order e processa de forma fake ela e retorna para o Order Service o resultado

