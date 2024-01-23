# CVWO Assignment Golang App
1. Fork or clone the repo
2. Open the project directory in any code editor
3. Install Golang if not done so already at the following link: https://go.dev/doc/install
4. Install MySQL by running "sudo apt install MySQL-server" on the terminal
5. Check if MySQL is installed successfully by running "MySQL --version" to see the version number
6. Start MySQL server by running "sudo /etc/init.d/MySQL start"
7. To start using MySQL, run "sudo MySQL"
8. To create a MySQL database called “internetforum”, run "create database internetforum;" at the MySQL command line (NOTE ALL COMMANDS BELOW until instruction 12 IS RUN ON THE MYSQL COMMAND LINE)
9. Then, to use the created database, run "use internetforum;"
10. Then, create a user called “user1” with password “Pas$w0rd” by running "create user 'user1'@'localhost' identified by 'Pas$w0rd';"
11. Next to grant all privileges so that user1 can use the internetforum database, run the command "GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO 'user1'@'localhost' WITH GRANT OPTION;"
12. Next, exit mysql to log in as user1 by running "exit;"
13. Login to user1 by running on the terminal "sudo mysql -u user1 -p" and type password 'Pas$w0rd'
14. Use the created database by running "use internetforum;" on the MySQL command line.
15. Run all the MySQL commands inside cvwo-backend/database/db.sql by copying and pasting all 5 commmands one by one to create all 5 necessary tables in the MySQL database "internetforum" or just simply run the command "source cvwo-backend/database/db.sql;" to create all 5 tables at ones.
16. Next, after all 5 tables are created, type the following command in the terminal (make sure that the current directory is the project directory): “go run cmd/server/main.go".
17. Backend and database is ready!
