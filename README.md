# Port Forwarding with Go

This is a simple Go program that allows you to perform port forwarding, similar to the `socat` command. It listens on a specified local port and forwards incoming connections to a target host and port.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository and navigate to the project directory.

3. Build the program using the following command:

   ```bash
   go build
   ```

4. Run the program with the following command:

   ```bash
   ./pforward_v1 -localPort=<local-port> -targetAddr=<target-address>
   ```

   Replace `<local-port>` with the desired local port number and `<target-address>` with the target host and port (e.g., `google.com:443`).

5. The program will start listening on the specified local port and forwarding incoming connections to the target address.

6. You can now connect to the local port to access the forwarded traffic.

## Example

To forward traffic from local port 8080 to `google.com:443`, run the following command:

```bash
./pforward_v1 -localPort=8080 -targetAddr=google.com:443
```

Now, any connection made to `localhost:8080` will be forwarded to `google.com:443`.

## Dependencies

This program uses the standard Go packages for networking.

## License

This project is licensed under the [MIT License](LICENSE).
```

Feel free to modify the content and structure of the `README.md` file according to your specific needs.
