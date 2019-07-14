# Sesh - A Web Server Session Package

## Overview

There are three main concepts in this package:

- **`SessionInterface`** is responsible for maintaining session data for a user. This package
contains a implementation called Session.

- **`StoreInterface`** facilitates the adding and retrieving of individual sessions. This package contains an in-memory implementation known as MemStore.

- **`CookieManager`** allows easy retrieval of a session from HTTP requests containing a session cookie along with setting the session cookie for a HTTP response after creating a new session.

## Usage

To get a session using a HTTP request:

    session, err := cookieManager.Get(r)
    if err != nil {
        // no session
    } 
    // proceed with session
    session.Attribute("age")

Creating and adding a session to a HTTP response is just as easy:

    session := NewSession(idByteLength, sessionTimeout)
    cookieManager.Add(session, w)


Removing a session works in a similar manner:

    cookieManager.Remove(session, w)
