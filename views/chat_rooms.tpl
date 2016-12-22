{{.Title}}
{{range .List}}
<div class="card">
	<div class="card-content">
		<div class="card-date-box">
			{{.CreateTime}}
		</div>
		<h1>
			<strong>{{.Id}}.{{.Name}}</strong>
		</h1>
		<small>创建者：{{.Creator.Name}}</small>
	</div>
</div>
{{end}}