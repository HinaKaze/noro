class Noro {
    constructor() {
        this.chatRoomId = -1;
        this.chatRoomWS = undefined;
    }
    getFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
    showChatRoom() {
        $("#user-dashboard-function-input").fadeIn(500);
        $(".main-header").css("display", "none");
    }
    sweepChatRoom() {
        if ($("user-dashboard-function-input").attr("display") != "none") {
            $("#user-dashboard-function-input").fadeOut(500);
        };
        $("#user-dashboard-function-input-text").val("");
        $("#main-content-left").html("");
        $(".main-header").css("display", "block");
    }
    changeMainBackground(imgurl) {
        $("#main-content").css("background-image", 'url("' + imgurl + '")');
        $("#main-content").css("background-position", "center");
        $("#main-content").css("background-repeat", "repeat-y");
    }
    clearMainBackground() {
        $("#main-content").css("background-image", "none");
    }
}

var noro = new Noro();

//自适应
(function(doc, win) {
    var docEl = doc.documentElement,
        resizeEvt = 'orientationchange' in window ? 'orientationchange' : 'resize',
        recalc = function() {
            var clientWidth = docEl.clientWidth;
            if (!clientWidth) return;
            if (clientWidth >= 1920) {
                docEl.style.fontSize = '100px';
                console.log("1", docEl.style.fontSize)
            } else {
                docEl.style.fontSize = 100 * (clientWidth / 1920) + 'px';
                console.log("2", docEl.style.fontSize)
            }
        };

    if (!doc.addEventListener) return;
    win.addEventListener(resizeEvt, recalc, false);
    doc.addEventListener('DOMContentLoaded', recalc, false);
})(document, window);
