Session Cookie Authentication
=============================

Session cookies are cookies that last for a session. This happens when you launch a web app/website and ends when you leave that website or close the browser window.

Session cookies contain information such as your website settings, preferred language, login status, location, and also can store email address, name, other personal information (if you provide it).

In this tutorial, explain how to authenticate users using session cookies in a server application in order to detect who they are that across our http methods and routes.

How session cookie works in authentication process?
---------------------------------------------------

- Login

    User tries to login and new session token will be generated. That session will be stored in server.
    User will obtain a set of cookie as session token.

- Refresh
    
    This will be called after the expiry time of a session is expired. This is needed to refresh the token to keep app stays alive without re-login. 

- Logout

    The session will be removed from storage as well as the users client (no session info of that user in server).