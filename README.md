# slaMDown
 A simple, local markdown presenter. I named this slaMDown out of a bit of frustration with not 
 being able to quickly render local markdown from a directory and step through all of those files
 quickly. There are solutions to this, but I took it as an opportunity to build a small tool for
 myself.

 ## TODO:
 - [x] create basic cli that will accept filepath(s)
 - [x] given a file path, list back the markdown files that we will present
 - [ ] given a list of markdown files, create web server and serve those files
    - [x] first spin up a web server
    - [ ] spin up pages for each found md with a link from there file path relative to the serarch directory?
    - [ ] clean up the list of pages and display a proper directory structure?

## Future Thoughts
- Present rendered markdown in the terminal; TUI?