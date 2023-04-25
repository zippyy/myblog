+++
author = 'Nicholas Bennett'
date = '2021-10-13T21:48:34Z'
year = "2021"
month = "2021-10"
description = ''
draft = false
image = ''
slug = 'manage-local-admin'
title = 'Manage Local Admin'
tags = ['Tech','Sysadmin']
+++


So working in the MSP space, Something that comes up a lot is managing Local Admin Access, Admin Access is something that can be daunting if you have a large Env. or have a large amount of users in your Env. that require access to local admin on a consistent basis.

Enter stage left [MakeMeAdmin](https://github.com/pseymour/MakeMeAdmin), I found make me admin years ago while browsing through github for similar projects, It has been a godsend for me in clients that cant afford or do not want to spend the money on something like [AutoElevate](https://www.autoelevate.com/).



[Make me admin](https://github.com/pseymour/MakeMeAdmin) is a great tool that allows you to install a client on your endpoints, create a security group in AD and point the Access to that groups SID, Then when users that are part of that security group need to elevate to admin access they request a time based elevation via an application icon in the system tray. When a user requests elevation they will be temporarily moved to the local administrators group and after a set time limit they will be automatically be removed from the group.

This has made my life so much easier and actually has increased productivity of our Help Desk and Tier 2 Staff that are usually responsible for handling those types of tickets. I have rolled this out at 2 previous employers for many of their clients and its has worked great.



I will say that it is not as great as [AutoElevate](https://www.autoelevate.com/) due to the fact that you must trust the users in the escalation group to be weary of threats when they are elevating themselves but it will cut down on attack vector because the elevated vector is only available for a short duration when requested so something like a malicious email attachment or link will only have effect if clicked while elevated. I will do another post on AutoElevate as it is an incredibly useful tool which provides way more security while also increasing productivity in a similar metric.

So long and Thanks for all the fish!