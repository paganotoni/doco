---
index: 6
---

Beyond general customization, Doco allows to setup some document specific configurations through the meta block in the document. This block is a yaml block that can be added to the top of the document and it is used to customize the document. It looks like this:

```yml
---
title: Getting Started
description: Learn how to get started with Doco, from installation to initializing the documentation folder.
keywords: getting started, installation, doco
index: 2
---
```

When these values are not stablished doco will apply defaults or fallback to the global configuration. The following sections will explain each of the options available in the meta block.

### Title
The title of the document, is used on the pages title and in the meta tags. When this is not set at the page level it will fallback to the global configuration.

### Description 
The description of the document, is used on the pages description and in the meta tags when this is not set at the page level it will fallback to the global configuration.

### Keywords 
The keywords of the document, is used on the pages description and in the meta tags. When this is not set at the page level it will fallback to the global configuration.

### Index 
Defines where the document will be positioned in the navigation. When not set it defaults to 10k so the documents will show at the end of the navigation.
