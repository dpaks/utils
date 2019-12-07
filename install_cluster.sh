#!/bin/bash

host=$1

sudo swapoff -a
if [[ -z $host ]]; then
    sudo kubeadm init --kubernetes-version 1.15.6 --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all
else
    sudo kubeadm init --kubernetes-version 1.15.6 --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all --apiserver-cert-extra-sans=$host
fi
mkdir -p $HOME/.kube
sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
sudo sysctl net.bridge.bridge-nf-call-iptables=1
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/62e44c867a2846fefb68bd5f178daf4da3095ccb/Documentation/kube-flannel.yml
kubectl taint nodes --all node-role.kubernetes.io/master-

sudo kubeadm token create --print-join-command

