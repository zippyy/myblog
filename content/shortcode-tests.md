+++
title = "Shortcode Tests"
description = "A post that tries out this set of shortcodes."
date = "2016-12-29T13:39:24-05:00"
draft = false
+++

Hello,
Thanks for downloading and testing our my shortcode pack. Below will
load all of the shortcodes and display the various options they have.
If you like and enjoy these short codes, please post a link on your
blog to [GeekThis.net][geekthis], I would really appreciate it.

## Installing New Shortcodes
To install shortcodes, copy each HTML file relating to a shortcode you
want and place it into the shortcode directory inside of Hugo. This
location will vary depending on if you are adding shortcodes to a
template or to your website. If you are adding shortcodes just to your
website, I suggest adding them to the layout directory instead of to a
theme.

To add a shortcode to your Hugo site, place the HTML file inside of
`/layouts/shortcodes/`. For templates add the shortcode file to
`/themes/theme-name/layouts/shortcodes/`.

## Twitch Shortcode
With Twitch shortcodes, you can specify a variable of either channel or
video. If both are set for a single shortcode, only the video variable
will be set and channel will be ignored.

To get a video ID from Twitch, visit a video page and the ID is the
number at the end of the URL, along with the letter in the path name
before it. With the example of `https://www.twitch.tv/example/v/123456789`
the video ID would be `v123456789`.

{{< twitch channel="example" >}}

{{< twitch video="v123456789" >}}

## JSFiddle Shortcode
JSFiddle shortcodes embed a read-only JSFiddle and displays the HTML,
CSS, JavaScript along with the Result. You can customize the appearance
to be either `light` or `dark`, adjust the height, and select what
tabs to display.

To get the JSFiddle ID and Revision Number, visit a JSFiddle and the
alphanumeric number as the first section in the path is the identifier.
The revision number is the numeric number in the next part of the path.
The revision number may not exist can be set to 0 or ignored. With
the example url of `https://jsfiddle.net/347cw3h4/1/` the ID is
`347cw3h4` and the revision number is `1`.

{{< jsfiddle id="347cw3h4" rev="0" color="dark" view="js,html,css,result" height="200" >}}

## Pastebin Shortcode
Pastebin is a great place to paste source code, logs, and other content
that you want to share with others. If you want to share it with your
website visitors, you can use the pastebin shortcode.

The pastebin shortcode has two options you can set, the pastebin id
and the height of the paste to display. The heigt is optional, and if
not set, the full height of the paste will be displayed.

To get the pastebin ID, go to the paste you want to publish, and in the
url, the alphanumeric string directly after the pastebin domain is the
pastebin id. Using the URL `http://pastebin.com/JwzQbTwh` as an example,
the pastebin ID would be `JwzQbTwh`.

{{< pastebin id="JwzQbTwh" height="500px" >}}

## Google Maps
Adding Google Maps can sometimes be useful to add inside of posts or
onto pages of your website. Our shortcode provides the general Google
Maps Place API to embed a Google Map with a search query. A search can
include an address, city, street, or almost any other location.

With Google Maps, you need to first setup a Google Maps API Key. It's
free and easy to setup. Visit the [Google Developer API
Console][google-console] and create a new application to get your API
key. You can store this key inside your Hugo configuration file or each
time you ues the shortcode, you can add the API Key to the parameter
`key`.

{{< highlight ini >}}
[params]
	googleMapsAPIKey = "000000000000"
{{< /highlight >}}

The Google Maps shortcode accepts the parameters of `height`, `key`,
and `location`. The height defaults to 400px if it's not set and the
api key defaults to the site configuration API Key. The location is a
search query for a location you want to display for the map.

{{< google-maps key="000" height="300" location="Googleplex, Amphitheatre Parkway, Mountain View, CA" >}}

## Box Documents
If you use Box to share files, adding embedded files to your website
is easy. Box provides an embedded view for shared files that displays
a preview of almost any file available, such as PDF, DOC, Spreadsheets
and others.

The box shortcode accepts two parameters, the share ID of the file and
the height you want the document to be displayed at. To get the ID of
a file inside of Box, right click on a file inside of Box and click on
"Share." The publicly available URL will be similar to
`https://app.box.com/s/000000000000`, you will want to grab the unique
string of characters ofter the `/s/` path.

{{< box-preview id="000000" height="400" >}}


[google-console]: https://console.developers.google.com/
[geekthis]: http://geekthis.net
