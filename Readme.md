## How to run

- Clone the repository to your computer.
- Navigate to the root folder of the project. i.e currency-exchange-app
- Run the following command. 

```bash
docker-compose up -d
```

- This starts the server (port 3001), client (port 3000) and database(port 3306) application simultaneoulsy.

- If the Golang app throws an error, restart the MySQL container and Golang app container respectively.

Some of the pages in the Nuxt client side application are:
/login
/signup
/transact
/transfer
/profile
/wallet