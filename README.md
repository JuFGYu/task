# Task

## Requirements:

You need to have **GO SDK** Installed


## STEPS TO START THE PROJECT
1. Open the terminal.
2. Go to the file location using "cd", For example:
```
cd Downloads/Task
```
3. Once your are in the file location, run the file 
```
go run task.go
```
4. Windows Defender will ask you to give permission to access the network, just press "Allow Access".
5. Open another terminal and start using the service 
```
curl "http://localhost:5000"
```
Note: if the curl happens to give you the next error:
```
curl : The response content cannot be parsed because the Internet Explorer engine is not available, or 
Internet Explorer's first-launch configuration is not complete. Specify the UseBasicParsing parameter 
and try again. 
At line:1 char:1
+ curl "http://localhost:5000"
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   + CategoryInfo          : NotImplemented: (:) [Invoke-WebRequest], NotSupportedException
   + FullyQualifiedErrorId : WebCmdletIEDomNotSupportedException,Microsoft.PowerShell.Commands.InvokeW  
   ebRequestCommand
```
You can use this command in the terminal
```
Remove-item alias:curl
```

Thanks,
