# goterminus

A simple package that provides a bunch of terminal output helpers.


## Platform support

Goterminus has full support on 'Nix-like systems, but not quite complete
support on Windows. On windows the cursor movement helpers aren't yet implemented
as is `EraseInLine()`.


## Usage

### Colour

There's a series of helpers for colourising strings, there's two variants
for each colour:

``` go
term.Red("Hello World")
term.Redf("Hello %s", "World")
```

### Output and formatting

Action(actionName string, msg string)
* `Info`, `Infof` Outputs text in the default colour.
* `Action`, `Actionf` Outputs a action name, and a msg. The name will be
bright green and right aligned, the msg will be left aligned and the default colour.
* `Warn`, `Warnf` Will display the word "WARNING" in yellow, and the msg in
the default colour.
* `Error`, `Errorf` WIll displays the word "ERROR" in red, and the msg in
the default colour.
* `List` Displays an array of items as a multi-column list, using the desired
number of columns. The column width must be supplied.
* `Table` Displays a set of data in a table, with an optional header.
