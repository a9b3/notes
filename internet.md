#Protocol

On a given computer an outgoing request goes through these layers in this order, and an incoming
request goes in the reverse order.

- Application Protocol: Specific to applications
- Transmission Control Protocol: Direct packets to port numbers
- Internet Protocol: Direct packet to computer using IP address
- Hardware Layer: converts binary to network signal


#Internet Infra

Top Down

1. NSP - Network service providers, exchange packet traffic connects to NAP (Network Access Points) or MAEs
2. IX - Internet Exchange Points
    - NAP - Send packets between NSPs
    - MAE - Send packets between NSPs (privately owned)
3. ISP - Use bandwidth from NSPs to traffic packets
4. Routers - Connected between networks to route packets. Each router knows about it's sub-network's IP addresses. The router does a lookup in a routing table and if the IP is found it sends the packet there if not it will send the packet up the chain of routers.

# DNS

A distributed database hosted by many servers. Each one has a table of the computer's name and their IP address. Any one DNS server has a subset of the entire database, if the current server does not have the domain name requested it will re-direct to another DNS server. Local network has to set up the primary DNS server to connect to first, without it it won't know where to send the DNS requests to and the DNS request will go unresolved.

# Protocols

### HTTP

Connectionless text based protocol. Client and servers send request back and forth through this protocol.

1. If url is domain name, connects to DNS to retrieve IP address of domain name.
2. HTTP request goes out  to the IP address received from the DNS.
3. Server receives request, processes the request, and sends appropriate response back.
4. Host computer receives response and the connection is closed.

### TCP

TCP attaches a TCP header to the packet as well as manage chunks and recompiling chunks. Right under the application layer is the TCP layer. Which is responsible to telling the packet where to go in a given machine, this is done through port numbers.

### IP

The Internet Protocol attaches an IP header to the packet.


# Wrapup

1. URL inputted into browser, packet needs to resolve DNS looks at OS hosts files.
2. Hosts entry not found in local machine, packet is sent to DNS server configured by OS's internet settings.
3. DNS server either knows the Host entry or it doesn't in which case it will send request to parent or another node.
4. Once DNS is found the IP address is sent back to the client.
5. Client takes the IP address and sends a "TCP SYN" request.


#### Sources
http://web.stanford.edu/class/msande91si/www-spr04/readings/week1/InternetWhitepaper.htm
