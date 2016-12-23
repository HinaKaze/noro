<script type="text/javascript">
$(document).ready(function () {
	var ws = new WebSocket('ws://'+window.location.host+'/ws?user_name='+$('#chat_username').text());
	//fetch room detail
	chatRoomDetail = {}
	$.get("/chat_room",function(data,status){
		chatRoomDetail = data
    	//chatRoomDetail = JSON.parse(data)
  		//console.log('111',chatRoomDetail)
    	showRoomMates()
    });
	ws.onmessage = function(event){
		var data = JSON.parse(event.data);
		switch(data.Type){
			case 0: //join
				addJoin(data.User.Name)
				break;
			case 1: //leave
				addLeave(data.User.Name)
				break;
			case 2: //message
				addMessage(data.User.Name,data.Text,data.Time)
				break;
		}
	}

	$("#chat_send").click(function(){
		var message = $("#chat_text").val()
		ws.send(message)
		$("#chat_text").val("")
	});
	function addMessage(name,text,time){
	var http = '\
		<div class="message">\
			<div class="message-user">\
				<img class="message-user-avatar img-thumbnail" src="static/img/user/avatar1.jpg"></img>\
				<div class="message-user-name">'+name+'</div>\
			</div>\
			<div class="message-content">\
				<div class="message-content-text text">'+text+'</div>\
				<div class="message-content-time">'+time+'</div>\
			</div>\
		</div>\
	'
	$("#messages").append(http)
	}

	function addJoin(name){
		$("#messages").append("<div class='chat_tip'>"+name+" 加入 房间</div>")
		console.log(chatRoomDetail)
		chatRoomDetail.Mates.push("name");
		showRoomMates()
	}

	function addLeave(name){
		$("#messages").append("<div class='chat_tip'>"+name+" 离开 房间</div>")
		for (i in chatRoomDetail.Mates){
			if (chatRoomDetail.Mates[i].Name == name){
				chatRoomDetail.Mates.splice(i,1)
				break;
			}
		}
		showRoomMates()
	}
	function showRoomMates(){
		alert("11121")
		for (u of chatRoomDetail.Mates){
			var http = '\
				<div class="chat-room-mate">'+u.Name+'</div>\
			';
			$("#right_content").append(http);
		}
	}
})


</script>
<input type="text" id="chat_text" placeholder="输入消息"></input>
<input type="text" id="chat_username" placeholder="输入用户名"></input>
<button id="chat_send" class="btn">发送</button>
<div id="messages">

</div>