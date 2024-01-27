---
title: Cloudflare Pages
---

Cloudflare Pages is a great way to deploy your static website. It's free and it comes with a lot of features like custom domains, automatic SSL certificates, and more. To deploy your site from github to Cloudflare Pages you will need to create a new page in Cloudflare Pages. 

Once that's done you can add a new file to your repository called `.github/workflows/cloudflare-pages.yml` with the following content:


```yaml
name: Cloudflare Pages
on: 
  push:
    branches:
      - main

jobs:
  Build:
    name: build docs
    runs-on: ubuntu-latest
    environment:
      name: github-pages
      url: ${{ steps.deployment.outputs.page_url }}
    steps:
      - uses: actions/checkout@v2
      - name: Setup Go ${{ matrix.go-version }}
        uses: actions/setup-go@v4

      - name: install doco
        run: go install github.com/paganotoni/doco/cmd/doco@latest

      - run: doco build
      - name: Publish to Cloudflare Pages
        uses: cloudflare/pages-action@v1
        with:
          apiToken: ${{ secrets.CLOUDFLARE_API_TOKEN }} 
          accountId: ${{ secrets.CLOUDFLARE_ACCOUNT_ID }}
          projectName: ${{ secrets.CLOUDFLARE_PROJECT_NAME }}
          directory: ./public

```

This will trigger a build every time you push to the `main` branch. If you want to trigger a build on a different branch you can change the `branches` section of the `on` property. After you've added the file, commit it to your repository and github will take care of deploying it.

This action depends on 3 secrets that you will need to add to your repository. You can find the `CLOUDFLARE_API_TOKEN` by going to your Cloudflare account settings and clicking on the API Tokens tab. the `CLOUDFLARE_ACCOUNT_ID` can be found in the same page. Finally, the `CLOUDFLARE_PROJECT_NAME` can be found in the Cloudflare Pages dashboard.

These need to be set in your repository settings under the secrets tab. You can find more information about how to add secrets to your repository [here](https://docs.github.com/en/actions/reference/encrypted-secrets#creating-encrypted-secrets-for-a-repository).
