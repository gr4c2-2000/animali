# ANIMALI - WIP
ANIMALI is a mobileapp build in golang using fyne : https://developer.fyne.io/index.html

## Purpose
There a four main goals for this project :
- Creating a app for my daughter
- Testing fyne as a mobilapp developing solution
- Going thru publishing the mobile application to the google play store
- Add something on github because all of my project end up on private repository of my company



## Build 
```
make 
```
Builds desktop app preview

```
make build-go
```
Runs go build whitout building assets 

```
make build-android
```
Builds Android APK

```
make build-assets
```
Generate's assets from ./assets dir to internal/animali/assets.go

```
make build-apk
```
Builds only APK whitout regenerating assets

## TO DO : 
 - add grid with animals sounds
 - add sound and images for music and animal screens
 - refactor
 - test
 - add language support

## Nice To Have :
 - Add badtime story's
 - Add white noises