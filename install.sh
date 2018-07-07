#!/bin/sh

clear

if [ "$EUID" -ne 0 ]
  then echo "Please run install script as root!"
  exit
fi

err() {
  clear
  echo Error connecting to download server! Check your internet connection.
  exit
}

uninst() {
  rm /usr/bin/cgsf-encode
  rm /usr/bin/cgsf-decode
  rm /home/"$USER"/share/CGSF
  clear
  echo Uninstalled successfully!
}

inst() {
  mkdir /home/"$USER"/share/CGSF
  cd /home/"$USER"/share/CGSF
  wget https://gitlab.com/mbednarek360/CGSF/-/archive/master/CGSF-master.zip || err
  unzip CGSF-master.zip
  rm CGSF-master.zip
  mv CGSF-master/encode .
  mv CGSF-master/decode .
  mv CGSF-master/install.sh .
  ln -s encode /usr/bin/cgsf-encode
  ln -s decode /usr/bin/cgsf-decode
  rm -r CGSF-master
  clear
  echo Installed successfully.
}

if [ -d ~/.local/share/CGSF ]; then
  while true; do
      read -p "CGSF is already installed on this system. Would you like to uninstall?" yn
      case $yn in
          [Yy]* ) uninst; exit;;
          [Nn]* ) exit;;
          * ) echo "Please answer yes or no.";;
      esac
  done
fi

if [ ! -d /home/"$USER"/share/CGSF ]; then
  while true; do
    read -p "Are you sure you want to install this program?" yn
    case $yn in
        [Yy]* ) inst; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
  done
fi