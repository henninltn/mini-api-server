#!/bin/sh

# Usage: xxxxxx.sh url times json
#
# Arguments:
#   url         
#   times       
#   json        it MUST be surrounded single-quartations
#
# Example
#   sh xxxxxx.sh http://localhost:8080/webapi/user/XXXXX 10 '{"token":"xxxxx","version":"1.3"}'
#

URL=$1
TIMES=$2
JSON=$3

if [ -z "$URL" ]; then
    echo "no url"
    exit;
fi
if [ -z "$TIMES" ]; then
    echo "no times"
    exit;
fi
if [ -z "$JSON" ]; then
    echo "no json"
    exit;
fi

n=1
until [ $n -gt $TIMES ]
do
    curl -H 'Content-Type:application/json' -d "$JSON" $URL
    n=$(( n+1 ))
    echo ""
    sleep 1s
done
