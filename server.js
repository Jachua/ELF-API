'use strict';

var express = require('express');
const app = express();
const http = require('http').Server(app);
const underscore = require('underscore');
const fs = require('fs');
var sprintf = require("sprintf-js").sprintf;
var vsprintf = require("sprintf-js").vsprintf;

const Tool = require('./tools');

const WebSocket = require('ws');
 
const wss = new WebSocket.Server({ port: 9090 });

var PROTO_PATH = __dirname + '/../play.proto';

var grpc = require('grpc');
var play_proto = grpc.load(PROTO_PATH).play;

wss.broadcast = function broadcast(data) 
{
    wss.clients.forEach(function each(client) 
    {
        if (client.readyState === WebSocket.OPEN) {
          client.send(data);
        }
    });
};

wss.broadcastEveryoneElse = function broadcast(ws, data) 
{
    wss.clients.forEach(function each(client) 
    {
        if (client !== ws && client.readyState === WebSocket.OPEN) {
            client.send(data);
        }
    });
};

var EVENT_TYPE = 
{
    'TEST'              : 'TEST',
    'REGISTER_USER'     : 'REGISTER_USER',
    'CREATE_ROOM'       : 'CREATE_ROOM',
    'GET_ROOM_LIST'     : 'GET_ROOM_LIST',
    'JOIN_ROOM'         : 'JOIN_ROOM',
    'PLAYER_JOIN_ROOM'  : 'PLAYER_JOIN_ROOM',
    'LEAVE_ROOM'        : 'LEAVE_ROOM',
    'UPDATE_BOARD'      : 'UPDATE_BOARD',
    'SAVE_DEAD_BOARD'   : 'SAVE_DEAD_BOARD',
    'AI_PLAY'           : 'AI_PLAY'
};

function getSendData(event, data)
{
    var  obj = {};
    obj["event"] = event;
    obj["data"] = null;
    if(data)
    {
        obj["data"] = data;
    }
    return obj;
}

