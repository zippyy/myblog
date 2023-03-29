+++
author = 'Nicholas Bennett'
date = '2022-01-22T02:58:03Z'
description = ''
draft = false
series = 'Hugo Netlify Setup'
slug = 'hugo-netlify-easy'
title = 'Hugo With Github & Netlify Part 2'
+++

Okay so for this side of the fork in the road we will be using the netlify account you setup in part 1 and we are assuming that you signed up for netlify with your github account and not just normal email sign up so your repos will already be listed within netlify. 

Lets get started, open netlify and click Add New site > Start from Template > Hugo Blog

Click the button to gonnect your git provider, in this case github and then choose a name and select your desired privacy level for the repo (public or private) and then click deploy site. 

it will bring you to the netlify builds page and you should see your randomly assigned netlify site name (we will go over how to change this) and right under that youll see "Site Deploy In Progress" give it a few minutes and it should show a link under your site name, If you click on that link it will take you to your site.

now lets change the site name and netlify url (we will cover later how to setup a custom domain as well). Right below your site name you will see a button called site settings, click on that and go to first card on that page and you should see a button that says change name, click on that and enter a name for your site, for this I chose hugo-tut which gave me a url of hugo-tut.netlify.app which I will leave up incase you want to see what the final product is like before you doeploy your own! 

In the next post we will cover how to do this the more manual route and pick a custom theme to start from!

So long and Thanks for all the fish!


