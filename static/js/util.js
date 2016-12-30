class util {
    GetFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
}

var noro = new util();
