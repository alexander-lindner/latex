---
layout: blog
date: 2022-04-17
authors:
- name: Alexander Lindner
  link: https://alindner.org
  avatar: https://avatars.githubusercontent.com/u/25225552?v=4
route: history-of-textool
---

# The history of TeXtool and it's docker image

In 2016, I created my first latex document and was frustrated, how difficult it was (and still is) 
to install an up-to-date version of texlive and also install necessary packages.
Unfortunately, using windows makes it even worse.

So, the big docker fan I was at the time, I created a docker image with a full installed texlive copy and shared it later through
Docker Hub and Github.
This image did a good job as I also could use it at my company in the Gitlab CI or at my private CI to automatically push a build
document to my NextCloud.
I therefore created some bash scripts to simplify the installation.

When I had done a university seminar with a friend of mine some draw backs popped out:
* the size of the image
* missing watch command
* creating a project and utilising the docker image was a little mess because the entrypoint had to be overridden most of the time

So I decided to rework my image, however I wasn't motivated enough to really to it.
When I started my master thesis, I was finally motivated and because I wanted to try out `GO`,
I switched from by bash scripts to a go cli.

These changes were a lot of fun, so I decided to add more and more features, and now I spend some time from week to week 
and improve this project even further so everyone can use it with ease.
