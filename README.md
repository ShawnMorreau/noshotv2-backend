I went through a lot of https://quii.gitbook.io/learn-go-with-tests/ (https://github.com/ShawnMorreau/tddGolang) and picked up on TDD and creating testable code. The initial https://github.com/ShawnMorreau/no-shot-backend was hacked together for the most part, with code that wasn't very testable. It was a small achievement because I had never used websockets before and I was able to get a server up and running that allowed me to play the game with friends with the odd bug here and there. This is a more methodical approach to creating the backend for my game No Shot! I am trying my best to stick to TDD to ensure I'm testing my code. I am also creating all of the game logic before attaching a websocket this time as I have plans to expand the game further.

TODOS
Milestones 1: 
    - Create a refactored version of https://github.com/ShawnMorreau/no-shot-backend while following TDD as much as possible
    - separated tests, gamelogic, and main.go 
    - I did a bad thing. I started writing a bunch of code again without writing the tests. A lot of code changed meaning there's a good chance my tests don't pass anymore. On the plus side - I have the game working as it did before with much more managable code, meaning adding bots will be much easier than it would have been before. My main focus now will be to ensure I am testing the things that for sure need to be tested. Maybe add in some integration tests to ensure that all my components are working well together on the backend.  

   
I have ran into a point where I can either keep building this project out and make it a production grade game or stay where I am. The purpose of this project was to be able to play the card game "Red Flags" with my girlfriend. The game was sold out so I made this and we are able to play. It has a few known bugs as of right now and is missing logic, but it serves it's purpose. Given that there's tons of edge cases ahead, I believe it will be more productive for me to move on to better things and maybe come back to this at a later time. 

Known bugs: 
    - Enters closing state after no events have been sent for a period of time. I though I had this fixed but it came up again.
    - closing the browser mid game is a no no. Especially with bots
    
