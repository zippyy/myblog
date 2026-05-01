---
title: "Dank Diary"
date: 2026-04-28
draft: false
tags: ["ios", "cloud", "tech"]
---

## Introduction

Alright so I built a new app called **Dank Diary**.

This wasn’t originally meant to be a full project, it started as something simple just to track sessions and strains but like most things it turned into a full iOS app with a web component and proper sync.

Main goal was being able to log entries, attach images, and access everything from anywhere without depending on a single device.

---

## Overview

At a high level this is what it does:

- Log sessions with notes, ratings, strain info  
- Attach images  
- Sync data between devices  
- View everything from a web interface  

Nothing crazy feature wise, most of the work went into making sure it actually works reliably.

---

## Access

Web app is live here:  
https://dankdiary.xyz

TestFlight is here:  
https://testflight.apple.com/join/HZxEHcz6

---

## The Actual Problem

Text is easy.

You can throw that in pretty much anything and sync it without issues.

Images are where things start to break.

If you embed them directly your data gets bloated, sync slows down, and everything starts feeling sluggish especially on mobile.

So the main problem this project ended up solving was:

- How to handle media properly  
- How to sync that media between iOS and web  
- How to do it without blowing up storage or performance  

---

## Setup / Architecture

### iOS App

Native Swift app.

Local first data model so entries are instant and not dependent on network.

Background sync handles pushing changes up.

---

### Web App

Pretty simple for now.

Mainly read access and sharing.

Not trying to replicate the full mobile experience here, just make the data accessible.

---

### Sync + Auth

Using Sign in with Apple for auth.

Data lives in cloud storage and is tied to the user account.

Images are stored as assets instead of being embedded directly in records.

This keeps things smaller and makes fetching more predictable.

---

## Issues I Ran Into

Couple things that were more annoying than expected:

- Keeping data consistent between iOS and web  
- Making image uploads not feel slow  
- Avoiding unnecessary storage usage  
- Not exposing user data when sharing entries  

None of this is complicated individually but when you combine it all it adds up quick.

---

## Monetization

Keeping this simple.

Free version:
- Full functionality  
- Sync enabled  
- Ads in non critical screens  

Paid version:
- Removes ads  
- Will add more features later  

No ads in the editor. Not worth degrading the core experience.

---

## Why I Built This

This was less about the actual app idea and more about working through:

- Cross platform sync  
- Media handling  
- Keeping things lightweight  
- Building something that actually works outside of a demo  

Same type of problems I run into on client work just packaged into a standalone project.

---

## Next Steps

- Improve sharing between users  
- Add basic analytics on entries  
- Continue optimizing sync performance  

---

## Closing

It works.

It syncs.

It does what it needs to do.

Now it just needs iteration.

As always, So long and Thanks for all the Fish!
