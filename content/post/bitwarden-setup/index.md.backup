+++
author = "Nicholas Bennett"
date = 2022-02-10T02:58:03Z
description = ""
draft = false
slug = "bitwarden-setup"
title = "Bitwarden Setup"
series = "Hosting Bitwarden"
+++

Alright so lets get started. You will need a few prerequisites as follows:

+ Github Account.
+ Heroku Account. (this can be done with a free account but you need to be verify your free account.)
+ A Free service like Freshping. (or some other pingdom alternative to send continuous pings to the herko instance to keep it from sleeping.)

# **_Deployment_**

Follow these steps:

1. Fork this [Repo](https://github.com/davidjameshowell/vaultwarden_heroku "Vaultwarden Heroku") on github.
2. Edit the file the file in the code block below to enable/disable Duo Auth.
```
.github/workflows/deploy.yml

DUO_ENABLE: 1 or 0
```

3. Go to your forked repo and click Settings > Secrets > Actions and add secrets for:


+     HEROKU_API_KEY
 	Your Heroku API key - can be found in Account Setings -> APi Keys
+     HEROKU_APP_NAME
 	The name of the Heroku application, this must be unqiue across Heroku and will fail if it is not [Value alphanumerical]
+     HEROKU_VERIFIED
 	Required regardless, if you have added a credit card on, your account will be verified to use built in addons, if not please see "NON VERIFIED ACCOUNTS" section [Value 0/1]


4. In your Github repo page click on the actions tab and then click the BitwardenRSOnHerokuAIO_Deploy job and wait for it to complete. 

5. The job with run through github actions and deploy to your heroku account via the API info that you filled in earlier. This should take about 15 minutes and you can keep an eye on it through the output on the screen.


# **_Update_**

Updating is simple and can be done one of two ways:

+ Running the workflow manually via Github Action 
+ Making a commit to the main branch, forcing a Github Actions workflow to initiate

*Either one of these will force the Github Actions workflow to run and update the app. If you need to modify to enable/disable settings, you should re run it as well. *

In the next post we will discuss how to host this same Vaultwarden RS server at home with docker!
