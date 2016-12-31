class Noro {
    getFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
    constructor(){
    	this.lastRoomId = -1;
    	this.roomWS = undefined;
    }
}

var noro = new Noro();
