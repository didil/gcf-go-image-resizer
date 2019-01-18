# Go Image Resizer Cloud Function

Pure Go Image Resizer function, deployable to Google Cloud Functions, using the [disintegration/imaging](https://github.com/disintegration/imaging) package

Medium article: [Image Resizing with Go and Cloud Functions](https://medium.com/@didil/image-resizing-with-go-and-cloud-functions-792a47e6473d)

*Original gopher image:*   
![Big Gopher](/example/gopherizeme_orig.jpg?raw=true "Big Gopher")

*Resized gopher image:*  
![Small Gopher](/example/gopherizeme_resized.jpg?raw=true "Small Gopher")


*Example gopher image generated using [gopherize.me](https://gopherize.me/)*

## Usage

https://{gcf-endpoint}/ResizeImage?url={url}&height={height}&width={width}
- url: url of the image to resize
- height: height of the output image in pixels
- width: width of the output image in pixels

*if width or height is missing, the aspect ratio is preserved*

## Deploying

```` 
$ gcloud functions deploy ResizeImage --runtime go111 --trigger-http
````

## Local testing
There is an http server included in cmd/server.go allowing you to test locally

```` 
$ export GO111MODULE=on
$ go get -u
$ go run cmd/server/server.go

````


## Todo
- Add tests
- Cache input images
- Cache output images