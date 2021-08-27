# test-linkaja

## Installation Step

- Extract test-linkaja.zip into your directory
- Copy .env.example into .env
- Configure .env depend on your DB Configuration

DB_HOST=<YOUR_HOST>  
DB_PORT=<YOUR_PORT>  
DB_USER=<YOUR_USER_DB>  
DB_PASS=<YOUR_PASSWORD_DB>  
DB_NAME=<YOUR_DB_NAME>  

** Don't forget to create DB based on your configuration **  

- Run the application using  
**go run server.go**  
When the application is running, the application will migrate the database from the struct, creating table **Account** and **Customer** and automatically seed table with default data   

