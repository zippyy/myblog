+++
categories = ['Technology']
codeLineNumbers = false # Override global value for showing of line numbers within code block.
codeMaxLines = 10 # Override global value for how many lines within a code block before auto-collapsing.
date = {{ .Date }} # Date of post creation.
year = {{ .Date | dateFormat "2006" }} # year for archives.
month = {{ .Date | dateFormat "2006-01" }} # month for archives.
description = '' # Description used for search engine.
draft = false # Sets whether to render this page. Draft of true will not be rendered.
featureImage = '' # Sets featured image on blog post.
featureImageAlt = '' # Alternative text for featured image.
featureImageCap = '' # Caption (optional).
figurePositionShow = true # Override global value for showing the figure label.
shareImage = '' # Designate a separate image for social media sharing.
tags = ['featured', ''] # tags for hugo, this should always have featured because featured = true is not working.
thumbnail = '' # Sets thumbnail image appearing inside card on homepage.
title = "{{ replace .Name "-" " " | title }}" # Title of the blog post.
toc = false # Controls if a table of contents should be generated for first-level links automatically.
usePageBundles = true # Set to true to group assets like images in the same folder as this post.
# comment: false
+++

**Insert Lead paragraph here.**