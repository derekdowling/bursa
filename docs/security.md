# Security Model and Vision

## Overview
First, we need to have some context of who the user is as this drives api
requests as well as potentially isolating local storage to that specific user's
computer.

There might be multiple people using a computer. Someone at a house party
using your computer should have to log in in order to be able to decrypt
the keys in your browser's storage.

Using a public computer is obviously a bad idea, plus we could only send them
the private key, which we've agreed not to. Otherwise, the user would have to
import his keys into the public computer - obviously ill-advised.

## Goals

### Independence of Client App and Server
In order to preserve a user's rights to the money they own, we should strongly
consider the notion that a user should be able to log in to the client application
independently of the server. E.g. if only to retrieve their private keys with the
intention of migrating.

Failing to make a payment on their app subscription should not render the user
helpless - thousands of potential dollars vs $20 / month is, simply, a reprisal.

Keep in mind that access to their keys isn't, given our security scheme,
something they could phone us up for.  We don't have 2 of the 3 keys that were
generated for a given wallet.

At the very least, if a user is locked out of our app, they should be well
informed about the dangers of losing access to their keys. Instead of solving
this technically, we can aggressively inform our users of how important it is to
secure their keys.

LocalStorage is browser specific. I use three different browsers. I can delete
caches and entire browsers on a whim. I don't think people value the data in
their browsers because the expectation is that it's all in the cloud and can be
retrieved on a whim.

#### Hypothetical: allow for client / server authentication to fall out of sync.
Suppose we guarantee that a user can always access the app, whether they have an
account, or internet access.

In doing so, we open up a new problem: synchronizing credentials between the
client and server. If the user changes their password locally, it makes sense to
update it remotely. An option, that I think suffers from poor usability, is to
keep local and remote auth credentials separate. E.g. contextual password prompts
based on whatever zone/endpoint/view the user was trying to access.

There's a few scenarios here:

 1. Credentials are updated on the client and server.
 2. Credentials are not updated on the client or the server.
 3. Credentials are updated on the client, not on the server
    We re-encrypt local data, but the user is no longer able to interact
    with our server because as far as we know it's someone entering a wrong
    password, not that the password changed.
 4. Credentials fail to update on client, but succeed on the server.

It may be tricky to accurately identify which problem we're having. Once we now
which case we're dealing with, we'd need to initiate some conflict resolution.

Using point 3 as an example. A simple mechanism might to require that the user
first sign in remotely using their old credentials.

## Scenarios

### Basic Wallet Flow

The user creates an account.

 1. Bursa Account
 2. Local Storage

The user creates a wallet.

 1. The wallet must be secured. We keep it encrypted until it's needed, storing
    it encrypted according to the user's pass phrase, in some form of browser storage.
 2. With GC'd languages like js, what are the best practices for limiting exposure?
    We don't want things like an accidental console.log spamming output. The LoginAction
    for example, would have a username and pass provided as context.
 3. The cold storage wallet is not stored, it is simply presented. The user must
    download it and put it somewhere safe.
 4. The server is provided with the server wallet.

### The user returns to Bursa.
