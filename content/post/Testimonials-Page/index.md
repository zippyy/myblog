+++
categories = ['Technology']
codeLineNumbers = false
codeMaxLines = 10
date = "2023-04-24T22:02:43-06:00"
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
title = "Testimonials Page"
toc = false
usePageBundles = true
+++

I have been getting a consistent number of people either paying me to complete the [Digital Nomad VPN Setup](https://techrelay.xyz/post/nomad-vpn) or that have set it up on their own, are super satisfied and want to let me know so I decided to start working on a testimonials page for those folks as well as anyone else I do any kind of work for. 

As you know I build this site with Hugo. I started looking for a theme that had a testimonials functionality to see if I could use it on my own site. After a while I found the Hugo-Universal-Theme that had exactly what I wanted, and that was more than just some html to make a carousel or something, I wanted to use TOML files in a directory to pull and display the testimonials in the carousel. 

which is exactly what Hugo-Universal-Theme is doing for their testimonials functionality so I started by taking the partial, the CSS and JS files and the config.TOML excerpt and put them in my site, low and behold I had a testimonials page but the carousel wasn't working, after some help from ChatGPT I realized that the JS and CSS weren't loading right or rather they were loading but jQuery wasn't and its needed so I set off to get jQuery working and come across another issue, something is still not right with the CSS and when I get the CSS to load the page is just blank.

I spent longer than I wanted to up to this point and decided to put it on the back burner and move on to something else, After a few days and several other additions I decided to try another hand at it and this time fully with the help of ChatGPT. I got ChatGPT to fix the ShortCode and properly load the JS files and jQuery, as well as the CSS files but something still wasn't right so I downloaded just the CSS files and put them in my static folder (I am loading everything else via CDN but for some reason the CSS files weren't working right till I put them local) and it was working! the carousel, the nav buttons, but not the dots.

ChatGPT could not seem to get the dots part right and it turned out to be so simple, I decided to head over to stack overflow and found a post that mentioned needing the owl-theme with the owl-carousel class which my code did not have, and ChatGPT seemed to have missed as well although the rest of the code was correct that one thing missing caused the dots to just not show and me to pull my hair out for a while. 

Eventually it is all working now and once I get these testimonials organized and into TOML files I will post the link in this post as well as update the links drop down in the menu with a link to it as well. 

I am really loving Hugo and ChatGPT has really helped me understand more about how it works and get better at building my stuff like Partials and ShortCodes which I had made plenty of at this point. I am just obsoletely in love with Hugo, although it takes some finagling to get some things to work or to take things from other themes its so worth it when you get everything to work right.

*Update: Here is the link to the [Testimonials Page](https://techrelay.xyz/testimonials/)

Head over to the [Services Page](https://techrelay.xyz/services) for a quote on getting a Hugo site built!

As always So long and Thanks for all the Fish!