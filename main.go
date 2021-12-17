{{ $args := parseArgs 1 "Syntax: -atar @command"
(carg "string" "command to execute") }}
{{ $command := $args.Get 0 }}
{{ $host := "https://raw.githubusercontent.com/DNIB/atar-talking/main/" }}
{{/* JPG */}}
{{ if (eq $command "africa") }}
	{{ $host }}{{"img/africa.jpg" }}
{{ else if (eq $command "hentai") }}
	{{ $host }}{{"img/hentai.jpg" }}
{{ else if (eq $command "noBoki") }}
	{{ $host }}{{"img/noBoki.jpg" }}
{{ else if (eq $command "noGG") }}
	{{ $host }}{{"img/noGG.jpg" }}
{{ else if (eq $command "shit") }}
	{{ $host }}{{"img/shit.jpg" }}
{{ else if (eq $command "tits") }}
	{{ $host }}{{"img/tits.jpg" }}
{{ else if (eq $command "yuna") }}
	{{ $host }}{{"img/yuna.jpg" }}
{{/* GIF */}}
{{ else if (eq $command "handshake") }}
	{{ $host }}{{"gif/handshake.gif" }}
{{ else if (eq $command "lie") }}
	{{ $host }}{{"gif/lie.gif" }}
{{ else if (eq $command "totoko") }}
	{{ $host }}{{"gif/totoko.gif" }}
{{ else if (eq $command "yuksu") }}
	{{ $host }}{{"gif/yuksu.gif" }}
{{ else if (eq $command "help" "-h" "--help") }}
	{{ "``` @see https://github.com/DNIB/atar-talking/blob/main/readme.md ```" }}
{{ else }}
	{{ "Error: Command Not Existed" }}
{{ end }}
