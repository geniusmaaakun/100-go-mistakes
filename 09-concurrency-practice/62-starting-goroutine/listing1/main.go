package main

func main() {
	newWatcher()

	// Run the application

	// newWatcher() が終了する前に、main() が終了してしまう
}

func newWatcher() {
	w := watcher{}
	go w.watch()
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}
