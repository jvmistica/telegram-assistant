# telegram-assistant
![Build status](https://github.com/jvmistica/telegram-assistant/workflows/build/badge.svg)
![Coverage](https://img.shields.io/sonar/coverage/jvmistica_telegram-assistant/main?server=https%3A%2F%2Fsonarcloud.io)

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
4. Run the server with `make run`

## Commands
```
/start                                - List the available commands  
/listitems                            - List items in your inventory  
/listitems sort by <field> asc        - List items in your inventory in ascending order  
/listitems sort by <field> desc       - List items in your inventory in descending order  
/listitems filter by <field> = <text> - List filtered items in your inventory (operations can be =, >, <, >=, <=, <>)  
/showitem                             - Show prompt for entering the item to be shown  
/showitem <item>                      - Show an item's details  
/additem                              - Show prompt for entering the item to be added  
/additem <item>                       - Add an item to your inventory  
/updateitem                           - Show prompt for entering the item to be updated  
/updateitem <item> <field> <value>    - Update an item in your inventory  
/deleteitem                           - Show prompt for entering the item to be deleted  
/deleteitem <item>                    - Delete an item in your inventory    
/importitems                          - Import records from a CSV file 
```


## Usage
Adding the bot to a chat for the first time will present you with a START button. Upon clicking it, the Menu button will become available and clicking it will show a list of commands accepted by the bot.
![image](https://github.com/jvmistica/telegram-assistant/assets/53989745/82e36532-5c2e-4f89-84d6-df43a6889b42)  

Listing and adding items  
![image](https://github.com/jvmistica/telegram-assistant/assets/53989745/e40d1166-ac5a-4706-959e-f8cfcb299dde)  

Updating and showing items  
![image](https://github.com/jvmistica/telegram-assistant/assets/53989745/a7c0dc70-cffb-453b-b694-e3ca7f71e3c5)
![image](https://github.com/jvmistica/telegram-assistant/assets/53989745/227a929f-4860-41a7-b8ba-724d8943a173)  

Deleting items  
![image](https://github.com/jvmistica/telegram-assistant/assets/53989745/1d87265e-bd5f-42cc-84b0-58f0f219fbc0) 
