# go_API_CBCardDev
Follow the Structure mentioned in here 

https://medium.com/@ott.kristian/how-i-structure-services-in-go-19147ad0e6bd

/*Root*/
This is where I like to keep anything used to work with the source code that helps getting it up and running, i.e. build tools, configuration files, dependency management, etc.. It also gives the reader/developer a good starting point — everything they need to run the service is right there in the root of the project.



/*Cmd*/
Anything that I want to run as part of our service will go in here. E.g. the API, cron jobs, etc.. Essentially anything that will be compiled to a binary will have a main package here, where we initialize the configuration and packages required to start any part of the service.




/*Pkg*/
This contains the meat of the project: the packages that define our services behaviour. Let me quickly explain each of these.

api/
Here I define how to wire up the API by initializing the DB, services, HTTP router + middleware & define the configuration we need to run the API. I usually have a single Start(cfg *Config) function that gets called from cmd/api/main.go.

db/
Pretty self-explanatory — this is where the connection & migration logic live. I also tend to put any migration folders/files in here.

utils/
This is where I’ll throw any package that’s related to helping with requests, logging, custom middleware, etc.. I’m not the biggest fan of this name but I haven’t landed on anything better just yet.

services/
This one requires a bit more explanation because I structure each of these services in a specific way. In general, each of these packages defines a feature of the service (structuring by feature rather than function).
