#!/bin/bash

host=$1

sudo swapoff -a
if [[ -z $host ]]; then
    sudo kubeadm init --kubernetes-version 1.15 --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all
else
    sudo kubeadm init --kubernetes-version 1.15 --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all --apiserver-advertise-address=$host
fi
mkdir -p $HOME/.kube
sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
sudo sysctl net.bridge.bridge-nf-call-iptables=1
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml
kubectl taint nodes --all node-role.kubernetes.io/master-

sudo kubeadm token create --print-join-command

