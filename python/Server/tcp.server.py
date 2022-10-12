import socket
def server_program():
    host = "127.0.0.1"
    port = 5000  
    server_socket = socket.socket() 
    server_socket.bind((host, port))  

    server_socket.listen(2)
    
    i=1
    while True:
            print("count" + str(i));
            conn, address = server_socket.accept()
            print("Connection from: " + str(address))
            try:
                data = conn.recv(1024).decode()
            except:
                data="None"
            print(f"DATA :{data}")
            
            i+=1
        
    conn.close()


if __name__ == '__main__':
    server_program()