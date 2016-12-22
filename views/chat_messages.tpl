{{range .List}}
<div class="message">
	<div class="message-user">
		<img class="message-user-avatar img-thumbnail" src="static/img/user/avatar{{.User.Id}}.jpg"></img>
		<div class="message-user-name">{{.User.Name}}</div>
	</div>

	<div class="message-content">
		<div class="message-content-text text">{{.Text}}</div>
		<div class="message-content-time">{{.Time}}</div>
		<!-- <span ownFlag=""></span> -->
	</div>
</div>
<br/>
{{end}}