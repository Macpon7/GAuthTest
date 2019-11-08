# GAuthTest
This little web app is deployed to heroku, and allows a user to log in through google.

`https://goauthtest.herokuapp.com/login`

Google returns this as a response to getUserDataFromGoogle:

    email:string
        The user's email address.
    family_name:string
        The user's last name.
    gender:string
        The user's gender.
    given_name:string
        The user's first name.
    hd:string
        The hosted domain e.g. example.com if the user is Google apps user.
    id:string
        The obfuscated ID of the user.
    link:string
        URL of the profile page.
    locale:string
    The user's preferred locale.
        name:string
    The user's full name.
        picture:string
    URL of the user's picture image.
        verified_email:boolean
    Boolean flag which is true if the email address is verified. Always verified because we only return the user's primary email address.

The only thing we are currently saving is email, but more can be added to the userInfo struct if required.