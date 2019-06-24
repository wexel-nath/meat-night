package template

const AlertHostSubject = "Meat Night"

// AlertHostHtml uses model.Mateo
const AlertHostHtml = `
<html>
	<h1>{{ .FirstName }},</h1>
	<p>You're up for meat night this week.</p>
	<p>Let everyone know if you can make it or not.</p>
	<img src="https://i.ibb.co/yPQBzKc/15ead358-dddc-4e5d-a5c7-5a26ad86e469-200x200.png" alt="Mateo Corp Logo" />
</html>
`

const AlertHostText = `
{{ .FirstName }},

You're up for meat night this week.

Let everyone know if you can make it or not.

Mateo Corporation
`
