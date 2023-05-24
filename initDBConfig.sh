#!/bin/bash

# Проверка наличия всех параметров
if ! [ $1 ] && ! [ $2 ] && ! [ $3 ] && ! [ $4 ] && ! [ $5 ] ; then

	echo "input DB_HOST, DB_PORT, DB_NAME, DB_USER and DP_PASSWORD"
	fi

# Запись параметров в переменные окружения
if [ $1 ] && [ $2 ] && [ $3 ] && [ $4 ] && [ $5 ]; then
	export	"DP_HOST=$1" 
	export	"DB_PORT=$2"
	export	"DB_NAME=$3"
	export	"DB_USER=$4"
	export	"DB_PASSWORD=$5"
	echo "Database config successfuly writed"
	fi