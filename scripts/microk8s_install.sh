#! /bin/bash
apt-get update
apt-get install -yqq tmux
snap install microk8s --classic
microk8s.enable dns
