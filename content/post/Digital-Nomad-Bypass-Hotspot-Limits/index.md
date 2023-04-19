+++
categories = ['Technology', 'Digital Nomad', 'tech', 'Pinned']
codeLineNumbers = false
codeMaxLines = 10
date = "2023-04-18T23:16:27-06:00"
year = "2023"
month = "2023-04"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', 'Digital Nomad', 'tech', 'Pinned', 'TTL']
thumbnail = ''
title = "Digital Nomad Bypass Hotspot Limits for Router"
toc = false
usePageBundles = true
series = 'Digital Nomad'
weight = 4
+++

Alright, In the [VPN Post](https://techrelay.xyz/post/nomad-vpn) I mentioned I would add other posts to the series to cover useful things, One of which I mentioned was [Bypassing Hotspot Limits](https://techrelay.xyz/post/bypass-mobile-hotspot-windows/) which is specific to windows and really has the most use for machines with built in WWAN or a Hotspot connected to a windows device but in this post we will cover how to modify the A Travel Router to set the TTL to 65 which in most cases will bypass hotspot limits put forth by carriers, This post will be specific to GL-inet routers but I will do more parts to cover other hardware/software vendors.


1. Login to your TravelRouter admin portal.

2. Go to *More Settings* >> *Advanced*.

3. You will be prompted with another login page, use the same credentials you just used to login to the regular admin portal.

4. Go to *Network* >> *Firewall* and Click on *Custom Rules*.

5. Add a Custom Rule at the bottom to set the outgoing TTL to 65

```
#Change TTL
iptables -t mangle -I POSTROUTING 1 -j TTL --ttl-set 6
```

6. Click *Restart Firewall* and the device should reboot. You should be good to go now as any device that connects to the Travel Router will mask their TTL as if it were a mobile device. 


**If you want to test it for peace of mind and are comfortable with SSH you can connect to the Travel Router and run the command *ping 127.0.0.1* and in the responses you should see the TTL as 65.**


I hope this post was helpful and keep an eye out for more parts covering other solutions! Let me know what you do to bypass hotspot limits or request a guide on a specific brand in the comments below!

So Long and Thanks for all the Fish!