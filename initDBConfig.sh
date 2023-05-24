#!/bin/bash

if [ $1 ] ; then
	export	"DB_URL=$1" 
	echo "database env has been imported"
	fi