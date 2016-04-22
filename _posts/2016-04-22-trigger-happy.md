---
layout: post
title: Trigger happy
date: '2016-04-21T20:10:00.000-02:00'
author: Adam Wołk
image: /images/Ponderosa.jpg
---

I became trigger happy and I'm not happy about it. What do I mean by that? Not taking enough time to think through a decision and checking all the facts.
[Ducking it](https://duckduckgo.com) before I head to the [man pages](http://man.openbsd.org) or [project documentation](http://www.openbsd.org/faq/). Reading and basing my solutions on Stack Overflow even though I knew that it's a really toxic thing to even read.

The end result are sub-par solutions, simple mistakes and feeling embarrassed by the amount of mistakes that can be pointed out in my contributions. This has been going on for a long time and has to change. I tried to identify the main cause of this and it's a combination of laziness & the ease of use of search engines. It's really easy and immediatly rewarding to slap a few search queries into $SEARCH_ENGINE, coming up with a 80% solution that is good enough. Why check your language documentation for that API if a single query shows you the proper usage in a few seconds? Yeah it's a trap. Suddenly you are not producing code & solutions. You are slapping together examples of existing solutions to make the square peg fit into the round hole. I made that error a lot initially when starting with ports. In retrospect my biggest problem was starting by copying an existing port and modyfing it to fit the software I was porting. You can't imagine how much easier (yes *easier*) it is to start off with /usr/ports/infrastructure/templates/Makefile.template.

With the above in mind I challenged [bmercer@](https://twitter.com/knowmercymod) to join me for a week without using a search engine for any tasks. You can listen to his experience of the experiment in the [latest garbage.fm](http://garbage.fm/episodes/23) podcast. This post documents my personal findings.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">My one week challenge of not using ANY search engines at all starts today. 16.04.2016 00:00. Wish me luck!</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/721095778626838528">April 15, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

First thing that immediately struck me is how useless my bookmarks are. Instead of a organized set of documents I might want to visit again, they are just not there. I mostly know all the software I need to use and apparently it was easier to search the product/project/library name versus using boomkarks. I wondered if I am alone in this and (a rather small) poll on twitter shows I'm not the outlier.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">How do you use browser bookmarks?</p>&mdash; mulander (@mulander) <a href="https://twitter.com/mulander/status/722488972065271808">April 19, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Obviously reading documentation takes time. I found myself taking things much slower, being less annoyed by not arriving at a solution fast enough. Not having to wade through unrelated content during searches helped in my regular work. Yeah I had to use GNU info to read the GNU Make manuals. In the end it allowed me to solve a Makefile problem that I wasn't able to solve before by just trying to quickly $search for facts. The terms were just to hard to get a good search engine match but were things covered almost at the beginning of the related manuals.

I did fail a couple of times. Catching myself looking at search engine results and realizing that "oh shit I used the $search". It happened for two things mostly. Currency conversions, stuff like 20 USD in PLN queries and having to solve issues under pressure - on one day a server failed for $work and I started searching for related info before I even started to think about it.

In the end. I do feel like I improved my whole workflow with this experiment. I still make mistakes (obviously) and still squeeze the trigger too son. Though in just this week alone those are single shot mistakes instead of series of bad decisions of trying to get things done ASAP.

Give it a shot. Hang up the search engine guns for a week and take things slow.
