---
index: 3
---

External links are a good way to connect your business with your documentation. Doco allows you to add a set of links to the navigation bar. This is a great way to connect your documentation with your product. 

The external links are configured through the `_meta.md` file at the root of the documentation folder. you can change it through the following options:

```yml
external_links:
  - text: "Pricing"
    link: "https://doco.paganotoni.com/pricing"
  - text: "Contact"
    link: "https://doco.paganotoni.com/contact"
```

You can add as many links as you want. The links will appear in the order they are defined in the `_meta.md` file. so one important consideration is space. If you add too many links, they will overflow the navigation bar.