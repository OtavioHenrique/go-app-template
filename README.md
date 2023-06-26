Go App Example
=======

A simple template of Golang Application.

## What is this?

This is a simple Golang project template, applying concepts of Ports & Adapters (or hexagonal) architecture. It doesn't do anything, nor does it have the intention of being used. **The only purpose is to be a reference to how I personally structure Golang projects.** 

## How it works

### CLI

Every project is a CLI too, even if it is web-based app. To implement CLI logic, this project uses [cobra](https://github.com/spf13/cobra), which take care of all logic related to CLI. If the project is just an API like this, code a `run` command and call it directly on Dockerfile. 

### Architecture

It uses concepts of Ports & Adapters pattern, and all adapters implement interfaces that are previous declared on `ports` folder. Port interfaces can stay on adapters file that implements it (If only one adapter uses it), but I left it alone in `ports` because is easier to understand if you are not familiar with it.

Dependency injection is used on everything, so, everything used should be passed using it. All dependencies will be initialized on CLI command that uses it, this technique will make your application decoupled and easy to test. **Observability and config will be passed to almost everything, as needed.**

All adapters' initialization will occur on CLI command that need it, in this way is simple to test if needed, or just change between adapters on migration scenarios.

Application (`application.go`), will orchestrate all logic and calls, long-lived background tasks (like consuming messages) must be initialized on it (preferably on goroutines). After everything is done, you should start your API or any blocking task (like the example).

### API

It uses HTTP Api as example, because almost everything need a `healthcheck` endpoint on modern applications infrastructure, but you can remove it if you want.

### Observability 

This template assumes that every application will have a modern logger library (it uses [zap](https://github.com/uber-go/zap)) which all logic will be placed under `logger.go` and some logic to publish metrics (like [prometheus](https://github.com/prometheus/client_golang)).

All adapters and application will receive it and pass ahead. Is expected to almost all classes receives `logger` on its constructor and `metrics` when needed.

**Both observability objects should not be mutated, or global accessed. Use dependency injection.**

### Config

Config will always be the first thing to be initialized (like on run command on this template), and it will be responsible to load all config needed to the application run, including envs, secrets, everything! Many objects will have reference to this Config, so it will be passed by pointers to almost every one that needed it. 

**This object should not be mutated, and you need to pay extra attention on it. Don't access it globally**  

## How to use 

Just use it as a base to study or to use. You can clone or fork it if you want.