package main

func main() {
	cfg := &config{}
	cfg.commands = getCommands()
	startRepl(cfg)
}
