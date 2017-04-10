ufw allow OpenSSH
ufw --force enable
# turns on firewall
#replace your-dockerhug-name with your docker hub name
#and your-image-name with your image name and
#uncomment the line by removing the leading #
# --name defines the name, so docker doesn't give us weird name
#docker run --name zipsvr -d -p 80:80 weijen0330/zipsvr
