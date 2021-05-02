# TC-moderation
Contains all the code you'll need, Binod.

So
> Step 1
Go to [Moderation](https://yagpdb.xyz/manage/794236519543734273/moderation) **>** Then, `Warnings`
And make sure the **left** side looks like the image below  
![image1](https://user-images.githubusercontent.com/83547071/116819849-e9f51180-ab69-11eb-8627-466bfcc25d75.png)
Then, add the code from [WarnDM](https://github.com/TC-club/TC-moderation/blob/main/moderation/warnDM.cc.go) to where it says `WarnDM`

> Step 2
While you're still in **Moderation** go to `Mute`
Make sure the **left** side looks like the image below (Make sure to **ENABLE** `Mute/Unmute`  
![image2](https://user-images.githubusercontent.com/83547071/116819961-81f2fb00-ab6a-11eb-9d11-5ae05514e37c.png)
Then, add the code from [MuteDM](https://github.com/TC-club/TC-moderation/blob/main/moderation/muteDM.cc.go) & [UnmuteDM](https://github.com/TC-club/TC-moderation/blob/main/moderation/unmuteDM.cc.go) to where it says `Mute DM` & `Unmute DM`

> Step 3
Again. Whilst still being in **Moderation** go to `Kick`
Now I'm assuming this is right (the roles), make sure the **left** side looks like the image below  
![image3](https://user-images.githubusercontent.com/83547071/116820088-11001300-ab6b-11eb-9e0c-098b21964154.png)
Then, add the code from [KickDM](https://github.com/TC-club/TC-moderation/blob/main/moderation/kickDM.cc.go) to where it says `Kick DM`

> Step 4
Yep. You guessed it. Head over to the `Ban` tab
Now, same old same old. Make sure your settings mirror the image below. 
![image4](https://user-images.githubusercontent.com/83547071/116820162-72c07d00-ab6b-11eb-9cfb-4fc7ca29e579.png)
Then add the code from [BanDM](https://github.com/TC-club/TC-moderation/blob/main/moderation/banDM.cc.go) to where it says `Ban DM`

# Final steps

> Step 5
Make your way over to [Custom commands](https://yagpdb.xyz/manage/794236519543734273/customcommands)
You want to make 3 commands.
In command number one, copy and paste what is in [Viewcase](..blob/main/custom%20commands/Viewcase.cc.go)
Do the same with [cases](..blob/main/custom%20commands/cases.cc.go) & [delcase](..blob/main/custom%20commands/delete-case.cc.go)

# Note. When adding the custom commands, if you check the TOP comment in each file. That will tell you the trigger type+trigger, don't include the backticks when pasting the trigger
