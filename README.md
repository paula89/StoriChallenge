# StoriChallenge

This project try to resolve the challenge in the challenge.pdf

Pre requires: 

- Install Docker
- Install some IDE to see the Golang code
- Install some IDE to see the mySql db


How to run : 

-   Open a terminal pointing to the StoriChallenge folder
-   run ´docker-compose up´
-   Visualize the logs in the terminal
-   Visualize the email in the email folder


Logs

-   The application start to run but don't finish, so you can continue seeing the logs and consulting the db
-   In the logs you can see all the successfull workflow
-   How it's running continually, you can start to visualize the idempotence in the db 
Example: "2024-04-23 00:09:29 2024/04/23 03:09:29 Error saving transactions in the db: Error 1062 (23000): Duplicate entry '8ef2aac5-1116-44ae-961a-03273a682dc6' for key 'Transactions.PRIMARY'"


Db
- you can connect to the db and do a select query to review that the table was filled correctly
- Ex: 
  -   Host : localhost
  -   Database: StoriChallenge
  -   Username: root
  -   Password: Stori2024!!
  -   Port: 3306
  -   Query:  select * from Transactions
