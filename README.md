
## gofast 🏎️
Um template básico para projetos em Go, focado em rápida inicialização e desenvolvimento.

## Recursos

- **Chi Router**: Utiliza o Chi como roteador HTTP para gerenciar as rotas da sua aplicação.
- **Arquivo de Roteamento**: Inclui um arquivo de roteamento para facilitar a organização e manutenção das rotas.
- **Validação do Corpo das APIs**: Implementa validação do corpo das requisições para garantir a integridade dos dados recebidos.
- **Leitura de Variáveis de Ambiente**: Oferece suporte para carregar configurações e segredos através de variáveis de ambiente.

## Uso

1. **Clonar o Repositório**:

   ```bash
   git clone https://github.com/seu-usuario/gofast.git
   cd gofast
   ``` 

2.  **Configurar as Variáveis de Ambiente**:
    
    Crie um arquivo `.config.yaml` na raiz do projeto e defina as variáveis de ambiente necessárias.
    
3.  **Executar o Projeto**:
    
    
     
    ```bash
    make dev
    ``` 
    

## Comandos do Makefile

Para facilitar o desenvolvimento e a execução do projeto, você pode usar os seguintes comandos definidos no Makefile:

-   **dev**: Compila e executa o projeto.
-   **test**: Executa os testes do projeto.
-   **dist**: Limpa os artefatos de compilação e binários gerados.
-  **build**: Builda o projeto

Para executar um comando, basta digitar `make` seguido do nome do comando desejado no terminal.

## Estrutura do Projeto
 

```bash
├── cmd
├── pkg
│   ├── bootstrap
│   ├── config
│   ├── ctx
│   ├── database
│   └── server
│       ├── api
│       │   └── health
│       ├── router
│       └── validator
├── Makefile
└── README.md
``` 

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.
