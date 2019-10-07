# Light Flame

Ligth flame is a modern Era ultra **light height web framework** based on netty and made for people who  like to have more control over the application, since the input of the data entrance, to the output. Everything is highly configurable... 

If you want to configure all by yourself, and have all control of your code, business logical without infer directly with the framework, so, light flame its perfectly made for you. We also encourage you to use it In a way that your code could be totally readable and decouple from the infra parts.


- low level, low abstraction , high performance, fast compile and start time
- Composite everywhere and everything, **its functions babe!**
- microservices mindset, without application services (no tomcat, jboss....)
- high configurability, before and during running time
- testing everything, intregrate tests as pure functions
- no mixing between domain and infra
- reactive non blocking IO
- non reflections, annotations free (thanks god!!)
- do whatever you want, we are on the ground =P

 
You can have some of our boiler plates example of projects using our engine and some principles of DDD, and microservices structure. I hope you enjoy, use it and help us to improve even more this simple code.

## Instalation

Using maven, declare dependency:
```maven
<dependency>
	<groupId>com.github.light-flame</groupId>
	<artifactId>lightflame-core</artifactId>
</dependency>
```

## Quick Start


```java
${../flame-examples/misc/src/main/java/init/App.java}
```

create a class that contain the configuration function, in this example, HandleConfig:
```java
${../flame-examples/misc/src/main/java/init/HandlerConfig.java}
```
Now you can declare the simple handler:
```java
${../flame-examples/misc/src/main/java/init/Handler.java}
```

## Routing rules

There is a bunch of existing rules that you can to route to your handler. The **FlameHttpStore** provides a way to store your functions and generate the rule for all.

```java
${../flame-examples/misc/src/main/java/routing/HandlerConfig.java}
```

## Flame Stores

The stores is the core API to designate the entrance of data on API. Lightflame supports:

* Http layer using **FlameHttpStore**
* Websocket layer using **FlameWsStore**
* NSQ message queue declaring on main lightflame class (see example above)
* Apache Kafka [issue #1]
* TCP server [issue #2]

(other supports on demmand. open issue if needed)

## Flame<IN,OUT> chainning API
This is one of the most important thing about lightflame. Everything is function, so every endpoint that you declare you can concatenate one or more function to this handler.

For example, to pass a http function to **FlameHttpStore** you have to declare a Flame function with **FlameHttpContext** as input and a **FlameHttpResponse** as a output. But you can chain multiple functions on the middle that connect the input with the output. Look at this example below:

```java
${..//flame-examples/websocket/src/main/java/com/ws/StaticHandler.java}
```

In this one you have this ordering:

FlameHttpContext -> java.io.File -> FlameHttpResponse

So, you can chain function **proccess()** either 4 or 5!

## Testing

Lightflame provides a simple way to test you application. You can test either if the route works depending on request, and all the steps throw the route.  

```java
${../flame-examples/helloworld/src/test/java/com/helloworld/TestingHelloWorld.java}
```
## Multi port

You can open multiple ports to work. and declare different or the same function to each port. Look at this simple example how it works. Its simple as look like:

```java
${../flame-examples/misc/src/main/java/multihttp/App.java}
```

then when you configure the project you just have to declare the port on constructor of **FlameHttpStore** like this:

```java
${../flame-examples/misc/src/main/java/multihttp/HandlerConfig.java}
```

## Web Socket and Static files (2 in 1)

You can access the full example on **flame-examples** repository on github. In this one we openned two different ports, one to serve the static files and another to WS, but you can do it using the same port

```java
${..//flame-examples/websocket/src/main/java/com/ws/App.java}
```
In this example we:

1. add a listener on port 8080 to static file
2. add a listener on port 8081 to websocket

```java
${..//flame-examples/websocket/src/main/java/com/ws/Config.java}
```

3. Its important to use the same port declared in the main file to instanciate the stores. If you dont do it the handler will throw an error.

4. the first handler execute two functions on context "/" (getRootFile() -> proccess()). This one is responsible to get the root index.html file and the second one process the file

5. this handler execute function on a widecard context after "/static/*". This happens to bring all other files that html will call, like style and js. In this one we used another function as the first one, but we mantain the second one, since the process is the same.

```java
${..//flame-examples/websocket/src/main/java/com/ws/StaticHandler.java}
```

Now we have the static hander with our three functions. The last one is the main function that receive as parameter a File and return the final **FlameHttpResponse**. The first two functions are called from different routes to read the file depending on context.

```java
${..//flame-examples/websocket/src/main/java/com/ws/WsHandler.java}
```

at the end we have our final handler for websocket app.