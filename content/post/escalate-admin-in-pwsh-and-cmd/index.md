+++
author = "Nicholas Bennett"
date = 2021-09-07T20:40:46Z
description = ""
draft = false
image = "__GHOST_URL__/content/images/2021/10/kyle-glenn-dGk-qYBk4OA-unsplash-1-.jpg"
slug = "escalate-admin-in-pwsh-and-cmd"
title = "Escalate Admin in PWSH and CMD Shell"

+++


So as someone that spends a lot of time in a POSIX Compliant Shell (I have several Hypervisors, Docker Hosts, Web Servers, etc...) I was looking for something similar to how privilege escalation works in *nix using SUDO and after a bit of searching I found an awesome tool called [gSUDO](https://github.com/gerardog/gsudo)

After trying several things including a PWSH profile script called GOADMIN that worked for some computers but not others as well as some computers is worked on would display an error when PowerShell was launched but would still work.... This was not ideal as i have a range of devices that this would be useful for so I started searching further to find a real solution that was closer to or exactly the same as how I escalate the user privileges on many of my other devices and thus I found GSUDO.

Install GSUDO:

* Using [Scoop](https://scoop.sh/): `scoop install gsudo`
* Or using [Chocolatey](https://chocolatey.org/install): `choco install gsudo`
* Or manually: Unzip the latest release, and add to the path. Or let the following script do it for you:

```
PowerShell -Command "Set-ExecutionPolicy RemoteSigned -scope Process; iwr -useb https://raw.githubusercontent.co
```

Once you install via one of the above methods you should have an application in your PATH called GSUDO as well as a Symlink to SUDO that when typed in a CMD or Powershell Prompt to escalate to admin within that session

(an awesome thing and one of the main reason I wanted a solution like this is that it will preserve your Working Directory Path, i.e. you can open a file browser and go to a specific path, click in the path window where you see your PWD bread crumbs and type CMD or Powershell to have a prompt open in this directory, This functionality is awesome but opens in a standard user space prompt, all the other methods I found so far will open a new session or reset your PWD when escalated, so by typing sudo in that prompt you will get an escalated prompt in the same session with the same PWD)



I dropped a link to the github repo for this software and want to give a shout to the user who created it @geardog on github, Thank you for creating this wonderful software to fill a gap that should be built into CMD and PWSH.

So long and Thanks for all the fish!