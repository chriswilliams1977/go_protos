#-------HEADERS--------
#Options to manipulate headers
#Set a new cache-control header. If an existing "Cache-Control" header exists
#it will be overwritten.
w.Header().Set("Cache-Control", "public, max-age=31536000")

#In contrast, the Add() method appends a new "Cache-Control" header and can
#be called multiple times.
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")

#Delete all values for the "Cache-Control" header.
w.Header().Del("Cache-Control")

#Retrieve the first value for the "Cache-Control" header.
w.Header().Get("Cache-Control")

#Go send threesystem-generated headers Date and Content-Length and Content-Type
#Content-Type - uses http.DetectContentType() to detect content type otherwise default Content-Type: application/octet-stream
#cannot detect JSON from plain text
#can fix this using
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
#header name always canonicalized - converts first letter and any letter following - to uppercase
#can override using - w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
#When headers are written to a HTTP/2 connection the header names and values will
#always be converted to lowercase
# Del() method doesn’t remove system-generated headers
#-------HEADERS--------

#-------PROJECT STRUCTURE------
#Tim Hockin set up - https://github.com/thockin/go-build-template

#-------PROJECT STRUCTURE------

project_build:
	mkdir -p cmd/web pkg ui/html ui/static
	touch cmd/web/main.go
	touch cmd/web/handlers.go