[![Coverage Status](https://coveralls.io/repos/github/ala24013/tweetviz/badge.svg?branch=main)](https://coveralls.io/github/ala24013/tweetviz?branch=main)

# :bird::eyeglasses: Tweetviz

Project for visualizing locations of topical tweets.

## :page_facing_up: Usage

### :computer: Creating a Twitter Developer Account

Tweetviz requires a Twitter application (developer) account to be able to connect and stream tweets. The developer's account credentials are not provided here. To acquire a Twitter application, please check their website here: https://developer.twitter.com/apps

### :traffic_light: Acquiring a Bearer Token

Once you have created an application, you will need to get the Bearer Token associated with your account. To find this, go to the twitter developer page (listed above) select your project/application, and click on the Keys and Tokens tab. Under this tab there should be an "Authentication Tokens" section that will allow you to create and manage your Bearer Token (as well as other keys and credentials).

### :rocket: Launching the Application

To launch the application, a makefile has been provided at the top level. You must first set the environment variable TWITTER_BEARER_TOKEN to your bearer token before running the makefile. Once you have done so, simply type
```bash
make run
```
and a webserver will instantiate, running a website on your local computer that will visualize the incoming tweets.

Then, just open your favorite web browser and go to http://localhost:3000/ to open the server. Happy visualizing! :sunglasses:
