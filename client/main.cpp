#include <iostream>
#include <string>
#include <sys/socket.h> // For socket functions
#include <netinet/in.h> // For sockaddr_in and related structures
#include <arpa/inet.h>  // For inet_pton
#include <unistd.h>     // For close

#define SERVER_IP_ADDRESS "0.0.0.0"
#define SERVER_PORT 12345


int main() {
    int clientSocket = socket(AF_INET, SOCK_DGRAM, 0);
    if (clientSocket == -1) {
        // Handle error
        std::cerr << "Failed to create socket." << std::endl;
        return 1;
    }
    sockaddr_in serverAddress;
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_port = htons(SERVER_PORT); // Convert port to network byte order
    inet_pton(AF_INET, SERVER_IP_ADDRESS, &serverAddress.sin_addr); // Convert IP string to binary form

    std::string message = "Hello from UDP client!";
    ssize_t bytesSent = sendto(clientSocket, message.c_str(), message.length(), 0,
                               (struct sockaddr*)&serverAddress, sizeof(serverAddress));
    if (bytesSent == -1) {
        // Handle error
        std::cerr << "Failed to send data." << std::endl;
        close(clientSocket);
        return 1;
    }
    std::cout << "Sent " << bytesSent << " bytes: " << message << std::endl;

    char buffer[1024];
    socklen_t serverAddressLen = sizeof(serverAddress);
    ssize_t bytesReceived = recvfrom(clientSocket, buffer, sizeof(buffer), 0,
                                     (struct sockaddr*)&serverAddress, &serverAddressLen);
    if (bytesReceived == -1) {
        // Handle error
        std::cerr << "Failed to receive data." << std::endl;
    } else {
        buffer[bytesReceived] = '\0'; // Null-terminate the received data
        std::cout << "Received " << bytesReceived << " bytes: " << buffer << std::endl;
    }

    close(clientSocket);
    return 0;
}
