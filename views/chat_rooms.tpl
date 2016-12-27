<script type="text/javascript">
	$(".card-content").click(function(){
    	$.get("/chat_enter_room",function(data,status){
        	$("#main_content").html(data);
        });
    }); 
</script>

{{range .List}}
<div class="card">
	<div class="card-content" id="enter{{.Id}}">
		<div class="card-date-box">
		<span class="card-date-box-day">{{.CreateDay}}</span>
		<br/>
		<span class="card-date-box-month">{{.CreateMonth}}</span>
		<span class="card-date-box-year">{{.CreateYear}}</span>
		</div>
		<h1>
			<strong>{{.Id}}.{{.Name}}</strong>
		</h1>
		<small>创建者：{{.Creator.Name}}</small>
	</div>
</div>
{{end}}