# Discord-Chuck-Norris

This is a Discord Chuck Norris bot that returns you a Chuck Norris joke when asked for. Integrate it with your Discord server in under five minutes.

## Requirements

1. [hasura CLI tool](https://docs.hasura.io/0.15/manual/install-hasura-cli.html)
2. A Discord server

## Setup guide

### Get the project

```
$ hasura quickstart rishi/discord-golang-bot
```

### Create a webhook to send messages on your channel

1. Go to your Discord server.

2. Click on the down arrow next to your server name and click on `Server Settings`.

3. Select the `Webhooks` section from the left pannel. Create a webhook.

4. Copy the Webhook URL and add it to project secrets since it contains a secret token that you do not want to explicitly mention in the code.

```
$ hasura secret update discord.webhook <webhook_url>
```

### Create a bot on discord

1. [Create an application](https://discordapp.com/developers/applications/me).

2. Now, in your application home page, scroll down to `Bots` and create a bot.

3. Make the bot public.

4. Make sure that the `Require OAuth2 Code Grant` checkbox is not checked.

5. Add the bot to your server by clicking on `Generate OAuth URL` and navigating to the generated URL from the browser.

5. Copy the bot token and add it to the project secrets.

```
$ hasura secret update discord.bot.token <bot_token>
```

### Push the project to the cluster

```
$ git add .
$ git commit -m "First commit"
$ git push hasura master
```

## Usage

Type `!joke` in the server chat and see the magic :)

## Modification

This project can also be used as a starter for building more complex Discord bots. The source code for the bot lies in the `microservices/bot/app/src/` directory.

You might also want to look at the Dockerfile and the k8s.yaml at `microservices/bot/` if you are to add some extra packages or environment variables.

## References

1. Special thanks to "bwmarrin" for the [discord bindings for golang](https://github.com/bwmarrin/discordgo).
2. Built on top of Hasura boilerplate [hello-golang-raw](https://hasura.io/hub/project/hasura/hello-golang-raw).
