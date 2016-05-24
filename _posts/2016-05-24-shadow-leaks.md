---
layout: post
title: Shadow leaks
date: '2016-05-24T00:00:00.000-02:00'
author: Adam Wołk
image: /images/Bugbusters.jpg
tags:
- privacy
---

The recent [linkedin password leak](https://www.troyhunt.com/observations-and-thoughts-on-the-linkedin-data-breach/) is a huge deal. The data was obviously leaked in 2012 but it's [';--have i been pwned?](https://haveibeenpwned.com/) that notified me about it and not linkedin.

I did change my password more than once since 2012. Is that enough reason for linkedin to not inform me about it? What if I removed my account completely? These days I am using a [keepassx](https://www.keepassx.org/) for password management and generate unique passwords for each service I use. Not sure though what my password was on linkedin in 2012. This is a none issue but there are more alarming things that grew on me when I reviewed the sites I have an account on.

Each and every internet service that you create an account on, that stores data about you is a ticking bomb. Even when your passwords are random and unique what does a service get? Your email? Name, shipping address, personal messages/data, phone number heck maybe even your CC card details? You might never even know that some of that data leaked out. I decided to close down some accounts that I no longer use & see the need for.

First off I'm happy that some sites allow me to remove my account. Yammer, Assembla and a couple of others were a painless process. It's somewhat tricky to locate the option as it tends to be tucked away in the least visited portion of the site but at least it's there.

Not all sites are the same though. Several of them, like the pivotal tracker force you to do that extra step and emailing tech support to have your account removed. That's a minor annoyance but at least they mention it in one of their help pages. They were also pretty prompt at executing that change without making me jump any additional hoops.

There are uglier sites though. Those that you or your wife used as a 'one off' years ago for that specific purchase. They tend to have incorrect TLS certificates set up (if at all), no way to remove existing data (billing address is a required field you can't remove it) and no way to delete the account. Sometimes even finding a contact email is an expedition in it's own. One site I ordered something once from in 2011 changed owners 3 times by now. Your only hope is to overwrite your personal data with some garbage (the ones they allow you to change, but usually their validation sucks) and hope that someone will react to your 'remove my account' support email.

Seems like a lot of leg work but doable or is it? Does that 'delete my account' option really work? Maybe someone just marked a 'deleted' flag on my entry in the users table but my personal data remains there intact. There is absolutely NO WAY to tell that. This is the real problem we are facing with the web today. Facebook, Google and even the small shops are either too invested in keeping every tiny bit of info about you or too incompetent to do a good job of cleaning up things that could cause problems later down the road, not to mention being able to inform you of breaches.

Here is a small example from my Pivotal Tracker account removal:

![My helpful screenshot](/images/pivotal-removal-emails.png)

I asked for account removal. Josh replied at 6:34 PM:

```As requested, I've removed your 'account@domain' login.```

Apparently my data is still somewhere at Pivotal since I got marketing email from them at 8:39 PM. After my account was removed (7:02 PM timestamp marks a reply for some additional feedback I gave after account closure).

I trust Josh is not malicious and that Pivotal respects those requests. Probably the newlestter is a separate system and Josh doesn't have removing my data from there on his checklist. I replied back and asked to be removed from any of their systems. Waiting to see how this unfolds. This is not a stab at Pivotal. I like their service and it was usefull when I still had a use case for them, they just happened to serve as a perfect example for today.

What else could be missing from a checklist like this?

 - Marketing systems
 - Analytics systems
 - Backups
 - Test servers
 - Developer machines with partial/full production data dumps
 - Stale access tokens for third party integration with Google, Facebook etc.
 - Cache
 - Data print outs
 - much much more...

If that data ever leaks will someone feel obliged to inform you? Will they be even able to do so? Will you remember what password you used when you still had an account there or what type of data was there? For many big companies you are probably just a shadow record marked as deleted waiting to be leaked out. Those few that deleted your data probably left a few ghots lingering around.

It took 4 years since 2012 to be informed by a third party that my account at linkedin was compromised. I can't even imagine how many breaches we are not aware of. Think twice before entering your data on a web form and creating a new account.