---
layout: post
title: One of the three OpenBSD users
date: '2015-11-22T17:00:00.000-02:00'
author: Adam Wołk
image: /images/three-of-us.png
tags:
- openbsd
---

Recently I have been taking a look at [Syncthing](https://syncthing.net/) in the hope of being able to replace [ownCloud](https://owncloud.org/) with it after I received a [recommendation](https://news.ycombinator.com/item?id=10208565) on HN.

Initial impressions were quite good. It's obviously a different category, closer to initial use case represented by Dropbox versus the full software suite including nice web clients that are now expected out of Dropbox like solutions. My current ownCloud instance sports a full featured ownNote which my wife loves and it will be really hard to get rid of but everything else about the service felt bad. Keeping it up to date on OpenBSD -current is a pain in the ass (like with every software installed from ports), desktop sync clients tend to fail when connecting to the service & mobile clients experience the same issues (including the paid ones). My initial setup worked OK but a few snapshot upgrades further and I'm quite happy when the basic online functionality works (I don't even dream of uploading large files to it). Hence the venture into Syncthing.

Syncthing is an application written in Go which already makes deployment a lot easier. OpenBSD does not provide a port package for it which is somewhat understandable considering that each user should run their own instance. Although it's not a huge issue since the upstream project provides official OpenBSD binaries which are additionally [signed](https://syncthing.net/security.html).

My initial test drive lasted more than a week and during that time the project released a few minor releases. My Linux boxes automatically upgraded (and those automatic upgrades can be disabled) which I really liked. Unfortunately my OpenBSD machines didn't follow the same path erroring out with a sad `upgrading: readlink /proc/curproc/file: no such file or directory`.

Most of you know that I'm running current and that the /proc filesystem is no longer a thing on OpenBSD. It wasn't hard to find other people reporting the [same problem](https://github.com/syncthing/syncthing/issues/1272) in the upstream providers repository. Opened sometime in January, I decided to subscribe to any updates and continued to manually update my OpenBSD Syncthing instances until I got a notification that the issue was closed.

When I saw that the problem is generally dismissed I replied back with a naive solution. To which one of the upstream developer [replied](https://github.com/syncthing/syncthing/issues/1272#issuecomment-158780141):

"Maybe one of the [three](https://data.syncthing.net/#metrics) OpenBSD users feel strongly enough about this to propose a patch. :D"

This kind of grinded my gears in a positive way. I have been using Linux since the early 90s. I remember people using that same excuse to completely ignore or close issues reported by a completely valid userbase. Guess what? That platform somehow matters now. Yet a keen reader will notice the leading screenshot to this post. The gaming industry is still treating Linux as the red headed step child.

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr"><a href="https://twitter.com/ramma_gaming">@ramma_gaming</a> <a href="https://twitter.com/KittyMeowGames">@KittyMeowGames</a> Linux is a second class citizen, we don&#39;t run it internally because only 17 people use it</p>&mdash; Garry Newman (@garrynewman) <a href="https://twitter.com/garrynewman/status/615071229947564032">June 28, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

By now, the [gaming on linux](https://www.reddit.com/r/linux_gaming) subreddit has over 30k subscribers and Steam machines are rolling out. If you time travelled back to the early 90s and told me that we would be THAT relevant I would laugh in your face.

Though the Linux community did one thing differently than the OpenBSD one does. It's pretty open about what software they use. After that bug report I enabled anonymous statistic reporting for Syncthing essentially **doubling** their OpenBSD user base. It's a bit sad. I went out of my way to disable it and I'm sure most OpenBSD users did the same as privacy matters, in doing so we degrade the platform in the eyes of the software providers. Maybe all of us should be a bit more loud mouthed and get ourselves heard?

All in all I am happy that one of the three OpenBSD users fixed an upstream library used by the Syncthing project. So thank you, [ajacoutot@](https://github.com/kardianos/osext/commit/b4814f465fb1f92d46e37f7ef84d732ece7c3e3a) you saved me some time today. :)
