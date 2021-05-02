{{/*
        Made by Maverick Wolf (549820835230253060)
        Modified by Ranger (765316548516380732)
        Credit to LemmeCry (664118444739919882)

    Trigger Type: `RegEx`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(cases|allcase)(\s+|\z)`
©️ Dynamic 2021
MIT License
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{$check := 0}}
{{$roles := cslice 807972168045690941 794243296657670174 808767399493894196 807610850541568031 794237238686515210}} {{/* All staff roles */}}
{{range $roles}}
    {{if eq $check 0}}
        {{if hasRoleID .}}
        {{$check =1}}
        {{end}}
    {{end}}
{{end}}

{{if eq $check 0}}
    {{sendMessage nil (cembed
            "author" (sdict "name" (print .User.Username) "icon_url" (.User.AvatarURL "512"))
            "description" (print "<:Cross:817828050938363905> I'm sorry. You don't have permission to use this command.")
            "color" 0x36393f
            )}}
    {{else}}
    {{$args := parseArgs 1 "correct usuage is `-cases <User/Mention> [Page]`" (carg "member" "target user") (carg "int" "page number")}}
    {{$user := ($args.Get 0).User}}
    {{$id := $user.ID}}
    {{$display := ""}}
    {{$page := 1}} {{/* Default page to start at */}}
    {{if ($args.IsSet 1)}}
        {{if ge ($args.Get 1) 1}}
            {{$page = ($args.Get 1)}}
            {{else}}
            {{$page = 1}}
        {{end}}
    {{end}}
    {{$skip := mult (sub $page 1) 10}}
    {{$cases := dbTopEntries $id 10 $skip}}
    {{range $cases}}
        {{- $value := str .Value }} 
        {{- $display = joinStr "" $display "\n" $value -}} 
        {{else}}
        {{$display = "`No cases exist on this page`"}}
    {{end}}
    {{$id := sendMessageRetID nil (cembed
            "author" (sdict "name" (print $user.Username " (" $user.ID ")") "icon_url" ($user.AvatarURL "512"))
            "title" "Cases" 
            "description" $display
            "footer" (sdict "text" (print "Page " $page))
            "color" 0x36393f
            )}}
{{end}}
