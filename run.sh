#! /bin/bash -

if [ -z $PORT ];then
  PORT=8080
fi

if [ -z $HOST ];then
  HOST=0.0.0.0
fi

./main -port=$PORT -host=$HOST