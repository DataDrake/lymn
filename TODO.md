# TODO

## UI

- [x] Titlebar
    - [x] Title Text
    - [x] Location Text
- [ ] TextArea
    - [ ] Display a border box
    - [ ] Display text with wrapping in Box
        - [ ] Stretch: Hyphenation while wrapping
    - [ ] Stretch: Display Text Character by Character
    - [ ] Stretch: Play randomized type-writer sounds as text prints
- [ ] PlayerStats
    - [ ] Stat bars
        - [ ] Text Label
        - [ ] ProgressBar (with colors)
            - [ ] Add Glyphs for empty portion of the bar
        - [ ] 2D Grid?
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