wss.people = 0;
wss.on('connection', function connection(ws) 
{   
    wss.people++;
    Tool.log("connection people:" + wss.people);
    let obj = {};
    obj["data"] = "123123123";
    ws.send(JSON.stringify(obj));
    
    ws.on('open', function open(message)
    {
        console.log('open: %s', message);
    });

    ws.on('message', function message(message) 
    {
        var mData = Tool.parseMessageData(message);
        
        if(mData && mData.event)
        {
            var data = Tool.getMsgData(mData);
            if(data instanceof Object)
                Tool.log(sprintf("request type:%s data:%s", mData.event, JSON.stringify(data)));
            else
                Tool.log(sprintf("request type:%s data:%s", mData.event, data));
            switch(mData.event) 
            {
                case EVENT_TYPE.TEST:
                    wss.broadcast(JSON.stringify(getSendData("TEST", "Test Event")));
                    break;
                case EVENT_TYPE.REGISTER_USER:
                    let index = findPlayerIndexById(data.id);
                    if(index > 0) //改名
                    {
                        players[index].name = data.name; 
                    }
                    else //新建
                    {
                        let player = new Player();
                        player.id = data.id;
                        player.name = data.name;
                        players.push(player);
                    }
                    break;
                case EVENT_TYPE.CREATE_ROOM:
                {
                    let room = createRoom(data.roomName, data.roomSize, data.playerId);
                    roomSockets[room.roomId] = [];
                    ws.send(JSON.stringify(getSendData("CREATE_ROOM", room.roomId)));
                    break;
                }
                case EVENT_TYPE.GET_ROOM_LIST:
                    // 获取最近的聊天记录
                    ws.send(JSON.stringify(getSendData("GET_ROOM_LIST", rooms)));
                    break;

                case EVENT_TYPE.JOIN_ROOM:
                {
                    let mode = 2;
                    Tool.log('join room'+ data.roomId);
                    let roomSign = 'Room' + data.roomId;
                    let roomIndex = findRoomIndexById(data.roomId);
                    let room = rooms[roomIndex];
                    
                    roomSockets[room.roomId].push(ws);  
                    ws.roomSign = roomSign;
                    Tool.log(data.playerId + '加入了' + roomSign);
            
                    //加入
                    if(data.playerId == room.playerId1) //主机
                    {
                        //加入或者继续
                        let p1 = findPlayerIndexById(data.playerId);
                        room.playerInfo1 = findPlayerByIndex(p1);
            
                        mode = 2;
                    }
                    else //P2
                    {
                        if(room.playerId2 == "")//加入
                        {
                            room.playerId2 = data.playerId;
                            let p2 = findPlayerIndexById(data.playerId);
                            room.playerInfo2 = findPlayerByIndex(p2);
            
                            mode = 2;
                        }
                        else if(room.playerId2 != "" && data.playerId == room.playerId2)
                        {
                            //继续
                            let p2 = findPlayerIndexById(data.playerId);
                            room.playerInfo2 = findPlayerByIndex(p2);
            
                            mode = 2;
                        }
                        else
                        {
                            //观战
                            mode = 3;
                        }
                    }

                    ws.send(JSON.stringify(getSendData("JOIN_ROOM", {mode:mode, room:room})));
                    wss.broadcastRoomEveryoneElse(ws, room.roomId, JSON.stringify(getSendData("PLAYER_JOIN_ROOM", room)));
                    break;
                }
                case EVENT_TYPE.LEAVE_ROOM:
                {
                    let roomSign  = ws.roomSign ;
                    if(roomSign && roomSign != "")
                    {
                        wss.broadcastRoom(ws, "LEAVE_ROOM", null);
                        ws.roomSign = "";
                        let roomId = roomSign.substr(4);
                        let roomClients = roomSockets[roomId];
                        if(roomClients)
                        {
                            let index = -1;
                            for(var i = 0; i < roomClients.length; i++)
                            {
                                var client = roomClients[i];
                                if (client == ws) {
                                    index = i;
                                }
                            }
                            roomSockets[roomId].splice(index);
                        }
                    }
                    break;
                }   
                case EVENT_TYPE.UPDATE_BOARD:
                {
                    let roomId = data.roomId;
                    let op = Operation.jsonToObj(data.operation);
                    let roomIndex = findRoomIndexById(roomId);
                    let room = finRoomByIndex(roomIndex);
            
                    room.report.addOperation(op);
            
                    let roomSign  = ws.roomSign ;
                    Tool.log("roomSign" + roomSign);
                    if(roomSign && roomSign != "")
                    {
                        wss.broadcastRoomEveryoneElse(ws, roomSign.substr(4), JSON.stringify(getSendData("UPDATE_BOARD", op)));
                    }
                    break;
                }
                case EVENT_TYPE.SAVE_DEAD_BOARD:
                {
                    Tool.log('save dead board' + data.toString());

                    let boardName = data["boardName"];
                    if(!boardName || boardName == "")
                        return;
                    Tool.log(__dirname);
                    var fileName = __dirname + "/save/" + boardName + ".json";
                    fs.exists(fileName, function(exist) {
                        if(exist)
                        {
                            fs.unlink(fileName, function (err) {
                                if(err) throw err;
                                Tool.log('删除成功')
                            });
                        }
                        fs.open(fileName, "w",function (err) {
                            if(err)
                            {
                                ws.send(JSON.stringify(getSendData("SAVE_DEAD_BOARD", {"result":2})));
                                return;
                            }
                        });
            
                        fs.readFile(fileName, function(err,fileData)
                        {
                            if(err){
                                console.error(err);
                                ws.send(JSON.stringify(getSendData("SAVE_DEAD_BOARD", {"result":2})));
                            }
                            var str = JSON.stringify(data);
                            fs.writeFile(fileName, str,function(err)
                            {
                                if(err){
                                    console.error(err);
                                    ws.send(JSON.stringify(getSendData("SAVE_DEAD_BOARD", {"result":2})));
                                    return;
                                }
                                Tool.log('----------保存成功-------------');
                                ws.send(JSON.stringify(getSendData("SAVE_DEAD_BOARD", {"result":1})));
                            })
                        })
                    });
                    break;
                }
                case EVENT_TYPE.AI_PLAY:
                {
                    let op = Operation.jsonToObj(data.operation);
                    var client = new play_proto.Turn('localhost:50051',
                    grpc.credentials.createInsecure());
                    var move = {x: op.x, y: op.y, color: op.player}    
                    client.playerMove(move, function(err, status) {
                        ws.send(JSON.stringify(getSendData(EVENT_TYPE.AI_PLAY, {"result":1})));
                        // console.log('Player places stone at coordinates: ', move.X, ", ", move.Y);
                    });


                    break;
                }
                default:
                    break;
                }

                
        }
    });

    ws.on('close', function close(message) 
    {
        wss.people--;
        Tool.log("close:" + message + "people:" + wss.people);
    });
});

