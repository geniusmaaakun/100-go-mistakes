package main

import "context"

func main() {
	// mainが終了する前に、newWatcher() が終了するようにする
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newWatcher(ctx)

	// Run the application
}

func newWatcher(ctx context.Context) {
	w := watcher{}
	go w.watch(ctx)
}

type watcher struct { /* Some resources */
}

func (w watcher) watch(context.Context) {}
