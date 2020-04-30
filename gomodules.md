This file contains important aspects of go modules

How Go GET WORKS
================
When you install any dependency packages using go get command, it saves the package files under $GOPATH/src path. 
Let’s understand how go get works and how to host your own custom Go repository.

go get domain.com/path/to/sub/directory
Go visits website domain.com/path/to/sub/directory securely using the HTTPS protocol. If the given URL does not support HTTPS or results in SSL error, and if GIT_ALLOW_PROTOCOL environment variable contains HTTP protocol, 
then Go attempts to resolve the package URL with HTTP protocol.
A Go package located on the internet is a version control system repository (VCS) like GIT or SVN. Supported version control systems are mentioned below.
Bazaar        .bzr
Fossil        .fossil
Git           .git
Mercurial     .hg
Subversion    .svn

If the domain.com is one of the code hosting sites recognized by Go, Go first tries to resolve domain.com/path/to/sub/directory.{type} where type can be .git, .svn etc. supported by the recognized website mentioned below.
Bitbucket              (bitbucket.org)     .git/.hg
GitHub                 (github.com)        .git
Launchpad              (launchpad.net)     .bzr
IBM DevOps Services    (hub.jazz.net)      .git

If Go successfully resolves the above URL, the VCS repository is cloned (using VCS tools like Git) and saved 
to the relevant path in $GOPATH/src as discussed earlier.

ISSUES in GO GET
================
A Go program cannot import a dependency unless it is present inside $GOPATH. Also, 
go build command creates binary executable files and package archives inside $GOPATH. Hence, GOPATH is a big deal in Go.

you can’t install multiple versions of the same dependency package since go get puts the dependency package at the same location despite having different versions
(as directory name is the same as the package name).

The solution for the above probems is provided by go modules.

GO MODULES
==========
A module is nothing but a directory containing Go packages. These can be distribution packages or executable packages.
A module is also like a package that you can share with other people. Hence, it has to be a Git or any other VCS repository which we can be hosted on a code-sharing platform like Github. Hence, Go recommends
  A Go module must be a VCS repository or a VCS repository should contain a single Go module.
  A Go module should contain one or more packages
  A package should contain one or more .go files in a single directory
  


