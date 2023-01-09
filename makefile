build : build-assets build-go
build-android: build-assets build-apk
build-go:
	go build ./cmd/animali
build-assets:
	rm -rf ./internal/animali/assets.go
	fyne bundle --pkg animali ./assets/ >> ./internal/animali/assets.go
build-apk: 
	cd cmd/animali && fyne package -os android -appID com.animali.test -icon ../../assets/icon.png
	cd cmd/animali && my -f ../../apk/animali.apk
