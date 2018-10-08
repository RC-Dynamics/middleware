package main

// Main Server
func main() {
	address := ":8080"
	client_proxy := ClientProxy{address}

	for {
		invoke(client_proxy)
		// fmt.Fprintf(os.Stderr, "%s\n", data)
	}
}
