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
    showChatLobby() {
        $("#musubi-omamori").text("分类");
        $(window).scroll(function() {
            if (document.body.scrollTop > $("#main-header").height()) {
                $("#musubi-omamori").text("回首");
            } else {
                $("#musubi-omamori").text("分类");
            }
        });
    }
    sweepChatLobby(){
        $("#musubi-omamori").text("回首");
        $(window).unbind("scroll");
    }
    showChatRoom() {
        $("#user-dashboard-function-input").fadeIn(500);
        $("#main-header").css("display", "none");
    }
    sweepChatRoom() {
        if ($("user-dashboard-function-input").attr("display") != "none") {
            $("#user-dashboard-function-input").fadeOut(500);
        };
        $("#user-dashboard-function-input-text").val("");
        $("#main-content-left").html("");
        $("#main-header").css("display", "block");
    }
    changeMainBackground(imgurl) {
        $("#main-content").css("background-image", 'url("' + imgurl + '")');
        $("#main-content").css("background-position", "center");
        $("#main-content").css("background-repeat", "repeat-y");
    }
    clearMainBackground() {
        $("#main-content").css("background-image", "none");
    }
    scrollTop() {
        $("html,body").animate({ scrollTop: 0 }, 500);
    }
    toggleBlur(element){
        if (element.css("filter") == "none"){
            element.css("filter","blur(4px)");
            element.css("-moz-filter","blur(4px)");
            element.css("-webkit-filter","blur(4px)");
            element.css("-o-filter","blur(4px)");
            element.css("-ms-filter","blur(4px)");
        }else{
            element.css("filter","none");
            element.css("-moz-filter","none");
            element.css("-webkit-filter","none");
            element.css("-o-filter","none");
            element.css("-ms-filter","none");
        }
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
            } else {
                docEl.style.fontSize = 100 * (clientWidth / 1920) + 'px';
            }
        };

    if (!doc.addEventListener) return;
    win.addEventListener(resizeEvt, recalc, false);
    doc.addEventListener('DOMContentLoaded', recalc, false);
})(document, window);
