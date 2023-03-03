# ANIMALI - WIP
ANIMALI is a mobileapp build in golang using fyne : https://developer.fyne.io/index.html

## Purpose
There a four main goals for this project :
- Creating a app for my daughter
- Testing fyne as a mobilapp developing solution
- Going thru publishing the mobile application to the google play store
- Add something on github because all of my project end up on private repository of my company



## Build 
Builds desktop app preview
```
make 
```

Runs go build whitout building assets 
```
make build-go
```

Builds Android APK
```
make build-android
```

Generate's assets from ./assets dir to internal/animali/assets.go
```
make build-assets
```

Builds only APK whitout regenerating assets
```
make build-apk
```


## TO DO : 
 - add grid with animals sounds
 - add sound and images for music and animal screens
 - refactor
 - test
 - add language support
 - about

## Nice To Have :
 - Add badtime story's
 - Add white noises


## Credits 
### Sound Effect 
Universfield, Pixabay


### MIDI 
https://www.romwell.com/kids/nursery_rhymes/kids_midi.shtml

### Images 
