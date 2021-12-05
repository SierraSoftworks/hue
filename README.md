# Hue CLI
**Control your Philips Hue lights from your command line**

Do you just **HATE** needing to reach for your phone to turn on the lights
when the sun goes down? Are you the kind of person who believes that GUIs
were a mistake? Just looking for a quick way to turn things off at the end
of the day?

This app is for you, it's a lightweight command line client which lets you
control your Hue lights with minimal fuss and no extra effort.

```powershell
# Install the latest version of the Hue CLI
go install github.com/sierrasoftworks/hue@latest

# Connect to your Hue bridge
hue setup

# Set up your lights just the way you want them
hue all=off tv=red,25% bedroom=orange,30%
```