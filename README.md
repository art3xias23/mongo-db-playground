# mongo-db-playground
- going through mongo installation and local setup
- simple queries with mongocompass and mongosh
- simple express app with the mongo driver

A fun thing I found was that since I was running nodemon on wsl windows terminal.
It could not find the 127.0.0.1 interface to connect to mongodb.
At that point I had two choices to change the C:\Program Files\MongoDB\Server\7.0\bin\mongod.cfg  blindip to

1. 0.0.0.0 which makes it accessible by all interfaces (insecure)
2. run ipconfig.exe
   a. get the IPv4 Address from the Ethernet adapter vEtherner(WSL)
   b. replce the blindip with it
