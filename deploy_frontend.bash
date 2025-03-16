#!/bin/bash
FRONTEND_USER=$1
FRONTEND_DOMAIN=$2

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <frontend_user> <frontend_domain>"
    exit 1
fi

tar -cf frontend.tar frontend
xz -9 frontend.tar
scp frontend.tar.xz $FRONTEND_USER@ftp.$FRONTEND_DOMAIN:/home/$FRONTEND_USER/
ssh $FRONTEND_USER@ftp.$FRONTEND_DOMAIN 
