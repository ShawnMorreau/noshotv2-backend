I went through a lot of https://quii.gitbook.io/learn-go-with-tests/ (https://github.com/ShawnMorreau/tddGolang) and picked up on TDD and creating testable code. The initial https://github.com/ShawnMorreau/no-shot-backend was hacked together for the most part, with code that wasn't very testable. It was a small achievement because I had never used websockets before and I was able to get a server up and running that allowed me to play the game with friends with the odd bug here and there. This is a more methodical approach to creating the backend for my game No Shot! I am trying my best to stick to TDD to ensure I'm testing my code. I am also creating all of the game logic before attaching a websocket this time as I have plans to expand the game further.

Milestone 2:
    - create separate lobbies so not every player joins into the same one everytime 
    - add in bots so there's no required amount of people to play with
Lifetime:
    - Lobby edgecases -> ensuring a lobby is "killed" if no actions are occuring.... (no human players in lobby, sitting in lobby for prolongued amount of time without adding anyone, etc)
    - Frontend is rough... it has structure but doesn't look very plesent 
Completed
Milestone 1: 
    - Create a refactored version of https://github.com/ShawnMorreau/no-shot-backend while following TDD as much as possible
    - separated tests, gamelogic, and main.go 
    - I did a bad thing. I started writing a bunch of code again without writing the tests. A lot of code changed meaning there's a good chance my tests don't pass anymore. On the plus side - I have the game working as it did before with much more managable code, meaning adding bots will be much easier than it would have been before. My main focus now will be to ensure I am testing the things that for sure need to be tested. Maybe add in some integration tests to ensure that all my components are working well together on the backend.  
    - Tests fixed and working now



   