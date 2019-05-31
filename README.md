# MZSearchHints
Quick and dirty (literally) MZSearchHints implementation that translates iTunes into Pornhub.

## How to use
Fiddler works quite well for this type of thing. One way to redirect hints requests with Fiddler is to add the following to OnBeforeRequest:
~~~~
if (oSession.HostnameIs("search.itunes.apple.com") && oSession.uriContains("/WebObjects/MZSearchHints.woa/wa/hints")) {
            oSession.hostname = "[server]";
            oSession.fullUrl = "http://" + oSession.url
}
~~~~

For trending searches, add the following AutoResponder rule:
~~~~
Match: https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/trends
Respond: http://[server]/WebObjects/MZSearchHints.woa/wa/trends/
~~~~

Once you've set up Fiddler correctly and launched MZSearchHints, you should see some rather NSFW trending searches and search suggestions when using the iTunes Store search.
