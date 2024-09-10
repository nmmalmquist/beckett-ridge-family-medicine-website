to deploy to EC2 use the command GOOS=linux GOARCH=amd64 go build

use the command scp -i ~/.ssh/brfm-server-key.pem ./bin/linux/amd64/app ec2-user@18.117.23.138:~/ to copy the file over

run sudo systemctl restart brfm on ec2 to run the new version.
