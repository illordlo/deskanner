# theskanner
distributed scanner for wide networks

# disclaimer
as you can see in the "about" page, this project has been created, developed, tested and deployed in 40 minutes, so you are using it at your on risk. moreover, **the example is running a command as root using as argument the content provided by the server**. again, even if HTTPS with certificate pinning is used and the input is lightly sanitized, you are doing it your one risk.

that said, you are fully responsible for using, misusing and blablabla of this software.

## description
when you want to scan huge networks, it takes hours (or maybe days?) to do it, if you do it with a single machine.
luckly, friends (or comminities, or pwned machines) are always helpfull on this kind of project so you can simply start the server with the right parameters and ask them to run the client.

## how the client works
the client is basically a stupid shell script doing the following stuffs:
1. make an http GET request to the server, asking for a range to scan
2. actually scan the IP range
3. upload the scan result to the storage server (in XML format, encoded in base64)

the example you can see here just scan for UPD 500 and 4500 ports. to accomplish this task, it has to run as root. sorry for that.

## how the server works
the server is written in golang (thanks dyst0ni3 and bestbug for writing it) and it works as follows:
1. the server is started passing the starting IP range, the ending IP range and the size of the subranges
2. the server creates the nmap strings required for scanning the subranges
3. the server returns every string just once to every connecting client

## how the storage server works
we absolutely don't care. you can use a stupid netcat socket, amazon s3 or just watch at the raw tcp packets reaching you server and manually write the resulting payload on a piece of paper.

## license
i don't care. rofl.