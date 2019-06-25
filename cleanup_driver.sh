#!/bin/bash

sudo find /lib/modules -name nvidia.ko -delete 
sudo depmod -a
sudo rm -f /etc/docker/daemon.json 
sudo systemctl restart docker
sudo apt remove nvidia-docker2 -y
sudo nvidia-uninstall -s
kubectl delete ds -n kube-system nvidia-device-plugin-daemonset
sudo rm -rf /usr/lib64/nvidia /home/kubernetes/bin/nvidia
echo "nvidia-bug-report.sh nvidia-cuda-mps-control nvidia-cuda-mps-server nvidia-debugdump nvidia-installer nvidia-modprobe nvidia-persistenced nvidia-settings nvidia-smi nvidia-uninstall nvidia-xconfig" | xargs  -n1 | xargs -I{} sudo rm -f /usr/bin/{}
sudo apt autoremove -y

printf "\nRestart now...\n"
