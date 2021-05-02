{{/*
        Made by Maverick Wolf (549820835230253060)
        Modified by Ranger (765316548516380732)

    Trigger Type: `Ban DM`
©️ Dynamic 2021
MIT License
*/}}

{{/* Configuration values start */}}
{{$LogChannel := 838432051094880306}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{if gt ( toInt ( currentTime.UTC.Format "15" ) ) 12}}
{{end}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
	{{$ext := "webp"}}
	{{if eq (slice .Guild.Icon 0 2) "a_"}} {{$ext = "gif"}} {{end}}
	{{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{$dm := 1}}{{/* change to 0 if you don't wanna DM the offender about the moderation action*/}}

{{if $dm}}
    {{$BanDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Ban`\n**Duration : **" .HumanDuration "\n**Reason: **" (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")) "\n\nIf the ban is permanent and you believe you have been wrongly banned you can fill out an [appeal](https://forms.gle/TX2CoL9y1opQnD589) in 1 month.")
            "thumbnail" (sdict "url" $icon)
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
{{sendDM $BanDM}}
{{end}}

{{$case_number := (toInt (dbIncr 77 "cv" 1))}}
{{$action := .ModAction.Prefix}}
{{$a := 0}}

{{if eq $action "Muted" "Unmuted"}}
    {{$a = (sub (len .ModAction.Prefix) 1)}}
    {{else}}
    {{$a = (sub (len .ModAction.Prefix) 2)}}
{{end}}

{{$title := (slice .ModAction.Prefix 0 $a)}}
{{$id := .User.ID}}
{{$channel := $LogChannel}}

{{$x := sendMessageRetID $LogChannel (cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print "<:TextChannel:800978104105304065> **Case number** " $case_number "\n<:Management:788937280508657694> **Who:** " .User.Mention " `(ID " .User.ID ")`\n<:Metadata:788937280508657664> **Action:** `Ban`\n<:Assetlibrary:788937280554926091> **Channel:** <#" .Channel.ID ">\n<:Manifest:788937280579698728> **Reason:** " (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")) "\n:clock12: **Time:** " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "footer" (sdict "text" (print "Duration: " .HumanDuration ))
            "color" 14043208
            )}}

{{dbSet $case_number "viewcase" (sdict "name" .Author.Username "warnname" .User.Username "avatar" (.Author.AvatarURL "512") "reason" .Reason "userid" $id "action" (.ModAction.Prefix) "channel" $channel "msgid" $x "userdiscrim" .User.Discriminator)}}{{/*for viewcase*/}}
{{dbSet $case_number $id (print "Case # **" $case_number "**\t\t**| " $title " Reason:** `" .Reason "`")}}{{/*for per user case viewing*/}}
{{dbSet $case_number "userid" (str $id)}}{{/* for delete case*/}}

{{$Response := sendMessageRetID nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Case type: Ban"))
            "description" (print .Author.Mention " Has successfully banned " .User.Mention " :thumbsup:")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            )}}
{{deleteMessage nil $Response 5}}
