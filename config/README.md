## Configuration Files

The default location for the configuration file is $BLOCKORGROOT/blockorg.json. If you want to change this location, you will need to change the code in this package.

By default, the code looks for an environmental variable named $BLOCKORGROOT. If it is not there it is assigned the default of '/home/user/.config/blockorg'.

If $BLOCKORGROOT/blockorg.json does not exist, the system will panic and exit. I have decided against using a default configuration for security reasons.

The provided example configuration should be used as a goby when creating your own configuration. Feel free to copy it into place and adjust the values.
