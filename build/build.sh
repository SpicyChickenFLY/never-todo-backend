#!/bin/bash

# initialize database
mysql -uroot -p -D never_todo < todo.sql

# move linux service conf
cp ./never-todo-backend.service

# compile go program
go build ../main.go

# move executable file to /usr/bin
mv 

# move static file to /usr/bin
mv 