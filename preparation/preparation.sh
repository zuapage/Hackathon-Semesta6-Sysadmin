#This script is for prepartion on first step on this lab
#By running this script you will install, start, enable docker service also net-tools
#Docker will used for setup ansible that will use later.

echo  "----- Preparation | Install Docker & Net Tools -----"
sudo apt install ca-certificates curl gnupg lsb-release -y
sudo mkdir -m 0755 -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
sudo apt install net-tools -y
sudo systemctl start docker && sudo systemctl enable docker

echo "-------------- Current status of docker -------------------"
STATUS_DOCKER=$(systemctl status docker)
echo "$STATUS_DOCKER"

# This network will use by (ansible, jenkins & grafana container)
echo "-------------- Create new docker network (dev_semesta) -------------------"
sudo docker network create --subnet=192.168.1.0/29 --gateway=192.168.1.1 dev_semesta
sudo docker network ls | grep dev_semesta
