String.prototype.interpolate = function (params) {
  const names = Object.keys(params);
  const vals = Object.values(params);
  return new Function(...names, `return \`${this}\`;`)(...vals);
};

document.addEventListener("DOMContentLoaded", () => {
  // Add highligthing to the code blocks
  hljs.highlightAll();

  // Adding the mobile menu toggle
  let toggles = document.querySelectorAll(".toggle-mobile-nav");
  toggles.forEach((toggle) => {
    toggle.addEventListener("click", () => {
      document.getElementById("mobile-menu").classList.toggle("hidden");
      document.querySelector("body").classList.toggle("overflow-hidden");
    });
  });

  // Replacing the $YEAR element with the current year
  let copy = document.getElementById("copy");
  copy.innerHTML = copy.innerHTML.replace("$YEAR", new Date().getFullYear());

  // TODO: Keyboard up and down
  // TODO: Keyboard enter to navigate to article
  document.querySelectorAll(".search-button").forEach((element) => {
    element.addEventListener("click", showSearch);
  });

  document
    .getElementById("search-overlay")
    .addEventListener("click", hideSearch);
  document.getElementById("close-search").addEventListener("click", hideSearch);
  document
    .getElementById("search-input")
    .addEventListener("keyup", (e) => search(e.target.value));

  document.addEventListener("keydown", () => {
    let paletteVisible = document
      .getElementById("search-palette")
      .classList.contains("hidden");

    if (paletteVisible) {
      return;
    }

    let selector = "#search-results li.selected";
    let quickLinksVisisble = document
      .querySelector("#search-quick-actions")
      .classList.contains("hidden");

    if (!quickLinksVisisble) {
      selector = "#search-quick-actions li.selected";
    }

    // get the current selected element
    let selected = document.querySelector(selector);
    if (selected == null) {
      return;
    }

    // on arrow down move the selected element down
    if (event.keyCode == 40) {
      let next = selected.nextElementSibling;
      if (next == null) {
        return;
      }

      selected.classList.remove("selected");
      next.classList.add("selected");
      next.scrollIntoView(false);
    }

    // on arrow up move the selected element up
    if (event.keyCode == 38) {
      let prev = selected.previousElementSibling;
      if (prev == null) {
        return;
      }

      selected.classList.remove("selected");
      prev.classList.add("selected");
      prev.scrollIntoView(false);
    }

    // on enter navigate to the selected element
    if (event.keyCode == 13) {
      let selected = document.querySelector("#search-results li.selected a");
      if (selected == null) {
        return;
      }

      window.location = selected.href;
    }
  });

  // Loading the index when the page it loaded so the search just uses it.
  fetch("/index.json")
    .then((response) => response.json())
    .then((data) => {
      window.searchIndex = new Fuse(data, {
        threshold: 0.8,
        tokenize: false,
        includeMatches: true,
        maxPatternLength: 32,
        minMatchCharLength: 1,
        location: 80_000,
        keys: ["content", "title", "section_name"],
      });
    })
    .catch((error) => console.error("Error:", error));

  let imageContainer = document.querySelector("#zoomed-image-overlay");
  imageContainer.addEventListener("click", () => {
    imageContainer.classList.add("hidden");
  });

  let zoomables = document.querySelectorAll("#htmlcontainer img");
  zoomables.forEach((zoomable) => {
    zoomable.classList.add("cursor-zoom-in");
    zoomable.addEventListener("click", () => {
      imageContainer.classList.remove("hidden");
      imageContainer.querySelector("img").src = zoomable.src;
    });
  });

  document.addEventListener("keydown", (e) => {
    if (e.keyCode === 27) {
      hideSearch();
      // hide image container
      imageContainer.classList.add("hidden");
    }

    if (e.keyCode >= 65 && e.keyCode <= 90) {
      let char = (e.metaKey ? "⌘-" : "") + String.fromCharCode(e.keyCode);
      if (char == "⌘-K") {
        showSearch();
      }
    }
  });
});

function showSearch() {
  document.getElementById("search-palette").classList.remove("hidden");
  document.getElementById("search-input").focus();
}

function hideSearch() {
  document.getElementById("search-palette").classList.add("hidden");
}

let tm = null;
let lastQuery = "";
function search(searchQuery) {
  if (searchQuery.length == 0) {
    document.getElementById("search-no-results").classList.add("hidden");
    document.getElementById("search-results").classList.add("hidden");
    document.getElementById("search-quick-actions").classList.remove("hidden");

    return;
  }

  if (searchQuery == lastQuery) {
    return;
  }

  lastQuery = searchQuery;

  if (tm != null) {
    clearTimeout(tm);
  }

  tm = setTimeout(() => {
    var result = searchIndex.search(searchQuery);

    // hide the quick actions
    document.getElementById("search-quick-actions").classList.add("hidden");

    if (result.length == 0) {
      document.getElementById("search-no-results").classList.remove("hidden");
      document.getElementById("search-results").classList.add("hidden");
      return;
    }

    document.getElementById("search-no-results").classList.add("hidden");
    document.getElementById("search-results").classList.remove("hidden");

    populateResults(result);
  }, 400);
}

function populateResults(result) {
  var template = document.getElementById("search-result-template").innerHTML;
  document.getElementById("search-results").innerHTML = "";

  result.forEach((value, index) => {
    if (value.item.title == null) {
      return;
    }

    const output = template.interpolate({
      key: index,
      title: value.item.title,
      link: value.item.link,
      // selecting the first one
      selected: index == 0 ? "selected" : "",
    });

    document.getElementById("search-results").innerHTML += output;
  });
}
