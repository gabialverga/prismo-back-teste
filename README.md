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
    
    - **MYSQL_USER**: Nome de utilizador da base de dados.
    - **MYSQL_PASSWORD**: Senha do usuário da base de dados.
    - **MYSQL_DATABASE**: Nome da base de dados.
    - **MYSQL_ROOT_PASSWORD**: Senha da base de dados.

2.  Build

    ```
    make build
    ```

3.  Run

    ```
    make up
    ```
    
### Comandos extras

- **Iniciar apenas um serviço**

    Pode receber os valores de 'db' ou 'back'
  ```
  make up s=<nome>
  ```

- **Ver logs dos containers** 

    Pode receber os valores de 'db' ou 'back' ou ficar vazio para retornar os logs de todos os containers
  ```
  make logs s=<nome>
  ```

- **Parar um container**

    Pode receber os valores de 'db' ou 'back' ou ficar vazio para parar todos os containers
  ```
  make down s=<nome>
  ```

- **Acessar um container**

    Pode receber os valores de 'db' ou 'back'
  ```
  make exec s=<nome>
  ```