# Prismo back teste

### Dependencias

- [Docker](https://docs.docker.com/desktop/install/linux-install/)
- [Docker Compose](https://docs.docker.com/compose/install/linux/)

### Deploy usando Docker

Aqui esta o passo a passo de como fazer o deploy da aplicação:

1.  Setup environment

    Faça uma cópia do arquivo env de exemplo e faça as mudanças das variáveis:

    ```
    cp .env.example .env
    ```

    Variaveis:
    
    - **MYSQL_USER**: Nome de utilizador da base de dados. Deve ser o mesmo utilizado no ficheiro backend init.sql.
    - **MYSQL_PASSWORD**: Palavra-passe da base de dados.
    - **MYSQL_DATABASE**: Nome da base de dados.

2.  Build

    ```
    make build
    ```

3.  Run

    ```
    make up
    ```
    
### Comandos extras

- Iniciar apenas um serviço
    Pode receber os valores de 'db' ou 'backend'

  ```
  make up s=<nome>
  ```

- Ver logs
    Caso queira ver os logs de um container em específico adicionando `s=<nome>`, nome pode ser 'db' ou 'backend'

  ```
  make logs
  ```

- Stop
    Caso queira ver os logs de um container em específico adicionando `s=<nome>`, nome pode ser 'db' ou 'backend'

  ```
  make down
  ```
