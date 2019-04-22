
Skip to content

    Why GitHub?
                          


                    
Enterprise
Explore
                      

                    
Marketplace
Pricing
                       


                        

Sign in
Sign up

1
17

    0

trigun117/TelegramBot-Go
Code
Issues 1
Pull requests 0
Projects 0
Insights
Join GitHub today

GitHub is home to over 31 million developers working together to host and review code, manage projects, and build software together.
TelegramBot-Go/installdocker.sh
@trigun117 trigun117 Added script for docker installing 2e97392 on 9 Mar 2018
26 lines (24 sloc) 1.61 KB
#!/bin/sh

sudo sh -c 'apt-get update
apt-get upgrade -y
apt-get dist-upgrade -y'

echo =========================================================================================================================================================================================
echo ==========START INSTALLING DOCKER====================START INSTALLING DOCKER====================START INSTALLING DOCKER==================================================================
echo =========================================================================================================================================================================================
sudo sh -c 'apt-get update
apt-get -y install \
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg2 \
	software-properties-common
curl -fsSL https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg |  apt-key add -
add-apt-repository \
	"deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
	$(lsb_release -cs) \
	edge"
apt-get update
apt-get install -y docker-ce'
echo =========================================================================================================================================================================================
echo ==========THE DOCKER ARE INSTALLED==============================THE DOCKER ARE INSTALLED==============================THE DOCKER ARE INSTALLED===========================================
echo =========================================================================================================================================================================================
