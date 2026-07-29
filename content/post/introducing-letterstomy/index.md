+++
title = "Introducing LettersToMy: A Private Family Time Capsule Built for Apple Devices"
date = "2026-07-29T10:37:00-06:00"
lastmod = "2026-07-29T10:37:00-06:00"
description = "LettersToMy is a new privacy-first family time capsule for writing letters, preserving memories, collaborating with loved ones, and delivering them to children in the future."
summary = "Meet LettersToMy, a new Apple-first app for preserving letters, photos, videos, audio, and family memories for the people who matter most."
slug = "introducing-letterstomy"
draft = false
categories = ["Technology", "Apps", "Apple"]
tags = ["featured", "LettersToMy", "iOS", "macOS", "SwiftUI", "CloudKit", "TestFlight", "App Development"]
featureImage = "feature.webp"
featureImageAlt = "LettersToMy app icon beside the words Introducing LettersToMy"
featureImageCap = "LettersToMy is entering beta testing on iPhone, iPad, and Mac."
thumbnail = "feature.webp"
shareImage = "share.webp"
figurePositionShow = true
codeLineNumbers = false
codeMaxLines = 20
toc = true
usePageBundles = true
canonicalURL = "https://techrelay.xyz/post/introducing-letterstomy/"
+++

Some applications begin with a business plan. **LettersToMy began with a much more personal question: how can Melissa and I preserve the things we want our daughter to hear years from now?**

Not just photographs scattered across camera rolls. Not a social-media timeline controlled by an advertising company. I wanted a private place for letters, voice messages, videos, milestones, family stories, and the little thoughts that are easy to forget as life moves forward.

That idea became **LettersToMy**, a new Apple-first family time capsule currently entering TestFlight beta testing.

## More Than a Digital Baby Book

A traditional baby book records what already happened. LettersToMy is also about what comes next.

Parents and trusted family members can create content now and choose when it should become available. A letter might unlock:

- On a specific calendar date
- On a future birthday, such as age 5, 10, 16, or 18
- When a meaningful life event happens
- When a parent decides the time is right

That makes it possible to write for moments that have not arrived yet: the first day of school, graduation, a wedding, becoming a parent, navigating heartbreak, or simply a future birthday when a few words from home may matter more than usual.

The app is not limited to daughters, either. A family archive can be organized for a **son, daughter, grandson, granddaughter, or another important recipient**.

## Built for iPhone, iPad, and Mac First

I chose to build LettersToMy as a native Apple application instead of starting with a generic cross-platform wrapper.

The current app uses:

- **SwiftUI** for the shared iPhone, iPad, and Mac experience
- **Core Data** with `NSPersistentCloudKitContainer`
- Separate **private and shared CloudKit stores**
- Native iCloud collaboration and invitation handling
- A portable Swift package for unlock rules and permission logic

The native approach gives the app proper Apple platform behavior while allowing most of the interface and business logic to be shared between iOS, iPadOS, and macOS.

A deeply integrated web application is planned next. It will use the same CloudKit-backed data rather than creating a disconnected second archive. Android is planned after the Apple and web data contracts are stable.

## Private iCloud Storage Instead of Another Social Network

Family memories should not exist to feed an engagement algorithm.

LettersToMy does not require a proprietary application server to store the Apple version of your archive. Content is stored locally and synchronized through the signed-in user's iCloud account using private and shared CloudKit databases.

That architecture is intended to provide:

- Private storage tied to the family's Apple accounts
- Synchronization across iPhone, iPad, and Mac
- Native sharing with invited family members
- Clear ownership of the original archive
- A foundation for controlled recipient access later

CloudKit is not the entire long-term preservation strategy. Before a public release, the project also needs encrypted exports, restore tools, recovery contacts, successor access, and a documented path for retrieving an archive even if the application eventually disappears.

An app designed to hold decades of memories cannot treat recovery as an afterthought.

## Collaboration Without Giving Everyone the Keys

LettersToMy is designed for more than one author.

A spouse might need full administrative access. Grandparents may be allowed to create letters only within their side of the family. Another relative might receive read-only access. A child should eventually receive only their own unlocked content—not every sealed letter waiting in the archive.

The current collaboration model includes roles for:

- Owner
- Parent or administrator
- Family organizer
- Contributor
- Viewer
- Recipient

Access can be scoped to:

- The complete archive
- A family side
- A folder
- A specific recipient

The app also supports explicit permission grants and denials. That makes it possible to distinguish between actions such as creating content, editing only your own letters, managing folders, inviting another family member, releasing a life-event letter, or exporting the archive.

CloudKit shares are separated along important security boundaries instead of placing the entire family into one enormous shared database space.

