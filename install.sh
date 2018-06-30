#!/bin/sh

if [ "$EUID" -ne 0 ]
  then echo "Please run install script as root!"
  exit
fi

mkdir ~/.local/share/CGSF
cd ~/.local/share/CGSF
wget https://gitlab.com/mbednarek360/CGSF/-/archive/master/CGSF-master.zip
unzip CGSF-master.zip
rm CGSF-master.zip
mv CGSF-master/encode .
mv CGSF-master/decode .
chmod +x CGSF-master/cgsf-encode
chmod +x CGSF-master/cgsf-decode
mv CGSF-master/cgsf-encode /usr/bin
mv CGSF-master/cgsf-decode /usr/bin
rm -r CGSF-master
clear
echo Installed sucecssfully.