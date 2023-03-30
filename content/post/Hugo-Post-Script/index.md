+++
categories = ['Technology']
codeLineNumbers = true
codeMaxLines = 10
date = "2023-03-30T07:15:55-06:00"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', '']
thumbnail = ''
title = "Hugo Post Script"
toc = false
usePageBundles = true
+++

I bounce around between Windows, Linux and MacOS, I have long since had a shell script that creates a new post with a page bundle and inserts the front matter and opens the file in sublime text. I wanted to do the same on windows with powershell so I could add the script to my path and call it from anywhere, I have a folder named bin in the root of C that is in my system path. You can find the script below.

```
# Change directory to blog root directory
cd "C:\Path\to\hugo\project"

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
"@ | Out-File -FilePath "content\post\$path\index.md" -Encoding UTF8

# Open the new post in Sublime Text
& "C:\Program Files\Sublime Text\sublime_text.exe" "content\post\$path\index.md"
```

Here is the Bash version as well

```
#!/bin/bash

# Prompt the user for the path to the new post
read -p "Enter the path to the new post (e.g. my-new-post): " path

# Create the new post using a page bundle
mkdir -p "content/post/$path"
touch "content/post/$path/index.md"

# Add front matter to the new post
cat <<EOT >> "content/post/$path/index.md"
+++
categories = ['Technology']
codeLineNumbers = false
codeMaxLines = 10
date = "$(date +"%Y-%m-%dT%H:%M:%S%z")"
description = ''
draft = false
featureImage = ''
featureImageAlt = ''
featureImageCap = ''
figurePositionShow = true
shareImage = ''
tags = ['featured', '']
thumbnail = ''
title = "$(echo "$path" | sed -E 's/-/ /g' | awk '{print toupper(substr($0, 1, 1)) substr($0, 2)}')"
toc = false
usePageBundles = true
+++

**Insert Lead paragraph here.**
EOT

# Open the new post in Sublime Text
subl "content/post/$path/index.md"
```

The script will change directory to the hugo project (You need to edit the line 2 to the correct path to your hugo directory) create a folder with an index.md using page bundles and add the front matter to the post. (You can edit the front matter to suite your needs just make sure that the output is UTF8 encoded, I had issues with posts not publishing or them showing up with weird characters instead of my text.)

Now your off to the races, This script has mad it so much easier to write content for hugo without having to remember syntax or long commands. 

So long and Thanks for all the fish!