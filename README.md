# Navimix
A custom "API injection" proxy for Navidrome and Deemix. 

# About
This app is a sole developer project to learn about APIs, backends, and the GO programming language. I enjoy writing it in my free time, and hope to continue to expand its functionality

# Features
Navimix is written entirely in GO, and runs exclusively as a backend. Its purpose is to transparently download any (copyright-free) music and stream it on the fly via subsonic. Navimix essentially sits between a normal subsonic server and the client, proxying unchanged API functionality and modifying certain API requests.

Current features (in addition your standard subsonic server) include:

- Search any song outside of library (available on deezer) and have it playback via streaming
- Fetch cover art for above songs

Navimix acts as a translation integration layer between the deezer api, deemix, navidrome, and subsonic. It is designed to run on top of navidrome (or any compatible subsonic server) and add functionality. It is still very much under development, with many new features currently under development.

Feel free to fork the repository or open a thread with suggested features!

# A note on piracy
Navimix is a tool designed to interact with local music streaming services and augment your personal music collection with **public domain** content. 
It **is not meant to provide access to content you do not own or have permission to use**. Please ensure that you have permission to use any content you stream or download.