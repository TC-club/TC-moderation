{{/*
        Made by Maverick Wolf (549820835230253060)
        Modified by Ranger (765316548516380732)
        Credit to LemmeCry (664118444739919882)

    Trigger Type: `RegEx`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(viewcase|seecase|vc|sc|case)(\s+|\z)`
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
    {{$args := parseArgs 1 "`-viewcase <case number>`" (carg "int" "case number")}}
    {{$a := ""}}
    {{$name := ""}}
    {{$iconurl := ""}}
    {{$warnname := ""}}
    {{$reason := ""}}
    {{$userid := ""}}
    {{$userdiscrim := ""}}
    {{$action := ""}}
    {{$link := ""}}
    {{$msgid := ""}}
    {{$channel := ""}}
    {{$server := .Guild.ID}}
    {{$b := ($args.Get 0)}}
    {{with (dbGet $b "viewcase")}}	
        {{$a = sdict .Value}}
        {{$name = $a.name}}
        {{$iconurl = $a.avatar}}
        {{$warnname = $a.warnname}}
        {{$reason = $a.reason}}
        {{$action = $a.action}}
        {{$userid = $a.userid}}
        {{$userdiscrim = $a.userdiscrim}}
        {{$msgid = $a.msgid}}
        {{$channel = $a.channel}}
        {{$embed := cembed
            "author" (sdict "name" $name "url" "" "icon_url" $iconurl)
            "description" (print "[**Jump to Case**](https://discordapp.com/channels/" $server "/" $channel "/" $msgid "/)\n**Who** : " $warnname "#" $userdiscrim "\t`(" $userid ")`\n **Action** : `" $action "`\n **Reason** : " $reason)
            "color" 0x36393f
            "footer" (sdict "text" (print "Case #" $b))}}
        {{sendMessage nil $embed}} 
        {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Could not find the case specified, Please make sure the case number is correct or the case has not been deleted")
            )}}
    {{end}}
{{end}}
