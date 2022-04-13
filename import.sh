#!/bin/bash
i=1;
MAX_INSERT_ROW_COUNT=$2;
HSET_KEY=$1;
while [ $i -le $MAX_INSERT_ROW_COUNT ]
do
  echo $i
  redis-cli -h localhost -p 6379 hset "$HSET_KEY" $(($i)) $(($i))
  i=$(($i+1))
done

exit 0