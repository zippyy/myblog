+++
categories = ['Technology']
codeLineNumbers = false # Override global value for showing of line numbers within code block.
codeMaxLines = 10 # Override global value for how many lines within a code block before auto-collapsing.
date = 2023-03-30T14:30:44-06:00 # Date of post creation.
description = '' # Description used for search engine.
draft = false # Sets whether to render this page. Draft of true will not be rendered.
featureImage = '' # Sets featured image on blog post.
featureImageAlt = '' # Alternative text for featured image.
featureImageCap = '' # Caption (optional).
figurePositionShow = true # Override global value for showing the figure label.
shareImage = '' # Designate a separate image for social media sharing.
tags = ['featured', ''] # tags for hugo, this should always have featured because featured = true is not working.
thumbnail = '' # Sets thumbnail image appearing inside card on homepage.
title = "Git Auto Commit" # Title of the blog post.
toc = false # Controls if a table of contents should be generated for first-level links automatically.
usePageBundles = true # Set to true to group assets like images in the same folder as this post.
# comment: false
+++


I have been using hugo for over 2 years now going on 3 now, I used ghost before but its really heavy and I wanted to ditch needing to host a webserver in favor of static site hosting which I have covered in previous posts, when I moved from ghost to hugo I used a converter to convert all my posts to markdown for a seamless transition to hugo.

since hugo uses git and I host my page with netlify I either live in github desktop with sublime text or I am in the terminal using nano or vim which due to the nature of git, requires me to commit my edits, and push those to the origin. 

that can get tedious and sometimes I forget to commit or push my changes so I started to look into auto commit options which for windows seems to be few and far between, I ended up trying all kinds of scripts and "apps" designed to monitor a git repo folder and commit and push any changes, none of them worked, every single one I tried I had issues getting it to run right as a service or scheduled task so I decided to enlist chatGPT which actually wasnt a huge help but it did help me isolate my issues and get a working script running as a windows service using nssm. 

after testing I decided that this might not be the best idea for a blog as sometimes I end up scraping content or making a bunch of small changes and running the hugo server locally before I push so I chose not to use the auto-commit script after all however I found another use for it

I have a folder on the root of my C:\ called Bin and this folder has been added to my path to allow usage of scripts, binaries, etc... that I may want to use without calling an explicit path, scripts like my TTL changer (My laptop has an LTE card in it with a SIM that only gets 10GB of data a month for hotspot but unlimited regular data thus the need to switch back and forth between 60 TTL and 128 TTL) that I use on a very regular basis so I thought why not use this script to sync that bin folder to github (its actually a symlink now to the github repo that I moved all the scripts onto)

below I will share the batfile as well as the powershell script I wrote to make it all work along with the right settings for nssm.

All in all it was a good learning experience and a reminder that sometimes you just cant use something pre-made and have to figure out a way to do what you want on your own. 

So long and Thanks for all the Fish!