<script type="text/javascript">
$(document).ready(function () {
	
	//fetch room detail
	msgIndex = {{.HistoryMsgLength}};
	chatRoomDetail = {
		"Id" : -1,
		"Name" : "noro作战本部",
		"Creator":{},
		"MaxMember":0,
		"Mates":[]
	}
	if (msgIndex > 0){
		$("html,body").animate({scrollTop:$("#message"+msgIndex).offset().top},500)
	}
	$.get("/chat_room",function(data,status){
		chatRoomDetail.Id = data.Id
		chatRoomDetail.Name = data.Name
		chatRoomDetail.Creator = data.Creator
		chatRoomDetail.MaxMember = data.MaxMember
		chatRoomDetail.Mates = data.Mates
    	//chatRoomDetail = JSON.parse(data)
  		//console.log('111',chatRoomDetail)
    	showRoomMates()
    	var ws = new WebSocket('ws://'+window.location.host+'/ws?user_name='+$('#chat_username').text());
    	ws.onmessage = function(event){
			var data = JSON.parse(event.data)
			switch(data.Type){
				case 0: //join
					addJoin(data.User)
					break;
				case 1: //leave
					addLeave(data.User)
					break;
				case 2: //message
					addMessage(data.User,data.Text,data.Time)
					break;
			}
		}
		$("#chat_send").click(function(){
			var message = $("#chat_text").val()
			if (message == ""){
				return
			}
			ws.send(message)
			$("#chat_text").val("")
		});
    });
	

	
	function addMessage(user,text,time){
		msgIndex += 1
		var http = '\
		<div class="message" id="message'+msgIndex+'">\
			<div class="message-user">\
				<img class="message-user-avatar" src="static/img/user/avatar1.jpg"></img>\
				<div class="message-user-name">'+user.Name+'</div>\
			</div>\
			<div class="message-content">\
				<div class="message-content-text text">'+text+'</div>\
				<div class="message-content-time">'+time+'</div>\
			</div>\
		</div>\
		'
	$("#messages").append(http)
	$("html,body").animate({scrollTop:$("#message"+msgIndex).offset().top},500)
	}

	function addJoin(user){
		$("#messages").append("<div class='chat-tip'>"+user.Name+" 加入 房间</div>")
		chatRoomDetail.Mates.push(user);
		showRoomMates()
	}

	function addLeave(user){
		$("#messages").append("<div class='chat-tip'>"+user.Name+" 离开 房间</div>")
		for (i in chatRoomDetail.Mates){
			if (chatRoomDetail.Mates[i].Name == user.Name){
				chatRoomDetail.Mates.splice(i,1)
				break;
			}
		}
		showRoomMates()
	}
	function showRoomMates(){
		for (u of chatRoomDetail.Mates){
			if (u != undefined){
				console.log(u)
			var http = '\
				<div><span class="glyphicon glyphicon-user"></span>\
				<span class="chat-room-mate">'+u.Name+'</span></div>\
			';
			$("#left_content").append(http);
			}
		}
	}
})


</script>
<div class="chat-room">
<div id="messages">
{{range .HistoryMsgs}}
	{{if eq .Type 2}}
	<div class="message" id="message{{.Id}}">
			<div class="message-user">
				<img class="message-user-avatar" src="static/img/user/avatar1.jpg"></img>
				<div class="message-user-name">{{.User.Name}}</div>
			</div>
			<div class="message-content">
				<div class="message-content-text text">{{.Text}}</div>
				<div class="message-content-time">{{.Time}}</div>
			</div>
	</div>
	{{else if eq .Type 0}}
	<div class='chat-tip' id="message{{.Id}}">{{.User.Name}} 加入 房间</div>
	{{else if eq .Type 1}}
	<div class='chat-tip' id="message{{.Id}}">{{.User.Name}} 离开 房间</div>
	{{end}}
{{end}}
</div>
<div class="input-box">
<input type="text" id="chat_text" placeholder="输入消息"></input>
<!-- <input type="text" id="chat_username" placeholder="输入用户名"></input> -->
<button type="button" id="chat_send" class="btn btn-warning">发送</button>
</div>

</div>