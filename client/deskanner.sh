#!/bin/bash

RANGE_SERVER_IP=${1}
RANGE_SIZE=¢{2}
RANGE_SERVER_CA_DIR=${3}
STORAGE_SERVER_URL=${4}

if [ ${#} -ne 4 ]; then
    echo -e "usage: ./theskanner <range-server-ip> <range-size> <range-server-ca-dir> <storage-server-url>"
    exit 1
fi

echo -e "[+] getting range from the server..."
RANGE=`wget --ca-certificate= https://$SERVER_IP/range | egrep -oE "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\-(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\-(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)"`

while true; do
    read -p "[+] can i scan for ipsec ports on range ${RANGE}?" yn
    case $yn in
        [Yy]* ) XML=`sudo nmap -oX - --open -T 5 -sU -p 500,4500 ${RANGE} | base64` && wget https://${STORAGE_SERVER_URL} --post-data=${XML} && break;;
        [Nn]* ) echo -e "[+] thanks anyway dude!" && exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

echo -e "[+] exiting..."

exit 0