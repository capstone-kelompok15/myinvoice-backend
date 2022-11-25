#!/bin/bash

database_username=$1
database_password=$2
database_host=$3
database_port=$4
database_name=$5

migrate -source file://db/migrations -database "mysql://$database_username:$database_password@tcp($database_host:$database_port)/$database_name" up