import http.server
import socket
import socketserver
import sys
import urllib.request
import json
import atexit

if len(sys.argv) != 3:
    print("Usage: python test_service.py <PORT> <SERVICE_NAME>")
    sys.exit(1)

PORT = int(sys.argv[1])
SERVICE_NAME = sys.argv[2]

class Handler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(f"Hello from {SERVICE_NAME}!\n".encode())

def get_host_ip():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect(("8.8.8.8", 80))
    host_ip = s.getsockname()[0]
    s.close()
    return host_ip

def unregister_service():
    unregister_data = {
        "service_name": SERVICE_NAME,
        "available_at": f"http://localhost:{PORT}"
    }
    unregister_data_json = json.dumps(unregister_data).encode('utf-8')
    req = urllib.request.Request("http://localhost:8081/unregister", data=unregister_data_json, method='POST')
    try:
        response = urllib.request.urlopen(req)
        if response.getcode() == 200:
            print("Unregistration successful")
        else:
            print("Unregistration failed")
    except urllib.error.URLError as e:
        print("Unregistration failed:", e)

atexit.register(unregister_service)

with socketserver.TCPServer(("0.0.0.0", PORT), Handler) as httpd:
    print(f"Starting {SERVICE_NAME} on port {PORT}...")

    # Send registration request to localhost:8081
    registration_data = {
        "service_name": SERVICE_NAME,
        "available_at": f"http://localhost:{PORT}"
    }
    registration_data_json = json.dumps(registration_data).encode('utf-8')
    req = urllib.request.Request("http://localhost:8081/register", data=registration_data_json, method='POST')
    try:
        response = urllib.request.urlopen(req)
        if response.getcode() == 200:
            print("Registration successful")
        else:
            print("Registration failed")
    except urllib.error.URLError as e:
        print("Registration failed:", e)

    httpd.serve_forever()
