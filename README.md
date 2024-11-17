

# Hands-on Software Engineering with Golang
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ibiscum/Hands-On-Software-Engineering-with-Golang)](https://goreportcard.com/report/github.com/ibiscum/Hands-On-Software-Engineering-with-Golang)
[![lint](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/lint.yml/badge.svg)](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/lint.yml)
[![build](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/build.yml/badge.svg)](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/build.yml)
[![codeql go](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/codeql.yml/badge.svg)](https://github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/actions/workflows/codeql.yml)

<a href="https://www.packtpub.com/in/programming/hands-on-software-engineering-with-golang?utm_source=github&utm_medium=repository&utm_campaign="><img src="https://www.packtpub.com/media/catalog/product/cache/e4d64343b1bc593f1c5348fe05efa4a6/9/7/9781838554491-original.png" alt="" height="256px" align="right"></a>

This is the code repository for [Hands-on Software Engineering with Golang](https://www.packtpub.com/in/programming/hands-on-software-engineering-with-golang?utm_source=github&utm_medium=repository&utm_campaign=), published by Packt.

**Move beyond basic programming to design and build reliable software with clean code**

## What is this book about?

This book distills the industry’s best practices for writing lean Go code that
is easy to test and maintain and explores their practical application on Links
‘R’ US: an example project that crawls web-pages and applies the PageRank
algorithm to assign an importance score to each one.

This book covers the following exciting features:

* Understand different stages of the software development life cycle and the role of a software engineer
* Create APIs using gRPC and leverage the middleware offered by the gRPC ecosystem
* Discover various approaches to managing package dependencies for your projects
* Build an end-to-end project from scratch and explore different strategies for scaling it
* Develop a graph processing system and extend it to run in a distributed manner
* Deploy Go services on Kubernetes and monitor their health using Prometheus

## Instructions
All of the code is organized into folders labelled after the chapter they
appear on. For example, Chapter02 contains the source code for the second book
chapter and so on.

The Makefile has been updated to manage dependencies via Go modules instead of
the dep tool. However, the dep tool will be used as a _fall-back_ for old Go
versions (that lack module support) or if the `GO111MODULE` environment
variable is set to `off` prior to running any of the Makefile targets.

Go 1.18+ is required for running the code/tests from the individual chapters.
The latest version of Go for your platform can be downloaded
[here](https://go.dev/dl/).

We also provide a PDF file that has color images of the screenshots/diagrams
used in this book. [Click here to download
it](https://static.packt-cdn.com/downloads/9781838554491_ColorImages.pdf).

### Intended audience
This Golang programming book is for developers and software engineers looking to use Go to design and build scalable distributed systems effectively. Knowledge of Go programming and basic networking principles is required.

### Related products
* Hands-On System Programming with Go  [[Packt]](https://www.packtpub.com/application-development/hands-systems-programming-go?utm_source=github&utm_medium=repository&utm_campaign=9781789804072) [[Amazon]](https://www.amazon.com/dp/1789804078)

* Go Programming Cookbook - Second Edition  [[Packt]](https://www.packtpub.com/in/application-development/go-programming-cookbook-second-edition?utm_source=github&utm_medium=repository&utm_campaign=9781789800982) [[Amazon]](https://www.amazon.com/dp/1789800986)

## Get to Know the Author
**Achilleas Anagnostopoulos**
has been writing code in a multitude of programming languages since the mid
90s. His main interest lies in building scalable, microservice-based
distributed systems where components are interconnected via gRPC or message
queues. Achilleas has over 4 years of experience building production-grade
systems using Go and occasionally enjoys pushing the language to its limits
through his experimental [gopher-os](https://github.com/gopher-os/gopher-os)
project: a 64-bit kernel written entirely in Go. He is a former member of
the [Juju](https://jaas.ai/) team at Canonical, and has contributed to one of the
largest open source Go [code bases](https://github.com/juju/juju) in existence.
He is currently working as an SRE at Google.

### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.
### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781838554491">https://packt.link/free-ebook/9781838554491 </a> </p>
