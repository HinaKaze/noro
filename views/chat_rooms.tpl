{{range .List}}
<div class="card">
	<div class="card-content">
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