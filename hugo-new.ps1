# Prompt the user for the path to the new post
$path = Read-Host "Enter the path to the new post (e.g. my-new-post)"

# Create the new post using a page bundle
New-Item -ItemType Directory "content\post\$path"
New-Item -ItemType File "content\post\$path\index.md"

# Add front matter to the new post
@"
+++
categories = ['Technology']
codeLineNumbers = false
codeMaxLines = 10
date = "$(Get-Date -Format "yyyy-MM-ddTHH:mm:sszzz")"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', '']
thumbnail = ''
title = "$(echo $path | %{ $_ -replace '-',' ' } | %{ $_.Substring(0,1).ToUpper() + $_.Substring(1) })"
toc = false
usePageBundles = true
+++

**Insert Lead paragraph here.**
"@ > "content\post\$path\index.md"

# Open the new post in Sublime Text
#& "C:\Program Files\Sublime Text 3\sublime_text.exe" "content\post\$path\index.md"
