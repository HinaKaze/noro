<script type="text/javascript">
$(document).ready(function () {
	var ws = new WebSocket('ws://'+window.location.host+'/ws?user_name='+$('#chat_username').text());
	ws.onmessage = function(event){
		var data = JSON.parse(event.data);
		console.log(data);
		switch(data.Type){
			case 0: //join
				addJoin(data.User.Name)
			case 1: //leave
				addLeave(data.User.Name)
			case 2: //message
				addMessage(data.User.Name,data.Text,data.Time)
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
	$("#messages").append("<div>"+name+"join room</div>")
}

function addLeave(name){
	$("#messages").append("<div>"+name+"leave room</div>")
}
})



</script>
<input type="text" id="chat_text" placeholder="输入消息"></input>
<input type="text" id="chat_username" placeholder="输入用户名"></input>
<button id="chat_send" class="btn">发送</button>
<div id="messages">

</div>