# Introduction
This is a simple toy project I use to send notifications to a locally running instance of ntfy

# Why not use bash
Yep, good question. I wanted to do something in go

# Usage
The usage is pretty straight forward there is a few arguments

> duptfy -s <-YOUR ntfy Server-> -c <-BASH COMMAND TO RUN, CAN INCLUDE PIPES-> -m <-THE TITLE OF THE ntfy MESSAGE->

# Branch Names

Branch names will adjust how tags are created:
* bugfix/description_of_the_change => patch
* feature/description_of_the_change => minor
* release/description_of_the_change => mayor