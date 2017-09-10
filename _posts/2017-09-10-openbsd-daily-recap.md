---
layout: post
title: OpenBSD Daily Recap
date: '2017-09-10T00:00:00.000-02:00'
author: Adam Wołk
image: /images/yearinthelife_right.jpg
tags:
- openbsd
---

The last OpenBSD daily read happened on August 18th. I decided to cancel the next one that was scheduled for August 25th and promised to do a recap on this blog over the weekend.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Had no votes for an <a href="https://twitter.com/hashtag/OpenBSD?src=hash">#OpenBSD</a>-daily so cancelling it for today. Will digest and throw out a blog post over the weekend.</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/901164301259558912">August 25, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

It's been over two weeks since that tweet so here is the overdue recap.

The daily [started in june](https://blog.tintagel.pl/2017/06/09/openbsd-daily.html) we kept up doing it daily for 42 days straight after which the schedule changed to weekly and continued for 5 more reads totalling 47 code reads by 3 people with ~130 people idling on the channel during most reads and seeing active participation from around 5.

All reads are archived in this [github repository](https://github.com/mulander/openbsd-daily) which was compiled & curated by @niamtokik (thanks!).

Why did the schedule change to weekly? Reading code on IRC in the format I devised proved to have some drawbacks that I didn't anticipate. It's very challenging to read code you see for the first time in your life, make sense of it and constantly update other people via short text messages and links to the code you are looking at. This leaves very little room for exploratory changes, especially in the kernel where the whole channel would have to wait for me to reboot before I could see the effects of what I did. That also meant not picking on low hanging fruits that were spotted during reads, if my evening is spent on reading code I am not stopping to actually do some changes that I think could be helpful, initially that was mitigated by people participating in the read attempting to do that work while I kept reading, hence making the effort much more worthwhile. I am extremely grateful to duncaen and DuClare for doing reads in the early days of the channel and actively participating in other code reads - that helped me kept going daily for 42 days in a row. With all of that said, I did notice the quality of my reads diving down very quickly as real life started catching up making me go into code reads completely unprepared

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">No <a href="https://twitter.com/hashtag/OpenBSD?src=hash">#OpenBSD</a>-daily today. Need to think how the format should continue. Reading daily without an overall plan started to degrade read quality</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/886350379616698370">July 15, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Going weekly was supposed to be the remedy. I still hoped that other people would pick up and fill in the missing days with their own code reads. I was even thinking that stopping doing dailies might make some people come up front now that each day isn't booked ahead by me - sadly this didn't happen. Activity on the channel started to dwindle, including during code reads - to a point where the last one didn't receive a single vote directing on what should we do on the read. Having no people interested in the reads beats the purpose of doing it since digesting the code for the channel as noted above doesn't leave much wiggle room for actually hacking on it. That is the point when I decided to cancel further reads for now.

Having everything said, there were very nice positive outcomes from the reads. They resulted in 11 commits to the OpenBSD source tree, personally I am most happy about the OpenSMTPD connection leak being diagnosed and fixed by DuClare during our reads. It was also great learning how the OpenBSD malloc and dynamic library loading works. I am completely convinced that reading the code is the best way to learn about the system and continue to do so on a daily basis - just on my own pace.

In summary, I think IRC is a terrible medium for live code reads and unfortunately I don't see myself doing twitch code reading streams ;) yet IRC is a very good medium for obtaining help when you get stuck. You can see that happening many times in my reads where DuClare jumped in to help me decipher specific bits of C and I saw this happen many times outside reads where people were discussing pieces of code they were reading and had questions about. I would like to encourage everyone to use the #openbsd-daily freenode channel exactly for that purpose. This idea is not dropped, I'm just putting a broken format on hold until I find a better way to do live reads. In the meantime, drop by and tell us about the code you have been reading - who knows, maybe we can help you make a crack in a brick wall you have been hitting for a while?