var PROTO_PATH = __dirname + '/../play.proto';

var grpc = require('grpc');
var play_proto = grpc.load(PROTO_PATH).play;
var client = new play_proto.Turn('localhost:50051',
      grpc.credentials.createInsecure());

//get_step returns the step made by the player
/*function get_step(){
  return {x: 0, y: 0, color: 1};
}

//send player step to the server
function main() {
  var client = new play_proto.Turn('localhost:50051',
  grpc.credentials.createInsecure());
  var move = get_step();
  client.playerMove(move, function(err, status) {

    console.log('Player places stone at coordinates: ', move.x, ", ", move.y);
  });
}*/

const readline = require('readline');

/*function getMove(){
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  var x;
  var y;
  rl.question("Enter coorindate x >", (answer) => {
    // x = parseInt(answer.toString());
    console.log("Player places stone at : ", answer);
    rl.close();
  })
  // rl.question("Enter coordinate y >", (answer) => {
  //   y = parseInt(answer.toString());
  //   rl.close();
  // })
  return {x: x, y: y, color: 1};
};*/

function main()
{
  let rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  rl.question("Enter 'B' for black, 'W' for white \n>", (answer) => {
    let humanColor;
    if (answer.toString().toUpperCase() == 'B'){
      humanColor = 1;
    } else {
      humanColor = 2;
    }
    let AIColor = humanColor%2 + 1;
    client.setPlayer({color: humanColor}, function(err, state){});
    //
    client.getAiPlayer({status: true}, function(err, player){
      console.log("The color of AI is ", player.color, "\n");
    
    //

    function prompt(){
      client.isNextPlayer({color: humanColor}, function(err, isNext){
        console.log(isNext.status)
        if (isNext.status){
          rl.question("Enter coordindate x \n>", (inputX) => {
            rl.question("Enter coordinate y \n>", (inputY) => {
              client.getMove({color: AIColor}, function(err, AImove){
                // if (AImove.player.color == 3) console.log("OpenGo has not made a move yet.")
                console.log("OpenGo places stone at coordinates: ", AImove.x, ", ", AImove.y, ".")
              var move = {x: parseInt(inputX.toString()), y: parseInt(inputY.toString()), 
                player: {color: humanColor}};
                client.setMove(move, function(err, state) {
                  client.updateNext({status: true}, function(err, state) {
                    client.isNextPlayer({color: humanColor}, function(err, next){
                      console.log("Next player is human? ", next.status);
                    })
                  });
                })
              })
              prompt();
            })
          })
        }
        else {
        var AITimer = setInterval(wait, 5000);
        function wait(){
          client.hasMoved({color: AIColor}, function(err, moved){
            if (moved.status) {
              client.getMove({color: AIColor}, function(err, AImove){
                console.log("OpenGo places stone at coordinates: ", AImove.x, ", ", AImove.y, ".");
                clearInterval(AITimer)
              })
            } else {
              console.log("Waiting for AI to make a move...")
            }
          })
        }
        prompt();
        //rl.close();
      }
      //prompt();
    })
    }
    setInterval(prompt, 5000);
      
    //}
  })
  })
  //TODO: set exit condition
};

main();
