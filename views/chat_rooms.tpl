<script type="text/javascript">
	$(document).ready(function () {
		$(".room-content").click(function(){
    		$.get("/chat_enter_room",function(data,status){
        		$("#main_content").html(data);
        	});
    	}); 
    	$("#create_room_btn").click(function(){
    		$("#room_config").toggle(500)
    	}); 
	});
</script>

<div class="lobby">
<div class="glyphicon glyphicon-plus" id="create_room_btn"></div>
<div id="room_config" class="room-config">
	<form role="form">
		<div class="form-group" action="create_room">
			<label for="room_topic">Topic</label>
			<input type="text" class="form-control" id="room_title" name="topic" placeholder="Please input room topic!"/>
		</div>
		<div class="form-group">
			<label for="room_maxmember">MaxMember</label>
			<input type="text" class="form-control" id="room_maxmember" name="maxmember" placeholder="Please input room max member size!">
		</div>
		<button type="submit" class="btn btn-primary">Create!</button>
	</form>
</div>
{{range .List}}
<div class="room-detail">
	<div class="room-content" id="enter{{.Id}}">
		<div class="room-date-box">
		<span class="room-date-box-day">{{.CreateDay}}</span>
		<br/>
		<span class="room-date-box-month">{{.CreateMonth}}</span>
		<span class="room-date-box-year">{{.CreateYear}}</span>
		</div>
		<h1>
			<strong>{{.Id}}.{{.Name}}</strong>
		</h1>
		<small>创建者：{{.Creator.Name}}</small>
	</div>
</div>
{{end}}
</div>