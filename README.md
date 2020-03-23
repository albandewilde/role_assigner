# Role Assigner

Discord bot for assign roles.

## Presentation

I made this bot for assignating role on discord server.

Often, when you have a discord server, you want people to self asign roles but
with out give them the right to manage roles.

This is exactly what this bot do.

You give him a message where the bot watchs reactions, and when someone add or remove
an emote to the message, the bot add of remove the role to the user.  
The bot also check each day if all users have (or haven't) roles in agreements
with the message's emotes reactions.

The link between roles and emotes is made with the `roles.json` file.

## Usage

This bot isn't made to run on multiple server. (It can, but the configuration
will be the same.)

You have to configure the `roles.json` file.  
To do that, you can rename the `roles.json.tpl` file to `roles.json` and
complete it.

The `roles.json` file must contain an object.  
Keys are a string which is the emote name.  
Values are a string which is the role name.

Example of `roles.json` file:
```json
{
    "py": "python",
    "rs": "rustacean",
    "go": "gopher",
    "hs": "haskell",
    "kt": "kotlin"
}
```

You have to put you discord token in a `.env` file.  
The `.env.tpl` file is here to help you complete it.
