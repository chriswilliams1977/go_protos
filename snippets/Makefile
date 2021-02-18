#-------HEADERS--------
#Options to manipulate headers
#Set a new cache-control header. If an existing "Cache-Control" header exists
#it will be overwritten.
#w.Header().Set("Cache-Control", "public, max-age=31536000")

#In contrast, the Add() method appends a new "Cache-Control" header and can
#be called multiple times.
#w.Header().Add("Cache-Control", "public")
#w.Header().Add("Cache-Control", "max-age=31536000")

#Delete all values for the "Cache-Control" header.
#w.Header().Del("Cache-Control")

#Retrieve the first value for the "Cache-Control" header.
#w.Header().Get("Cache-Control")

#Go send threesystem-generated headers Date and Content-Length and Content-Type
#Content-Type - uses http.DetectContentType() to detect content type otherwise default Content-Type: application/octet-stream
#cannot detect JSON from plain text
#can fix this using
#w.Header().Set("Content-Type", "application/json")
#w.Write([]byte(`{"name":"Alex"}`))
#header name always canonicalized - converts first letter and any letter following - to uppercase
#can override using - w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
#When headers are written to a HTTP/2 connection the header names and values will
#always be converted to lowercase
# Del() method doesn’t remove system-generated headers
#-------HEADERS--------

#-------PROJECT STRUCTURE------
#Typically
#cmd/web for app logic
#cmd/pkg for created packages
#ui/html and ui/static for templates, HTML, CSS, images..
#Tim Hockin set up - https://github.com/thockin/go-build-template
#-------PROJECT STRUCTURE------

#------ TEMPLATING------------
#Go’s net/http package ships with a built-in http.FileServer handler which you can use to
#serve files over HTTP from a specific directory.
#fileServer := http.FileServer(http.Dir("./ui/static"))
#-------TEMPLATING-------------

#-------REQUESTS---------------
#can occur when two or more goroutines try to use a piece of shared data (in this example the bank balance) at the same time,
#but the result of their operations is dependent on the exact order that the scheduler executes their instructions.
#-------REQUESTS---------------

#------CONFIGURATION-----------
#Go use conmmand line flags to set config
#addr := flag.String("addr", ":4000", "HTTP network address")
#set value in cmd lind
	#go run cmd/web/* -addr=":9999"
	#use -help to list cmd line flags
	#go run cmd/web/* -help
	#can alos use env vars
	#addr := os.Getenv("SNIPPETBOX_ADDR")
#but they do not have help or defaults
#can workaround this by passing the environment variable as a
#command-line flag when starting the application
	#export SNIPPETBOX_ADDR=":9999"
	#go run cmd/web/* -addr=$SNIPPETBOX_ADDR
#you can also store config in a struct in memory using flag.StringVar(),flag.IntVar(), flag.BoolVar()
	#type Config struct {
	#Addr string
	#StaticDir string
	#}

	#cfg := new(Config)
	#flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	#flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	#flag.Parse()
#------CONFIGURATION-----------

#------LOGGING-----------
#avoid using the Panic() and Fatal() variations outside of
#your main() function — it’s good practice to return errors instead, and only panic or exit
#directly from main().

#Custom loggers created by log.New() are concurrency-safe. You can share a single logger
#and use it across multiple goroutines and in your handlers without needing to worry about
#race conditions.

#recommendation is to log your output to standard streams and
#redirect the output to a file at runtime.
#Can do this in go in you prefer
	#f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	#if err != nil {
	#log.Fatal(err)
	#}
	#defer f.Close()
	#infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
#------LOGGING-----------


project_build:
	mkdir -p cmd/web pkg ui/html ui/static
	touch cmd/web/main.go
	touch cmd/web/handlers.go
build_local:
	#build and tag image
	docker build -t snippets-example .
	#ADD Image to GCR
	#docker tag snippets-example gcr.io/williamscj-gke-demos/snippets-example:v1
	#docker push  gcr.io/williamscj-gke-demos/snippets-example:v1
create_source_gsr:
	git config --global credential.'https://source.developers.google.com'.helper gcloud.sh
	gcloud source repos create snippets
	git remote add google \
	https://source.developers.google.com/p/williamscj-demos/r/snippets
create_af_repo:
	#Set up artifact registry repo for images
	#gcloud artifacts locations list - gets list of supported regions
	gcloud artifacts repositories create snippets \
    --repository-format=docker --location=europe-west4 \
    --description=snippets_demo --async
set_af_permissions:
	#NOT WORKING
	#get list of existing policies
	#gcloud projects get-iam-policy williamscj-demos
	gcloud projects set-iam-policy williamscj-demos policy.yaml
build_af:
	#Push image
	#Uses Artifact Registry on GCP
	#gcloud auth configure-docker europe-west4-docker.pkg.dev  - allows docker to push to artifact repo
	#tag image
	docker tag snippets-example europe-west4-docker.pkg.dev/williamscj-demos/snippets/snippets-image:v1
	#push image
	docker push europe-west4-docker.pkg.dev/williamscj-demos/snippets/snippets-image:v1
deploy_gke:
	#To view get svc LB external IP
	kubectl apply -f go-deploy.yaml
svc_gke:
	kubectl apply -f go-service.yaml