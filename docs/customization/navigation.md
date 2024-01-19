---
index: 4
title: Navigation Order
---
When its not specified the order of the pages is alphabetical. However, you can change the order of the pages by adding an `index` property to the folders and file. The `index` property is a number that defines the order of the page in the navigation. The lower the number, the higher the page will be in the navigation.

## Folders

To order the folders, you need to add an `index` property to the `_meta.md` file of the folder. The `index` property is a number that defines the order of the folder in the navigation. The lower the number, the higher the folder will be in the navigation.

## Files

In files you can add an `index` property to the `_meta.md` file of the file. The `index` property is a number that defines the order of the file in the navigation. Same rule applies: The lower the number, the higher the file will be in the navigation.

## Index vs Non-Index documents and folders
When the index is not specified, those items are assigned an index of 10_000_000. This is to ensure that the index is always higher than the index of the items that are specified. This is a way to ensure that the items that are not specified are always at the bottom of the navigation.