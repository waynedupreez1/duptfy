# Introduction
This is a simple toy project I use to send notifications to a locally running instance of ntfy

# Why not use bash
Yep, good question. I wanted to do something in go

# Usage
The usage is pretty straight forward there is a few arguments

> duptfy -s <-YOUR ntfy Server-> -c <-BASH COMMAND TO RUN, CAN INCLUDE PIPES-> -m <-THE TITLE OF THE ntfy MESSAGE->

# Branch Names

Branch names to adjust semver
* patch => bugfix/description_of_the_change  results in => v0.1.0 results in v0.1.1
* minor => feature/description_of_the_change results in => v0.1.0 results in v0.2.0
* major => release/description_of_the_change results in => v0.1.0 results in v1.0.0

Branch names to adjust semver
* <-ANYTHING ELSE->/description_of_the_change results in => v0.1.0 results in v0.1.0+1
