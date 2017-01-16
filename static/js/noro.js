class Noro {
    constructor(){
    	this.chatRoomId = -1;
    	this.chatRoomWS = undefined;
    }
    getFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
    sweepChatRoom(){
    	if ($("user-dashboard-function-input").attr("display") != "none"){
    		$("#user-dashboard-function-input").fadeOut(500);
    	};
        $("#user-dashboard-function-input-text").val("");
        $("#main-content-left").html("");
    }
}

var noro = new Noro();
