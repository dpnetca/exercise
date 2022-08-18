# CLI Task Manager
from gophercises course

Build CLI task manager that takes 3 commands, add, list and do., for this exercise we will use a 3rd party command line tool (cobra because that is what the course used) but there are lots of options shown here: https://github.com/avelino/awesome-go#command-line

For DB will use BoltDB for this, as that is what the course is using. However, this is not likely something that I will ever use again as the git repo is archived with no activity since July 2017.  There is a fork of the original bolt library (bbolt) that is maintained that could be worth considering (I am actually going to use the forked libary for this build). Bolt is a Key/value pair DB that is good for high read/low write database.
