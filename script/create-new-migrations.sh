#!/bin/bash

migration_name = $1

migrate create -ext sql -dir db/migrations $migration_name