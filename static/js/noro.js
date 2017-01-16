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
        this.clearMainBackground();
    }
    changeMainBackground(imgurl){
        $("#main-content").css("background-image","imgurl");
        $("#main-content").css("background-position","center");
        $("#main-content").css("background-repeat","repeat-y");
    }
    clearMainBackground(){
        $("#main-content").css("background-image","none");
        $("#main-content").css("background-position","none");
    }
}

var noro = new Noro();
