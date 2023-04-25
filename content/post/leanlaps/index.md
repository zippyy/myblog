+++
author = 'Nicholas Bennett'
date = '2023-03-04T01:33:53Z'
year = "2023"
month = "2023-03"
description = ''
slug = 'leanlaps'
tags = ['featured','Tech','Cloud']
title = 'Lean LAPS'
+++

For a long time I have just used a powershell script deployed via intune to manage a local admin for emergencies, this is not the best option but I dont have that many devices to manage and honestly most of them never leave the environment, I also use powershell to rotate this password. 

With Microsoft on the verge of releasing the updated LAPS option that works with AAD, I wanted to try another option in the interim while I wait for that feature to roll out, I settled on LeanLAPS and man was it easy to setup.

Getting going was a simple as downlading the script, editing the config section to set variables and configure things like excluded admins or groups if you have the auto remove other local admins variable enabled, and then create the remediation policy, upload the script and set your applied groups and schedules for each group, deploy and bam!

I have struggled to get the GUI app to work but its not been an issue as the password is displayed in the intune portal also. 

check it out at [LeanLAPS](https://www.lieben.nu/liebensraum/2021/06/lightweight-laps-solution-for-intune-mde/)