# TODO

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
- [ ] DialogueOptions
    - [ ] Prompt for continuation of text
    - [ ] Display Choices
    - [ ] Allow selection of Choice with Mouse and Keyboard
    - [ ] New Glyphs for Choices

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
