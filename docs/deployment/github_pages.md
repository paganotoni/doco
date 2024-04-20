---
title: GitHub Pages
---

One of the easiest ways to deploy your site is to use GitHub Pages. Github pages will take care of the hosting for you and it's free. Doco will generate a static website that can be hosted on GitHub Pages, then you can use a custom domain to point to your GitHub Pages site.

⚠️ This guide assumes you have a github repository already. If you don't, you can create one by following the [GitHub guide](https://docs.github.com/en/github/getting-started-with-github/create-a-repo).

## Enabling GitHub Pages

To enable GitHub Pages, go to your repository settings and scroll down to the GitHub Pages section. Click on the `Sources` dropdown and select `github actions` as the source.

![Github Pages](</assets/github-pages-source.png>)

Once you've selected `github actions` as the source we will need to create a workflow file. To do this we will modify the source adding a `.github/workflows/gh-pages.yml` file to your repository.

```yaml
name: Pages
on:
  push:
    branches:
      - main

permissions:
  contents: read
  pages: write
  id-token: write

concurrency:
  group: "pages"
  cancel-in-progress: false

jobs:
  Build:
    name: build docs
    runs-on: ubuntu-latest
    environment:
      name: github-pages
      url: ${{ steps.deployment.outputs.page_url }}
    steps:
      - uses: actions/checkout@v2
      - name: Setup Doco
        run: >
          wget https://github.com/paganotoni/doco/releases/latest/download/doco_Linux_x86_64.tar.gz &&
          tar -xvf doco_Linux_x86_64.tar.gz

      - run: ./doco build
      - name: Upload artifact
        uses: actions/upload-pages-artifact@v3
        with:
          path: 'public'
      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v4
```

This will trigger a build every time you push to the `main` branch. If you want to trigger a build on a different branch you can change the `branches` section of the `on` property. After you've added the file, commit it to your repository and github will take care of deploying it.
