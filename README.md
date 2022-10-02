# Henchmaid
![Build status](https://github.com/jvmistica/henchmaid/workflows/Go/badge.svg)  
A telegram bot that allows users to manage their inventory - items and recipes. Initially made to prevent food stored in the far reaches of my kitchen cabinet and fridge from going to waste.

## Running the Server Locally
1. Install and run PostgreSQL
2. Create a user, password, and database
3. Set the values for the environment variables or simply enter the commands below in the terminal:

```
export POSTGRES_HOST=<postgres_host>
export POSTGRES_PORT=<postgres_port>
export POSTGRES_USER=<user>
export POSTGRES_PASS=<password>
export POSTGRES_DB=<postgres_db>
export BOT_TOKEN=<telegram-bot-token>
```
4. Run then server with `make run`

## Commands
```
/start                                - List the available commands  
/listitems                            - List items in your inventory  
/listitems sort by <field> asc        - List items in your inventory in ascending order  
/listitems sort by <field> desc       - List items in your inventory in descending order  
/listitems filter by <field> = <text> - List filtered items in your inventory (operations can be =, >, <, >=, <=, <>)  
/showitem <item>                      - Show an item's details  
/additem <item>                       - Add an item to your inventory  
/updateitem <item> <field> <value>    - Update an item in your inventory  
/deleteitem <item>                    - Delete an item in your inventory  
```
