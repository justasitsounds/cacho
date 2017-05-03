# Cacho - A dice game

[Cacho or Dudo is a popular dice game played in South America](https://en.wikipedia.org/wiki/Dudo)

Play is as follows:

1. Each player starts with 5 dice
2. The players roll one die each to determine who starts (highest score wins, in the case of draws, the highest scorers roll against each other until one player has the highest roll)
3. the starting player rolls their dice using a cup and peeks at her dice, without letting the other players know what they have rolled. The player then makes a prediction about all the dice that everyone has. IE: 'Three fours'.
4. The next player can challenge the previous player's prediction by saying either 'Dudo' (doubt) or 'Calza' (spot-on?). If they 'doubt' the prediction then all the dice in the group are counted to see how accurate the prediction was. If the previous player's prediction was too high, the previous player loses a die. If the challenger was wrong, then they lose a die. If the player asserts 'Calza' and there are exactly the predicted number of dice, then the challenger can reclaim a lost die.
5. One's (Aces) count as wild-cards 
6. If a player does not challenge they must raise the current 'bid' by increasing the number of dice of the previoulsy bid number (ie from "three sixes" to "four sixes") or by increasing the die number ("three fives" to "three sixes")
6. If a player is left with only one die then the rules cahnge slightly:
    a: the aces do not count as wild-cards
    b: the chosen number cannot change, ie. "two fours" cannot become "two sixes"


## Gameplay states:

1. registration:

Clients are invited to join a game or start a new game.
Ways to share a game with your friends:
i) start a game and share a unique short url with them to invite them. (Sharing via FB or other social media is a good idea for people that aren't in the same room)
ii) use a QR code + the mobile device's camera
iii) generate a human-friendly reddit-style ID: LonelyBlackPotatoSpiper
iv) allow geolocation on the devices and group users that are at the same or close location
v) allow the first user to create the room name

//
Client login
 Start or Join

 Start
    choose gamename
    wait for players...

 Join
    list local games
    choose game
    wait for player...

2. startgame
players sorted randomly
everyone rolls
//first player then makes bid
currentBid = startRound(firstplayer)

startRound(
    if anyplayer.dieCount() == 1{
        lastRound = true
        scoringStrategy = scoringFunc(lastRound) func(alldice, currentbid)
        biddingRules = biddingRules(lastround) func(currentPlayer)
    }

    for each player{
        if not first player{
            //userAction:challenge or bid
            if challenge{
                //dudo or calza
                if dudo{
                    if scoringStrategy(alldice.score < currentbid.score){
                        previousPlayer.LoseADie()
                        startRound(currentPlayer)
                    }else{
                        currentPlayer.LoseADie()
                        startRound(currentPlayer)
                    }
                }else{
                    if calzo && scoringStrategy({alldice.score == curentbid.score}){
                        currentPlayer.GainADie()
                        startRound(currentPlayer)
                    }
                }
            }else{
                //bid
                currentBid = biddingRules.IncreaseBid(currentPlayer,currentBid)
            }
        }
    }
)

