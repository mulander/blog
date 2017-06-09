---
layout: post
title: OpenBSD Daily
date: '2017-06-09T00:00:00.000-02:00'
author: Adam Wołk
image: /images/yearinthelife_right.jpg
tags:
- openbsd
---

I made a new years resolution to read at least one C source file from OpenBSD.
The goal was to both get better at C and to contribute more to the base system and userland
developmnet. I have to admit that initially I wasn't consistent with it at all. In the first quarter
of the year I read the code of a few small base utilities and nothing else. Still, every bit counts and it's
never too late to get better.

Around the end of May, I really started reading code daily - no days skipped. It usually takes
anywhere between ten minutes (for small base utils) and one and a half hour (for targeted reads). I'm pretty happy with
the results so far. Exploring the system on a daily basis, looking up things in the code that I don't understand
and digging as deep as possible made me learn a lot more both about C and the system than I initially expected.

There's also one more side effect of reading code daily - diffs. It's easy to spot inconsistencies, outdated code or
an incorrect man page. This results in opportunities for contributing to the project. With time it also becomes less
opportunitstic and more goal oriented. You might start with a [drive by diff to kill](https://marc.info/?l=openbsd-tech&m=149591302814638&w=2) optional compilation of an old compatibility option in chown that has been compiled in by default since 1995.

<br />

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">chown was standardised in 1992 deprecating the `.` dot separator - 25 years later all chown implementations still support it due to wide use</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/868577622258851841">May 27, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Soon the contributions become more targeted, for example [using a new API](https://marc.info/?t=149677112300004&r=1&w=2) for encrypting passwords
in the htpasswd utility after reading the code of the utility and the code for htpasswd handling in httpd. Similarly it can take you from discussing
a doas feature idea with a friend to [implementing it](https://marc.info/?t=149694587100005&r=1&w=2) after reading the code.

I was having a lot of fun reading code daily and started to recommend it to people in general discussions. There was one particular
twitter thread that ended up starting something new.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Pick C. Read one source file from <a href="https://twitter.com/hashtag/OpenBSD?src=hash">#OpenBSD</a> /usr/src daily. It&#39;s fun. Me and <a href="https://twitter.com/tuxbsd">@tuxbsd</a> have been doing this lately. We could do group reviews.</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/871777396743184385">June 5, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Just registered <a href="https://twitter.com/hashtag/openbsd?src=hash">#openbsd</a>-daily on freenode. Feel free to jump in.</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/871780638805950464">June 5, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

This is still a new thing and the format is not yet solidified. Generally I make a lot of notes reading code, instead of slapping them inside a local file I
drop the notes on the IRC channel as I go. Everyone on the channel is encouraged to do the same or [share his notes](https://github.com/bsdtux/openbsd-daily) in any way he/she seems feasable.

In the first 3 days of the channel we:

 - read the source of httpasswd (sorry, I didn't grab the log for that one)
 - [learned how htpasswd files are handled in httpd](https://junk.tintagel.pl/openbsd-daily-httpd.txt)
 - [read the code of doas](https://junk.tintagel.pl/openbsd-daily-doas.txt)
 - [implemented a new feature in doas](https://junk.tintagel.pl/openbsd-daily-doas-confirm.txt)

I'm having fun so far, reading code in a group is proving to be a great way to learn. Feel free to drop by on #openbsd-daily @ Freenode! ;)