## Family Sides and Folders

Family structures are rarely simple, so the archive should not assume they are.

LettersToMy currently creates several starting areas:

- Parents
- Maternal Family
- Paternal Family
- Chosen Family

Custom family sides and folders can be added as needed. Letters can then be assigned to a recipient, family side, and folder.

This organization is not just cosmetic. It also determines which collaborators can see or contribute to particular content.

## Keeping Sealed Letters Away From Recipients

The recipient experience requires a stronger boundary than simply hiding a button in the interface.

Children are not added to the same CloudKit share that contains every sealed master letter. The planned recipient-delivery model uses a separate inbox for each recipient. When content unlocks, an appropriate delivery record is placed into that inbox.

The goal is to ensure recipients cannot inspect:

- Other children's content
- Sealed letters
- Hidden attachments
- Future delivery metadata
- Administrative archive records

Recipient inboxes are intended to be read-only by default, with optional replies that do not alter the original letter.

## What Works in the Current Beta

The current iPhone, iPad, and Mac builds include the foundation for:

- Creating, editing, favoriting, sealing, and deleting letters
- Fixed-date, birthday-age, and life-event unlock rules
- Photos, videos, audio files, and other attachments
- Recipient preview mode
- Family profiles
- Family sides and folders
- Collaboration roles and permission scopes
- Real CloudKit share creation and invitation acceptance
- Private and shared iCloud-backed stores
- Search and timeline views
- Synchronization across Apple devices

The project has automated tests for core unlock and collaboration rules, along with continuous builds for the iOS simulator and macOS targets.

This is still beta software. Media handling, recipient delivery, notifications, recovery, collaboration lifecycle details, and portions of the interface will continue to evolve.

## What I Need Beta Testers to Try

The most useful testing is not simply opening the app and confirming that it launches.

Testers should try to break the workflows that matter:

1. Create a letter, attach media, and seal it.
2. Schedule letters using each unlock method.
3. Organize content under different recipients and family folders.
4. Invite another iCloud user and test both read-only and contribution access.
5. Make changes on one Apple device and confirm they appear on another.
6. Work briefly offline, reconnect, and check synchronization.
7. Turn on recipient preview and verify sealed content remains hidden.
8. Confirm collaborators cannot edit or delete content outside their permissions.
9. Relaunch the app and verify letters and attachments remain intact.

The most important reports involve crashes, duplicates, missing attachments, incorrect permissions, invitation failures, synchronization problems, or any situation where sealed content appears for the wrong user.

Testflight link here: [https://testflight.apple.com/join/7YUj661j](https://testflight.apple.com/join/7YUj661j)

## Why I Am Building It in Public

LettersToMy is personal, but the engineering problems are broadly interesting.

The project combines native application design, CloudKit collaboration, long-lived data, role-based permissions, media storage, offline synchronization, and eventually a web client that must safely access the same archive.

I plan to share more of that development process here on Tech Relay, including:

- Building private and shared CloudKit stores
- Designing scoped family collaboration
- Handling recipient delivery securely
- TestFlight and App Store deployment
- Creating a CloudKit-integrated web client
- Designing encrypted archive export and recovery

## The Long-Term Goal

The goal is not merely to ship another journaling application.

I want LettersToMy to become a trusted family archive: a place where someone can receive the right words at the right moment, even if those words were written many years earlier.

Software changes quickly. Families, memories, and the things we wish we had said deserve a little more permanence.

**LettersToMy is now in its early Apple beta phase, and this is only the beginning.**

---

## Frequently Asked Questions

### What is LettersToMy?

LettersToMy is a private family time-capsule app for writing letters, saving media and memories, collaborating with trusted relatives, and making content available to a recipient in the future.

### Which devices does it support?

The current app is being developed first for iPhone, iPad, and Mac. A CloudKit-integrated web app is planned, followed later by Android.

### Is LettersToMy a social network?

No. It is designed as a private family archive rather than a public feed or engagement-driven social platform.

### Where is family content stored?

The Apple-first version stores content locally and synchronizes it through the archive owner's iCloud account using private and shared CloudKit databases.

### Can grandparents and other relatives contribute?

Yes. Invited family members can receive roles and scopes that determine whether they can view, create, edit, organize, or manage specific areas of the archive.

### Can children see letters before they unlock?

The recipient architecture is designed to keep sealed master records out of recipient shares. Recipients should receive only content that has been delivered to their own inbox.

### Is the app publicly available?

LettersToMy is currently in beta development and TestFlight testing. Features and data structures may change before a public App Store release.

Testflight is available [here](https://testflight.apple.com/join/7YUj661j)
