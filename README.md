# LambdaLabs Sniper

This is a LambdaLabs Sniper for getting your desired instance. I am NOT responsible for how you use this software. Read the LICENSE

This might also lower your score on stripe cause it loads a stripe cookie

Currently works as a notifier

## How to use?

1. Change `copy.env` to `.env`
1. Fill out the environment variables with your stuff
    - To get the SESSION_ID do the following:
        - Go to https://cloud.lambdalabs.com/ and log in
        - Press Ctrl+Shift+I (or right-click > inspect) to open up Dev Tools
        - Go to Application > Storage > Cookies > https://cloud.lambdalabs.com
        - Copy the `sessionid` cookie
    - I assume that you know how to create a discord webhook by now, so won't mention it
1. go run main.go and wait for notifs