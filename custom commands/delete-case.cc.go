{{/*
        Made by Maverick Wolf (549820835230253060)
        Modified by Ranger (765316548516380732)
        Credit to LemmeCry (664118444739919882)

    Trigger Type: `RegEx`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(delcase|deletecase|clearcase|dc|clc)(\s+|\z)`
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
    {{$check :=0}}
    {{$roles :=cslice 807972168045690941 794243296657670174 808767399493894196 807610850541568031 794237238686515210}} {{/* All staff roles */}}
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
        {{$args := parseArgs 1 "correct usuage is `-delcase <case number>`" (carg "int" "case number")}}
        {{$caseno := ($args.Get 0)}}
        {{$id := (toInt (dbGet $caseno "userid").Value)}}
        {{dbDel $caseno "viewcase"}}
        {{dbDel $caseno $id}}
        {{dbDel $caseno "userid"}}
        {{$embed := sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Deleted the case")
            "color" 0x36393f
            )}}
    {{end}}
{{end}}
