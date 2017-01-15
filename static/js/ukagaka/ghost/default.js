var ghost = {
    data: {
        eat_times: 0, //吃了几次
        talkself_arr: [],
        faces: {
            '1': "face1.gif",
            '2': "face2.gif",
            '3': "face3.gif",
            '4': "face4.png",
        },
        menu: [
            //['shownotice', '显示公告'],
            //['chatTochuncai', '聊&nbsp;&nbsp;&nbsp;&nbsp;天'],
            //['eat', '吃 零 食'],
            ['transform', '变身'],
            //['meetparents', '见 家 长'],
            //['lifetimechuncai', '生存时间']
        ],
    },
    shell: {},
    init: function(shell) {
        //some ghost init...
        this.shell = shell
    },

    /**
     * 事件列表
     */


    /**
     * 显示公告
     */
    // shownotice: function() {
    //     this.getdata("getnotice");
    //     ghost.data.Ukagaka.setFace(1);
    //     this.shell.closeChuncaiMenu();
    // },

    /**
     * 变身
     */
    transform: function() {
        this.shell.setFace(4);
        this.shell.stopTalkSelf();
    },

    /**
     * 与春菜聊天
     */
    chat: function() {
        talk_html = '';
        talk_html += '  <div class="addinput">';
        talk_html += '      <div class="inp_l">';
        talk_html += '          <input class="talk" type="text" name="mastersay" value=""/>';
        talk_html += '          <input id="talkto" class="talkto" name="talkto" onclick="ghost.talkto()" type="button" value=" " />';
        talk_html += '      </div>';
        talk_html += '      <div class="inp_r" onclick="ghost.inp_r()">X</div>';
        talk_html += '  </div>';

        $(".wcc.smchuncai").append(talk_html);
        $(".wcc.smchuncai .addinput input.talk").keypress(function(event) {
            if (event.which == 13) {
                ghost.talkto();
            }
        });
        this.showInput();
    },

    showInput: function() {
        this.shell.closeMenu();
        //this.shell.closeNotice();
        this.shell.chuncaiSay("............?");
        //ghost.data.Ukagaka.setFace(1);
        $(".wcc .addinput").css("display", "block");
    },

    closeInput: function() {
        ghost.data.Ukagaka.setFace(3);
        $(".wcc .addinput").css("display", "none");
    },

    inp_r: function() {
        this.closeInput();
        this.shell.chuncaiSay('不聊天了吗？(→_→)');
        ghost.data.Ukagaka.setFace(3);
    },

    talkto: function() {
        this.getdata("talking");
    },

    clearInput: function() {
        $(".wcc.smchuncai .addinput input.talk").val('');
    },

    /**
     * 春菜喂食系统
     */
    foods: function() {
        this.shell.closeMenu();
        this.getdata("foods");
    },

    eatfood: function(obj, setCookie) {
        var gettimes = this.shell.tools.getCookie("eattimes");
        if (parseInt(gettimes) > parseInt(9)) {
            this.shell.chuncaiSay("主人是个大混蛋！！");
            ghost.data.Ukagaka.setFace(3);
            this.closechuncai_evil();
        } else if (parseInt(gettimes) > parseInt(7)) {
            this.shell.chuncaiSay(".....................肚子要炸了，死也不要再吃了～～！！！TAT");
            ghost.data.Ukagaka.setFace(3);
        } else if (parseInt(gettimes) == parseInt(5)) {
            this.shell.chuncaiSay("我已经吃饱了，不要再吃啦......");
            ghost.data.Ukagaka.setFace(3);
        } else if (parseInt(gettimes) == parseInt(3)) {
            this.shell.chuncaiSay("多谢款待，我吃饱啦～～～ ╰（￣▽￣）╭");
            ghost.data.Ukagaka.setFace(2);
        } else {
            var id = obj.replace("f", '');
            this.getdata('eatsay', id);
        }
        this.data.eattimes++;
        this.shell.tools.setCookie("eattimes", this.data.eattimes, 60 * 10 * 1000);
    },

    closechuncai_evil: function() {
        this.shell.stopTalkSelf();
        $(".wcc .showchuncaimenu").css("display", "none");
        setTimeout(function() {
            $(".wcc.smchuncai").fadeOut(1200);
            $(".wcc.callchuncai").css("display", "block");
        }, 2000);
    },


    /**
     * 见我家长
     */
    meetparents: function() {
        this.shell.closeChuncaiMenu();
        this.shell.closeNotice();
        //$("#getmenu").css("display", "none");
        this.shell.chuncaiSay("马上就跳转到我父母去了哦～～～");
        ghost.data.Ukagaka.setFace(2);
        setTimeout(function() {
            window.location.href = 'https://github.com/HinaKaze/web-ukagaka';
        }, 2000);
    },

    /**
     * 生存时间
     */
    lifetimechuncai: function() {
        this.shell.closeChuncaiMenu();
        this.shell.closeNotice();
        ghost.data.Ukagaka.setFace(2);
        this.getdata('showlifetime');
    },


    /**
     * 读取数据
     */
    getdata: function(el, id) {
        //$("#dialog_chat").fadeOut("normal");
        $(".wcc .tempsaying").css('display', "none");
        $(".wcc .dialog_chat_loading").fadeIn("normal");

        $.getJSON(ghost.data.Ukagaka.data._weichuncai_path, { time: new Date().getTime() })
            .done(function(dat) {
                $(".wcc .dialog_chat_loading").css('display', "none");
                //$("#dialog_chat").fadeIn("normal");
                $(".wcc .tempsaying").css('display', "");

                if (el == 'defaultccs') {
                    ghost.data.Ukagaka.chuncaiSay(dat.defaultccs);
                } else if (el == 'getnotice') {

                    //整合data里读取的自言自语
                    if (ghost.data.talkself_arr.length < 1) {
                        ghost.data.talkself_arr = ghost.data.Ukagaka.data.talkself_arr;
                    }
                    ghost.data.Ukagaka.data.talkself_arr = ghost.data.talkself_arr.concat(dat.talkself_user);

                    ghost.data.Ukagaka.chuncaiSay(dat.notice);
                    ghost.data.Ukagaka.setFace(1);

                } else if (el == 'showlifetime') {

                    var showlifetime = dat.showlifetime.replace(/\$time_str\$/, '许久');
                    if (!isNaN(parseInt(dat.lifetime[dat.defaultccs]))) {
                        var this_time = new Date();
                        var build_time = new Date(dat.lifetime[dat.defaultccs]);
                        var time_str = ghost.data.Ukagaka.tools.dateDiff(build_time, this_time);

                        showlifetime = dat.showlifetime.replace(/\$time_str\$/, time_str);
                    }

                    ghost.data.Ukagaka.chuncaiSay(showlifetime);

                } else if (el == 'talking') {

                    var talkcon = $(".wcc .talk").val();
                    var i = ghost.data.Ukagaka.tools.in_array(talkcon, dat.ques);
                    var types = typeof(i);
                    if (types != 'boolean') {
                        ghost.data.Ukagaka.chuncaiSay(dat.ans[i]);
                        ghost.data.Ukagaka.setFace(2);
                    } else {
                        ghost.data.Ukagaka.chuncaiSay('.......................嗯？');
                        ghost.data.Ukagaka.setFace(3);
                    }
                    ghost.clearInput();

                } else if (el == 'foods') {

                    var str = '';
                    var arr = dat.foods;
                    var preg = /function/;
                    for (var i in arr) {
                        if (arr[i] != '' && !preg.test(arr[i])) {
                            str += '<ul id="f' + i + '" class="eatfood" onclick="ghost.eatfood(this.id)">' + arr[i] + '</ul>';
                        }
                    }
                    ghost.data.Ukagaka.chuncaiSay(str);

                } else if (el = "eatsay") {

                    var str = dat.eatsay[id];
                    ghost.data.Ukagaka.chuncaiSay(str);
                    ghost.data.Ukagaka.setFace(2);

                } else if (el = "talkself") {

                    var arr = dat.talkself;
                    return arr;

                }
            })
            .fail(function(jqxhr, textStatus, error) {
                ghost.data.Ukagaka.chuncaiSay('好像出错了，是什么错误呢...请联系管理猿' + textStatus + error + jqxhr);
            });
    }
};
