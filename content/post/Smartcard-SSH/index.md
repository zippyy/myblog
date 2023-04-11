+++
categories = ['Technology']
codeLineNumbers = false
codeMaxLines = 10
date = "2023-04-10T18:17:41-06:00"
year = "2023"
month = "2023-04"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', '']
thumbnail = ''
title = "Smartcard SSH"
toc = false
usePageBundles = true
+++

So I love Smartcards! I dont know why but I love the idea of physical passwordless authentication, I am also an early adopter of the Yubikey, My env. is configured for both Smartcard and FIDO2 Auth with the yubikeys able to auth either smartcard cert or FIDO2 (I have yet to find a Dual-Interface Smartcard that has both Contacted/Contactless PIV as well as FIDO2, I have found cards that have either PIV or FIDO2 but not both so if you know of one drop it in the comments!)

I have been using Smartcards for Auth for years, Same for the Yubikey, Something that drew me to the yubikey was the ability to use GPG, initially I used gpg on the yubikey to sign ssh keys and by using a pageant application I was able to use the gpg ssh key on the yubikey for ssh on windows. This was honestly a bitch to setup, Eventually the yubikey gained the ability to use FIDO2 for SSH as well as FIDO2 became supported in OpenSSH as a standard feature as well so that made the experience better but was not perfect, It only worked with the newer powershell openssh package, not the one that comes with windows and is isntalled in the System32 folder but the one you download and install manually and add to your path which is not the greatest as its just one more thing that needs to be installed on a new machine, So I started looking for another option, I found it but its still not perfect but seems like its as perfect as I will be able to get right now. 

Alright so Putty has an offshoot/fork called Putty-CAC that is enabled for Smartcards AND FIDO2, It ships with a Pageant application, This pageant application can output an openssh compatible pipe with the *--openssh-config C:\Path\Here\Pageant.conf* argument to the pageant.exe and then adding the *Include pageant.conf* to your ssh conf in *C:\ProgramData\ssh* This works awesome with powershell or cmd ssh alias however it doesnt work in MobaXterm which I use a lot. (It does work normally with MobaXterm but when pageant is run with the openssh config arguement it seems to break MobaXterm, If you know of a way to get it to work with both drop it in the comments!) so since I use MobaXterm a lot but also use the windows terminal ssh a lot as well I wrote a bat file to start pageant with the OpenSSH Compatible config and added pageant to my PATH. When I want to switch between OpenSSH and MobaXterm I just kill pageant and run either pageant from my PATH (for MobaXterm) or the Bat file for 




    ```
    @echo off
	cd "C:\Program Files\PuTTY-cac\" 
	pageant.exe --openssh-config C:\ProgramData\ssh\pageant.conf
	```