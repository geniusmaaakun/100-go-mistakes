package main

func main() {
	w := newWatcher()
	//明示的にclose()を呼び出す
	defer w.close()

	// Run the application
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}

func (w watcher) close() {
	// Close the resources
}
