prepare:
	go get github.com/gen2brain/beeep
	go get github.com/go-vgo/robotgo
	go get fyne.io/fyne/v2/app

build:
	@-$(MAKE) prepare
	go build

build-windows:
	GOOS=windows \
	   GOARCH=amd64 \
	   CGO_ENABLED=1 \
	   CC=x86_64-w64-mingw32-gcc \
	   CXX=x86_64-w64-mingw32-g++ \
	   CGO_CXXFLAGS="-static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic" \
	   CGO_CFLAGS="`go env CGO_CFLAGS` -I/usr/local/x86_64-w64-mingw32/include -I/usr/local/opt/zlib/include" \
	   CGO_LDFLAGS="`go env CGO_LDFLAGS` -L/usr/local/x86_64-w64-mingw32/lib -L/usr/local/x86_64-w64-mingw32/lib -L/usr/local/opt/zlib/lib" \
	   go build -x -trimpath -ldflags="-w -s -extldflags -static"

clean:
	rm -rf no-blank no-blank.exe
	go clean