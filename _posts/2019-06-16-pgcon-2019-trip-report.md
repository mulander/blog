---
layout: post
title: PGCon 2019 Trip Report
date: '2019-06-16T00:00:00.000-02:00'
author: Adam Wołk
image: /images/IMG_20190527_143527-s.jpg
tags:
- openbsd
- postgresql
---

During my flight I finally decided to watch the "Shape of Water". I knew the film was critically acclaimed but I wasn't expecting a lovecraftian theme and it never clicked with me that it was directed by Guillermo del Toro - it was a very fitting thing to watch during a flight above the ocean depths and now I want "At the mountains of madness" to finally happen even more.

The last leg of the flight was operated by Air Canada and thankfully uneventful apart from a small delay. I landed in Ottawa and was at my hotel at around 2 am where I ended up waiting for a battery replacement on a smart-lock.

I had two days before [@oshogbovx](https://twitter.com/oshogbovx) arrival. I spent the first one on a walk along the Rideau Channel. I was astounded by the amount of nature interveined with the city. Trees, flowers (so many tulips) and animals. On my way back I even spotted a beaver or a badger but didn't want to get too close to annoy it. That's in the middle of the city.

![tree with a small animal next to it](/images/IMG_20190527_160802-s.jpg)

You can open up the images in a new tab and remove the trailing '-s' from the file name to see the original quality picture.

![closer picture of the animal](/images/IMG_20190527_160807-s.jpg)

On Tuesday the tutorials of [PGCon](https://www.pgcon.org/2019/schedule/) have started but I haven't been subscribed to them. I decided to spend the day on museums. I was naive in thinking that I can do more than one in a day. I took 811 pictures inside the Museum of History and only managed to cover two displays ([Neanderthal](https://www.historymuseum.ca/media/neanderthal-exhibition-opens-at-the-canadian-museum-of-history-in-a-north-american-exclusive/), [First Peoples Hall](https://www.historymuseum.ca/event/first-peoples-hall/)). They very kindly didn't kick me out when I triggered an alarm by leaning too far over the display to take a picture and only gently hinted that I should probably go when it was well past 15 minutes from closing.

![Halibut hook, octopus design. The Octopus was also called Devilfish](/images/IMG_20190528_145908-s.jpg)
*Halibut hook, octopus design. The Octopus was also called Devilfish*

Mariusz arrived later that evening and we spend the rest of it touring the city itself. The parliament building is very impressive and apparently they do light shows using the building façades. Unfortunately those were scheduled after our planned departure so we didn't see the real deal but we did see a testing sequence they ran over the building in preparation for the seasonal performances.

![Parliament building during the evening](/images/IMG_20190528_195630-s.jpg)
*Parliament building during the evening*

![Parliament building with a light show test sequence](/images/IMG_20190528_212616-s.jpg)
*Parliament building with a light show test sequence*

On the next day the Unconference started. I went in expecting a hackathon experience. The group listed out 26 topics and a vote was taken to pick projects to hack on as the amount of time and rooms were both limited. Finally 12 topics emerged and composed into three tracks.

![PGCon Unconference voting on topics](/images/IMG_20190529_101424-s.jpg)
*PGCon Unconference voting on topics*

We all moved to the room corresponding with our track of interest. At that point I expected everyone to grab their laptops and start hacking on code and experimenting with the idea. That however did not happen. The unconference is essentially a series of unscheduled talks proposed by the attendees and picked by the attendees on site. Don't get me wrong - the topics were extremely interesting and I don't regret participating. I would however love another day scheduled or a track for hands on hacking.

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Unconference <a href="https://twitter.com/PGCon?ref_src=twsrc%5Etfw">@PGCon</a> has been started. And now we know TOP-12 hot topics for developers (not end users probably)! <a href="https://t.co/kXgDgwbp98">pic.twitter.com/kXgDgwbp98</a></p>&mdash; Pavlo Golub (@PavloGolub) <a href="https://twitter.com/PavloGolub/status/1133752833906544643?ref_src=twsrc%5Etfw">May 29, 2019</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

We picked the following path through the unconference:

- Zed store
- Direct IO & Async IO
- Fault injection framework
- Constrain session memory

The re-emerging topic through all of PGCon was [#fsyncgate](https://wiki.postgresql.org/wiki/Fsync_Errors). Both the direct IO and fault injection framework talks directly related to it. During the fault injection talk questions were raised how to test system level misbehaviours like fsync discarding data. This turned into a mini hackathon that I and Mariusz attempted in the evening by trying to hook the syscall with dtrace and alter it's behaviour (unsuccessfully at that time but we had other ideas later on).

After the unconference, it was time to claim our badges. I loved the idea of handing them out at the Royal Oak pub. We were in I think a pretty unique situation, we basically raided the conference - not knowing anyone in the community. On top of that PGCon is known as the most developer heavy PostgreSQL conferences so those attending usually knew each other already. Having the badge handout in a social setting made it easier to start some casual conversations. We even met an OpenBSD user among the conference attenders.

Off to Day 1. Short and to the point opening session by Dan, kicked off the regular talks.

Our track choice was as follows:

- Challenges of Concurrent DDL, Robert Haas
- Transactions, Thomas Munro
- An Introduction to Memory Contexts, Chris Travers
- Intro to Postgres Planner Hacking, Melanie Plageman
- Pluggable Table Storage in PostgreSQL, Andres Freund
- Lightning talks

I found transactions, introduction to memory contexts and the intro to planner hacking as the most interesting talks for someone intending to start hacking on PostgreSQL itself. Especially the planner hacking talk. I basically had a specific query that was not using an optimization that should have been obvious and was hoping to approach someone during PGCon and see how hard it would be to implement that myself after some guidance. The talk covered exactly that plus discussed when & why you might not want to do so. I approached the speaker after the talk and we spent some time hacking on my use case, it won't likely result in any diffs to PostgreSQL as I'm pretty much unable to reproduce it on a newer version (which is good).

The Oracle lightning talk was hilarious. Do watch the recording when it pops up online.


This was also the evening of the social event. The conference organized catering in a pub. Very nice location and the food was simply amazing. This is where we finally had a chance to speak with most of the people attending and learn more about the PostgreSQL community itself.

I was extremely surprised to learn that Postgres has around 30 committers. I expected much higher numbers considering it's popularity. We learned that most committers are hired full time to work on PostgreSQL and it's pretty hard to do significant meaningful work for the project without very large commitments in time and resources. That's just the nature of the project and it did show in the scale of topics that the unconference talks covered - they were usually large scale and well planned projects like direct IO, new storage etc. If I understood correctly, the expected time frame for obtaining a commit bit varies between months and years and that's mostly because a significant contribution is non-trivial to achieve.

It was great to speak with experienced people. Learn what workloads are usual for PostgreSQL and by extension where problems in our deployments are caused by hardware and where we are hitting design limitations or are doing something suboptimal ourselves. We were given very good suggestions to try for some of specific use cases we had that I described in casual talks. My awareness of other index types like RUM and BRIN have been raised significantly.

![Picture showing a totem reaching across two floors to the ceiling. Totems were much taller than I ever expected, surprisingly also not round but a halved, hollowed trunk](/images/IMG_20190528_141427-s.jpg)
*Totems were much taller than I ever expected, surprisingly also not round but a halved, hollowed trunk*

Day 2. I really wanted to attend "Viva, the NoSQL Postgres!". It's about jsonpath - a query language for json and presented by Alexander Korotkov and Oleg Bartunov. Both (among other things) are the GIN index implementation authors which our talk is about. We also plan to use the PostgreSQL json features at work - essentially I have been reading a lot of their code and articles in the past few months. I did have a chance to speak with Oleg and learned about [RUM indexes](https://github.com/postgrespro/rum) which I definitely need to try.

Unfortunately we were not able to attend their talk as ours was scheduled at the same time! Seeing that I was pretty sure our room would be pretty much empty with everyone learning about the new cool feature that only days ago fronted on hackernews. We came in early, about an hour before the talk started and there were 3 people in the room. We started preparing and fortunately the room filled up almost fully before we started. I was pretty stressed speaking about PostgreSQL internals to a bunch of people who most likely know more about the code base. In general I think we did good and conveyed the message we wanted. One person in the audience was on BSDCan and stayed for PGCon just to attend our talk. We ended up grabbing some beers with him later that evening.

!['namaxsala reflects on human relationships with the natural world. The sculpture depicts the artists grandfather helping a wolf cross treacherous waters by boat.](/images/IMG_20190528_141955-s.jpg)
*'namaxsala reflects on human relationships with the natural world. The sculpture depicts the artists grandfather helping a wolf cross treacherous waters by boat.*

For the remaining talks we picked the following to attend:

- Transparent data at rest encryption in PostgreSQL, Masahiko Sawada
- Pluggable table access methods (pluggable storage), Pankaj Kapoor
- Advanced Authentication, Stephen Frost
- Toward Implementing Incremental View Maintenance on PostgreSQL, Yugo Nagata
- Introducing PostgreSQL SQL parser through an experience of reusing it in other applications, Bo Peng

There were two encryption related projects being talked about during PGCon. Both focused on storage encryption (one on tablespace level), we were hoping for a bit more fine grained control like selective column based encryption and access control for that (ie. who can encrypt & who can decrypt that column). I guess those storage based features that were presented are nice if you are running Postgres on a system not offering full disk encryption.

I liked the advanced authentication talk by Stephen Frost. It was a very nice overview of what Postgres is capable of. I also took some of his time after the talk and bounced one or two crazy ideas off him (unrelated to the talk) like disabling PostgreSQL parser at runtime to prevent SQL injections (by allowing only prepared statements which is viable if you use a library like [anosql](https://github.com/honza/anosql)).

The SQL query parser talk was also extremely valuable. I actually did something very similar for work when we wanted to do some additional lint checks against our schema as a post-commit hook. Things like detecting most likely missing indexes on foreign-keys. Bo Peng gave a very good run-down of the query parser, it's structure and the intermediate AST generated by the parser. This would have been a great time saver if I heard this talk before working on my task.

![Falls inside Ottawa, the river they fall into is magnificent](/images/IMG_20190601_130401-s.jpg)
*Falls inside Ottawa city, the river they fall into is magnificent*

The conference ended with an auction, on the very next day we still had a few hours to go around the city. Lots of buildings (like embassies) were open for tourists that day - unfortunately we did not have that much time. There were still many places to see and things to do and that is a very good motivation to be there again for the next conference. Be it BSDCan, PGCon or both of them.

I would like to thank the [@PGCon](https://twitter.com/PGCon) conference organizers & the [sponsors](https://www.pgcon.org/2019/sponsors.php) for covering my travel & accommodation. I would also like to thank my employer [Fudo Security](https://fudosecurity.com) for making hacking on PostgreSQL and BSDs my day job and considering my PGCon attendance as part of it.