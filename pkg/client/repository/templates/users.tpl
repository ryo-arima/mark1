+----+--------------------------------------+-----------------------+-----------+-----------+---------------------+---------------------+---------------------+
| ID |                UUID                 |         Email         |   Name    |  Status   |     Created At      |     Updated At      |     Deleted At      |
+----+--------------------------------------+-----------------------+-----------+-----------+---------------------+---------------------+---------------------+
{{range .Users}}| {{.ID}}  | {{.UUID}} | {{.Email}} | {{.Name}} | {{.Status}} | {{.CreatedAt}} | {{.UpdatedAt}} | {{if .DeletedAt}}{{.DeletedAt}}{{else}}                     {{end}} |
+----+--------------------------------------+-----------------------+-----------+-----------+---------------------+---------------------+---------------------+
{{end}}