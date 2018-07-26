Date.prototype.format = function (format) {
    var args = {
        "M+": this.getMonth() + 1,
        "d+": this.getDate(),
        "h+": this.getHours(),
        "m+": this.getMinutes(),
        "s+": this.getSeconds(),
       "ms+": this.getMilliseconds(),
    };
    if (/(y+)/.test(format))
        format = format.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var i in args) {
        var n = args[i];
        if (new RegExp("(" + i + ")").test(format))
            format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? n : ("00" + n).substr(("" + n).length));
    }
    return format;
};


(function(exports) {
    var SimpleMap = exports.SimpleMap = function() {
            this.map = {};
            this.mapSize = 0;
        };

    SimpleMap.prototype.put = function(key, value) {
        var oldValue = this.map[key];
        this.map[key] = value;
        if(!oldValue) {
            this.mapSize++;
        }
        return(oldValue || value);
    };

    SimpleMap.prototype.get = function(key) {
        return this.map[key];
    };

    SimpleMap.prototype.remove = function(key) {
        var v = this.map[key];
        if(v) {
            delete this.map[key];
            this.mapSize--;
        };
        return v;
    };

    SimpleMap.prototype.size = function() {
        return this.mapSize;
    };

    SimpleMap.prototype.clear = function() {
        this.map = {};
        this.mapSize = 0;
    };

    SimpleMap.prototype.keySet = function() {
        var theKeySet = [];
        for(var i in this.map) {
            theKeySet.push(i);
        }
        return theKeySet;
    };

    SimpleMap.prototype.values = function() {
        var theValue = [];
        for(var i in this.map) {
            theValue.push(this.map[i]);
        }
        return theValue;
    };

    var CircleList = exports.CircleList = function(maxSize) {
            this.maxSize = (maxSize || 10);
            this.list = [];
            this.index = null;
        };

    CircleList.prototype.clear = function() {
        this.list = [];
        this.index = null;
    };

    CircleList.prototype.add = function(value) {
        if(null == this.index) {
            this.index = 0;
        }

        this.list[this.index++] = value;

        if(this.index == this.maxSize) {
            this.index = 0;
        }
    };

    CircleList.prototype.values = function() {
        var theValue = [];
        if(null != this.index) {
            if(this.list.length == this.maxSize) {
                for(var i = this.index; i < this.maxSize; i++) {
                    theValue.push(this.list[i]);
                }
            }

            for(var i = 0; i < this.index; i++) {
                theValue.push(this.list[i]);
            }
        }
        return theValue;
    };

    exports.EVENT_TYPE = {
        'LOGIN': 'LOGIN',
        'LOGOUT': 'LOGOUT',
        'SPEAK': 'SPEAK',
        'LIST_USER': 'LIST_USER',
        'ERROR': 'ERROR',
        'LIST_HISTORY': 'LIST_HISTORY'
    };

    // 服务端口
    exports.PORT = 9090;

    // 服务端口
    exports.HOST = "localhost";

    var parseMessageData = exports.parseMessageData = function(message) 
    {
        try {
            return JSON.parse(message);
        } catch(error) {
            // 收到了非正常格式的数据
            console.log('method:parseMessageData,error:' + error);
        }

        return null;
    }

    var getMsgData = exports.getMsgData = function(mData) 
    {
        if(mData && mData.data) 
        {
            return mData.data;
        }

        return null;
    }

    var log = exports.log = function(message) 
    {
        console.log(new Date().format("yyyy-MM-dd hh:mm:ss ms ") + message);
    }

})((function() {
    if(typeof exports === 'undefined') {
        window.Tool = {};
        return window.Tool;
    } else {
        return exports;
    }
})());
