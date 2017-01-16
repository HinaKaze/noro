class Noro {
    getFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
    constructor(){
    	this.chatRoomId = -1;
    	this.chatRoomWS = undefined;
    }
}

var noro = new Noro();
