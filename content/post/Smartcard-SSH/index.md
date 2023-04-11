+++
categories = ['Technology', 'Tips & Tricks']
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
tags = ['featured', 'tech', 'SSH', 'Terminal']
thumbnail = ''
title = "Smartcard SSH"
toc = true
usePageBundles = true
+++

## Intro

So I love Smartcards! I dont know why but I love the idea of physical passwordless authentication, I am also an early adopter of the Yubikey, My env. is configured for both Smartcard and FIDO2 Auth with the yubikeys able to auth either Smartcard Cert or FIDO2 (I have yet to find a Dual-Interface Smartcard that has both Contacted/Contactless PIV as well as FIDO2, I have found cards that have either PIV or FIDO2 but not both so if you know of one drop it in the comments!)

I have been using Smartcards for Auth for years, Same for the Yubikey, Something that drew me to the yubikey was the ability to use GPG, initially I used gpg on the yubikey to sign ssh keys and by using a pageant application I was able to use the gpg ssh key on the yubikey for ssh on windows. This was honestly a bitch to setup, Eventually the yubikey gained the ability to use FIDO2 for SSH as well as FIDO2 became supported in OpenSSH as a standard feature as well so that made the experience better but was not perfect, It only worked with the newer powershell openssh package, not the one that comes with windows and is isntalled in the System32 folder but the one you download and install manually and add to your path which is not the greatest as its just one more thing that needs to be installed on a new machine, So I started looking for another option, I found it but its still not perfect but seems like its as perfect as I will be able to get right now. 

## Pageant

Alright so Putty has an offshoot/fork called Putty-CAC that is enabled for Smartcards AND FIDO2, It ships with a Pageant application, This pageant application can output an openssh compatible pipe with the *--openssh-config C:\Path\Here\Pageant.conf* argument to the pageant.exe and then adding the *Include pageant.conf* to your ssh conf in *C:\ProgramData\ssh* This works awesome with powershell or cmd ssh alias however it doesnt work in MobaXterm which I use a lot. (It does work normally with MobaXterm but when pageant is run with the openssh config arguement it seems to break MobaXterm, If you know of a way to get it to work with both drop it in the comments!) so since I use MobaXterm a lot but also use the windows terminal ssh a lot as well I wrote a bat file to start pageant with the OpenSSH Compatible config and added pageant to my PATH. When I want to switch between OpenSSH and MobaXterm I just kill pageant and run either pageant from my PATH (for MobaXterm) or the Bat file (for OpenSSH from the Terminal). 

Like I said its not perfect, having to switch back and forth like that kinda sucks but considering before now It only worked with mobaxterm and for OpenSSH I had to use a key file on my disk which I am not a fan of, So this will work until something else comes along that gives me both at the same time! Below is the bat file, all you need is to have [Putty-CAC](https://risacher.org/putty-cac/) installed. (Run pageant without the openssh config first and import your keys and set it to remember your keys before using the bat file so that it already knows what CAPI or FIDO2 key to use.)


## Bat File
    
    @echo off
	cd "C:\Program Files\PuTTY-cac\" 
	pageant.exe --openssh-config C:\ProgramData\ssh\pageant.conf

So long and Thanks for all the Fish!