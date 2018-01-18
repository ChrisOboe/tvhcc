# tvhcc
A simple command line client for tvheadend.

This project uses the tvheadend web api. This is the internal api for their
webinterface. A proper tvheadend client would use the htsp protocol instead.
This client is done quick and dirty and can break anytime.

## Features:
 * Display the epg
 * Display a list of channels
 * Start a stream in mpv

## Usage:
 * `tvhcc <server> epg` for displaying epg
 * `tvhcc <server> channels` for listing channels
 * `tvhcc <server> play <channelName>` for starting mpv

### Example:
 `tvhcc http://192.168.1.1 epg`

## Todo:
 * Add bash completion
 * Add configuration file (for videoplayer and default server)
 * Nicer output with colors and fancy stuff
 * Try to find out if the broken sorting is the fault of my tvheadend configuration or if it's a tvheadend bug
