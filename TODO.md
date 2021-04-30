# TODO

## Jam Prep

- [ ] main.go cleanup
    - [x] Move Game to game.go
    - [x] Move Game UI bits to ui/window.go
    - [ ] Script gets given to Game on creation
    - [ ] Font gets given to Game on creation
    - [ ] Move script.yml and font.json to example/
    - [ ] Load Script and Font in main.go
    - [ ] Allow main.go to register Stat bars for the player
- [x] UI cleanup
    - [x] Create ui.Window
    - [x] Give ui.Window dimensions
    - [x] ui.Window.Size() -> grid.Size()
    - [x] Rework UI elements to have X,Y coordinates and fixed sizes
- [ ] Items
    - [ ] model: Allow Player to hold items
    - [ ] engine: Add Item type for Events
    - [ ] ui: New Item List on right side

## UI

- [x] Titlebar
    - [x] Title Text
    - [x] Location Text
- [ ] TextArea
    - [x] Display a border box
    - [x] Display text with wrapping in Box
        - [ ] Stretch: Hyphenation while wrapping
    - [ ] Stretch: Display Text Character by Character
    - [ ] Stretch: Play randomized type-writer sounds as text prints
- [x] PlayerStats
    - [x] Stat bars
        - [x] Text Label
        - [x] ProgressBar (with colors)
            - [x] Add Glyphs for empty portion of the bar
- [x] DialogueOptions
    - [x] Prompt for continuation of text
    - [x] Display Choices
    - [ ] Allow selection of Choice with Mouse and Keyboard

## Scripting
 - [x] Defined data types
 - [x] Parse script

## Engine
 - [x] Keep track of script execution
    - [x] Internal State
    - [x] State Machine
        - [x] Scenes
        - [x] Events
        - [x] Choices
        - [x] Text
        - [x] Stats

## Model
 - [x] Characters
    - [x] Stats
 - [x] Scene
    - [x] Title
    - [x] Location
 - [x] Choices
    - [x] Runes
    - [x] Selected

## Input

 - [x] Borrow keyboard input from Pixie
    - [x] Optimize to only allow one listener per Key Press (no Release)
