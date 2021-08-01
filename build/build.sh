#!/bin/bash

export GIN_MODE=release

# compile go program
go build ../main.go -o never-backend

# move executable file to /usr/bin
mv never-backend /usr/bin

# move static file to /usr/bin
mv ../config/never-backend.ini /etc/

# initialize database
never-backend init
# mysql -uroot -p -D never_todo < todo.sql

# move linux service conf
cp ./never-todo-backend.service