/////////////////////////////////////////////////////////////////

wss.broadcastRoom = function broadcast(roomId, data) 
{
    var roomClients = roomSockets[roomId];
    if(roomClients && roomClients.length > 0)
    {
        for(var i = 0; i < roomClients.length; i++)
        {
            var client = roomClients[i];
            if (client.readyState === WebSocket.OPEN) {
                client.send(data);
              }
        }
    }
};

wss.broadcastRoomEveryoneElse = function broadcast(ws, roomId, data) 
{
    var roomClients = roomSockets[roomId];
    if(roomClients && roomClients.length > 0)
    {
        for(var i = 0; i < roomClients.length; i++)
        {
            var client = roomClients[i];
            if (client !== ws && client.readyState === WebSocket.OPEN) {
                client.send(data);
              }
        }
    }
};


let MAX = 30;
let rooms = [];
let players = [];
let roomIndex = 0;
let roomSockets = {};

class Player
{
    constructor()
    {
        this.id = "";
        this.name = "";
    }
}

class Report
{
    constructor()
    {
        this.size = 0;
        this.start = 0;
        this.operations = [];
    }

    addOperation(oper)
    {
        this.operations.push(oper);
    }

    objToJson()
    {
        let res = {};
        res["size"] = this.size;
        res["start"] = this.start;
        res["operations"] = this.operations;
        return JSON.stringify(res);
    }

    static jsonToObj(json)
    {
        let report = new Report();
        report.size = json["size"];
        report.start = obj["start"];
        let operations = obj["operations"];
        for(let op in operations)
        {
            let operation = Operation.jsonToObj(op);
            report.operations.push(operation);
        }
    }
}

class Operation
{
    constructor(x, y, player, step, dead)
    {
        this.x = x;
        this.y = y;
        this.player = player; //1:hei 2 : bai
        this.step = step;
        this.dead = dead;
    }

    static jsonToObj(json)
    {
        let operation = new Operation();
        operation.x = json["_x"];
        operation.y = json["_y"];
        operation.player = json["_player"]
        operation.step = json["_step"];
        operation.dead = json["_dead"];
        return operation;
    }
}

class Room
{
    constructor()
    {
        this.roomId = 0;
        this.roomName = "";
        this.roomSize = 9;
        this.playerId1 = "";
        this.playerId2 = "";
        this.playerInfo1 = null;
        this.playerInfo2 = null;
        this.report = new Report(); //save
    }
};

function findPlayerIndexById(id) 
{
    let index = -1;
    for (var i = players.length - 1; i >= 0; i--) 
    {
        if(players[i].id == id)
        {
            index = i;
            break;
        }
    }
    return index;
}

function findPlayerIndexByName(name) 
{
    let index = -1;
    for (var i = players.length - 1; i >= 0; i--) 
    {
        if(players[i].name == name)
        {
            index = i;
            break;
        }
    }
    return index;
}

function findPlayerByIndex(index)
{
    if(index >= 0)
        return players[index];
    return null;
}

function findRoomIndexById(roomId) 
{
    let index = -1;
    for (var i = rooms.length - 1; i >= 0; i--) 
    {
        if(rooms[i].roomId == roomId)
        {
            index = i;
            break;
        }
    }
    return index;
}

function finRoomByIndex(index)
{
    if(index >= 0)
        return rooms[index];
    return null;
}

function createRoom(roomName, roomSize, playerId)
{
    let room = new Room();
    roomIndex++;
    room.roomId = roomIndex;
    room.roomName = roomName;
    room.roomSize = roomSize;
    room.playerId1 = playerId;
    rooms.push(room);
    return room;
}