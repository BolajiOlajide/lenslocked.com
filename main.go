package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

const styles string = `<style>
	.transition, ul li i:before, ul li i:after, p {
	  transition: all 0.25s ease-in-out;
	}

	.flipIn, ul li, h1 {
	  animation: flipdown 0.5s ease both;
	}

	.no-select, h2 {
	  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
	  -webkit-touch-callout: none;
	  -webkit-user-select: none;
	  -khtml-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}

	html {
	  width: 100%;
	  height: 100%;
	  perspective: 900;
	  overflow-y: scroll;
	  background-color: #dce7eb;
	  font-family: "Titillium Web", sans-serif;
	  color: rgba(48, 69, 92, 0.8);
	}

	body {
	  min-height: 0;
	  display: inline-block;
	  position: relative;
	  left: 50%;
	  margin: 90px 0;
	  transform: translate(-50%, 0);
	  box-shadow: 0 10px 0 0 #ff6873 inset;
	  background-color: #fefffa;
	  max-width: 450px;
	  padding: 30px;
	}
	@media (max-width: 550px) {
	  body {
		box-sizing: border-box;
		transform: translate(0, 0);
		max-width: 100%;
		min-height: 100%;
		margin: 0;
		left: 0;
	  }
	}

	h1, h2 {
	  color: #ff6873;
	}

	h1 {
	  text-transform: uppercase;
	  font-size: 36px;
	  line-height: 42px;
	  letter-spacing: 3px;
	  font-weight: 100;
	}

	h2 {
	  font-size: 26px;
	  line-height: 34px;
	  font-weight: 300;
	  letter-spacing: 1px;
	  display: block;
	  background-color: #fefffa;
	  margin: 0;
	  cursor: pointer;
	}

	p {
	  color: rgba(48, 69, 92, 0.8);
	  font-size: 17px;
	  line-height: 26px;
	  letter-spacing: 1px;
	  position: relative;
	  overflow: hidden;
	  max-height: 800px;
	  opacity: 1;
	  transform: translate(0, 0);
	  margin-top: 14px;
	  z-index: 2;
	}

	ul {
	  list-style: none;
	  perspective: 900;
	  padding: 0;
	  margin: 0;
	}
	ul li {
	  position: relative;
	  padding: 0;
	  margin: 0;
	  padding-bottom: 4px;
	  padding-top: 18px;
	  border-top: 1px dotted #dce7eb;
	}
	ul li:nth-of-type(1) {
	  animation-delay: 0.5s;
	}
	ul li:nth-of-type(2) {
	  animation-delay: 0.75s;
	}
	ul li:nth-of-type(3) {
	  animation-delay: 1s;
	}
	ul li:last-of-type {
	  padding-bottom: 0;
	}
	ul li i {
	  position: absolute;
	  transform: translate(-6px, 0);
	  margin-top: 16px;
	  right: 0;
	}
	ul li i:before, ul li i:after {
	  content: "";
	  position: absolute;
	  background-color: #ff6873;
	  width: 3px;
	  height: 9px;
	}
	ul li i:before {
	  transform: translate(-2px, 0) rotate(45deg);
	}
	ul li i:after {
	  transform: translate(2px, 0) rotate(-45deg);
	}
	ul li input[type=checkbox] {
	  position: absolute;
	  cursor: pointer;
	  width: 100%;
	  height: 100%;
	  z-index: 1;
	  opacity: 0;
	}
	ul li input[type=checkbox]:checked ~ p {
	  margin-top: 0;
	  max-height: 0;
	  opacity: 0;
	  transform: translate(0, 50%);
	}
	ul li input[type=checkbox]:checked ~ i:before {
	  transform: translate(2px, 0) rotate(45deg);
	}
	ul li input[type=checkbox]:checked ~ i:after {
	  transform: translate(-2px, 0) rotate(-45deg);
	}

	@keyframes flipdown {
	  0% {
		opacity: 0;
		transform-origin: top center;
		transform: rotateX(-90deg);
	  }
	  5% {
		opacity: 1;
	  }
	  80% {
		transform: rotateX(8deg);
	  }
	  83% {
		transform: rotateX(6deg);
	  }
	  92% {
		transform: rotateX(-3deg);
	  }
	  100% {
		transform-origin: top center;
		transform: rotateX(0deg);
	  }
	}
	</style>
`

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>Lens Locked</title>
		%s
	</head>
	<body>
		<h1>Welcome to my pretty awesome site!</h1>
	</body>
</html>
`, styles))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>Lens Locked | Contact</title>
		%s
	</head>
	<body>
		<h1>Contact page</h1>
		<p>To get in touch email me at <a href="mailto:bolaji@lenslocked.com">bolaji@lenslocked.com</a>
	</body>
</html>
`, styles))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>Lens Locked | FAQs</title>
		%s
	</head>
	<body>
		<h1>Lens Locked FAQ</h1>
		<ul>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>Is there a free version?</h2>
			<p>Yes! We have a free trial for 30 days on any paid plans.</p>
		  </li>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>What are your support hours?</h2>
			<p>We have support staff answering emails 24/7, though response time may be a bit slower on weekends.</p>
		  </li>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>How do I contact support?</h2>
			<p>Email us - <a href="support@lenslocked.com">support@lenslocked.com</a></p>
		  </li>
		</ul>
	</body>
</html>
`, styles))
}

func getSingleResourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userID := chi.URLParam(r, "userID")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>LensLocked | Single Resource</title>
		%s
	</head>
	<body>
		<h1>I am getting a single resource</h1>
		<p><b>userID</b>: %s<p>
	<body>
</html>
`, styles, userID))
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Page Not Found", http.StatusNotFound)
//	}
//}

//type Router struct {}
//
//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Page Not Found", http.StatusNotFound)
//	}
//}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.CleanPath, middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", homeHandler)
	r.Get("/user/{userID}", getSingleResourceHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	// not necessarily needed but okay to replicate former router structure
	r.NotFound(func (w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)

	//http.Handle("/", http.HandlerFunc(homeHandler))

	fmt.Println("Starting the server on port 4000")
	http.ListenAndServe(":4000", r)
}
