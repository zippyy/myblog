+++
categories = ['Technology']
codeLineNumbers = false
codeMaxLines = 10
date = "2026-01-21T16:18:48-07:00"
year = "2026"
month = "2026-01"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', '']
thumbnail = ''
title = "Cellular Era"
toc = false
usePageBundles = true
+++

We are totally in a Cellular Era right now, Cellular modems have come so far with more and more towers being updated as well. From Things like 5G SA and mmWave which are changing the came with insnane Gigabit or even Multi-Gig speeds to Satellite connections Direct-to-Phone. 

I have been using Nighthawk Hotspots forever, They consistently use the best available modems at any given time and release models with updated features on the reg, However they are very pricy and do not always get the best signal without external antenna setup's. 

I got my hands on the Puli AX and man is it an amazing router, It might be a little (okay, A alot. lol) on the large side however it makes up for that with its long battery life and amazing 5G connection. I have taken this across the country a few times while driving back and forth to the in-laws and it does great even in the more rural areas of the drive where towers are fewer and far between or have older hardware.

We drove almost 2000 miles across 3 states to see the in-laws for multiple holidays and the Puli AX never missed a beat! I have it setup with Tailscale so I can always access anything from my tailnet no matter where I am (this is nice because I have stuff in the tailnet that is not in my home network, like stuff at the in-laws or my parents house.) as well as 2 different wireguard enpoints, one running on my home unifi and another running on my Flint 3 connected to my backup ISP. These are running in policy mode so that A. None of the tailnet traffic ever goes through either of the VPN's, Traffic will first try the Unifi Wireguard endpoint as this is my main network and multi-gig fiber ISP, If that is ever down then it will automagically move on to the Flint 3 endpoint and any traffic that does not match (i.e both tunnels are down) gets dropped (Killswitch style) which is super nice to have as I do not have to rely on carrying both a travel router like my Slate 7 coupled with a hotspot like my MR6500 unless I absolutely am going to need those mmWave speeds. 

All in all I would say that the Spitz AX and Puli AX are in the same league as the Nighthawk MR5100 but slightly behind the way more expensive mmWave capable MR6500/MR6550. If small size and fastest speed are your most important need then go for the M6 Pro MR6500/MR6550, but If your wanting customization and longevity then the Spitz AX or Puli AX are hands down the go to here as they run OpenWRT with unlocked modems (full AT access!) and have some amazing external antenna options!


So long and thanks for all the fish!