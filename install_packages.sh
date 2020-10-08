#!/bin/bash

apt remove docker docker-engine docker.io -y
apt update -y
apt install apt-transport-https ca-certificates  curl  software-properties-common -y

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

apt update -y
apt install docker-ce=5:19.03.13~3-0~ubuntu-bionic -y

apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF

apt-get update -y
#apt install kubernetes-cni=0.6.0-00 --allow-downgrades -y;
#apt-get install -y kubelet=1.12.5-00 kubeadm=1.12.5-00 kubectl=1.12.5-00 --allow-downgrades
sudo apt-mark unhold kubeadm && apt update && apt install -y kubeadm=1.18.9-00 --allow-downgrades && apt-mark hold kubeadm
sudo apt-mark unhold kubectl && apt update && apt install -y kubectl=1.18.9-00 --allow-downgrades && apt-mark hold kubectl
sudo apt-mark unhold kubelet && apt update && apt install -y kubelet=1.18.9-00 --allow-downgrades && apt-mark hold kubelet
