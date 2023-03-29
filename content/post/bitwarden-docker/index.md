+++
author = 'Nicholas Bennett'
date = '2022-03-21T02:58:03Z'
description = ''
draft = false
series = 'Hosting Bitwarden'
slug = 'bitwarden-docker'
tags = ['featured']
title = 'Bitwarden Docker'
+++


In this final post in the Hosting Bitwarden Series we will discuss how to self host bitwarden on any docker compatible OS. I will go over both the Docker-Compose route as well as a simple one liner docker run command to get it up and going with persistent storage and as a bonus I will share my backup solution that generates a backup of not only the database but the entire persistent directory which inlcudes the attachments. 

Before we get started let me preface this the assumption that you already have docker installed on your operating system of choice so I will not go over setup or installation of docker itself. 

So lets get started! In your favorite terminal (Mine is Iterm2 for mac) get connected to your Docker Host and nagivate to your docker directory (For me since I deployed my Vaultwarden Instance on my Synology this will be inside of my /volume1 in a folder called docker) 

Create a directory inside of your docker directory called Vaultwarden with another folder inside of that called bw-data or vw-data (this will be your persistent volume mount point to make backup and access to the containers filesystem easier) once you have this folder created we are ready to run the docker command to get Vaultwarden going

In that same terminal window run 


docker run -d --name vaultwarden -v /path-to-vw-data-here/:/data/ -p 80:80 vaultwarden/server:latest

it will pull the image and set everything up, at which point you should now be able to visit ip or fqdn in a browser and you should be presented with bitwarden page. 

So long and Thanks for all the fish!

