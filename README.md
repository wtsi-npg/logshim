# logshim

## Overview

`logshim` is a [shim](https://en.wikipedia.org/wiki/Shim_\(computing\)) API to
provide library-independent logging statements.

It provides a common denominator of API calls to logging backends for the most
frequent operations i.e. log statements that will be found throughout code. It
does not attempt to abstract away tasks such as log configuration, which are
left to each backend's own API.
