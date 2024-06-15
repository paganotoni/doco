String.prototype.interpolate = function (params) {
  const names = Object.keys(params);
  const vals = Object.values(params);
  return new Function(...names, `return \`${this}\`;`)(...vals);
};

// Doco JS functionallity
let doco = {
  formatCode: function () {
    // Add highligthing to the code blocks
    hljs.highlightAll();

    document.querySelectorAll("pre").forEach((el) => {
      let bt = document.createElement("button");
      bt.classList.add("material-symbols-outlined", "absolute", "top-3", "right-2", "p-1.5", "bg-gray-50", "hover:bg-gray-100", "rounded");
      bt.appendChild(document.createTextNode("content_copy"));
      bt.addEventListener("click", () => {
        navigator.clipboard.writeText(el.textContent);

        // Add a notice that the code was copied
        let notice = document.createElement("div");
        notice.appendChild(document.createTextNode("Copied!"));
        notice.classList.add("absolute","top-3","right-2","p-2","bg-gray-50","rounded","text-sm");
        el.appendChild(notice);

        // Remove the notice after 1.5 seconds
        setTimeout(() => {
          notice.remove()
        }, 1500);
      });

      el.appendChild(bt);


      el.classList.add("relative");
    });
  },

  replaceYear: function () {
    // Replacing the $YEAR element with the current year
    let copy = document.getElementById("copy");
    copy.innerHTML = copy.innerHTML.replace("$YEAR", new Date().getFullYear());
  },

  search: {
    tm: null,
    lastQuery: "",
    index: null, // The search index
    setup: function () {
      // Loading the index when the page it loaded so the search just uses it.
      fetch("/index.json").then((response) => {
        return response.json()
      }).then((data) => {
        doco.search.index = new Fuse(data, {
          threshold: 0.8,
          tokenize: false,
          includeMatches: true,
          maxPatternLength: 32,
          minMatchCharLength: 1,
          location: 80_000,
          keys: ["content", "title", "section_name"],
        });
      }).catch((error) => {
        console.error("Error:", error)
      });

      // Search button toggles the search
      document.querySelectorAll(".search-button").forEach((element) => {
        element.addEventListener("click", doco.search.show);
      });

      document.getElementById("search-overlay").addEventListener("click", doco.search.hide);
      document.getElementById("close-search").addEventListener("click", doco.search.hide);
      document.getElementById("search-input").addEventListener("keyup", (e) => {
        doco.search.do(e.target.value)
      });
      // Up and down navigation.
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

      //Esc and CMD+k toggle
      document.addEventListener("keydown", (e) => {
        if (e.keyCode === 27) {
          doco.search.hide();
          // hide image container
          imageContainer.classList.add("hidden");
        }

        if (e.keyCode >= 65 && e.keyCode <= 90) {
          let char = (e.metaKey ? "⌘-" : "") + String.fromCharCode(e.keyCode);
          if (char == "⌘-K") {
            doco.search.show();
          }
        }
      });

    },

    show: function () {
      document.getElementById("search-palette").classList.remove("hidden");
      document.getElementById("search-input").focus();
    },
    hide: function () {
      document.getElementById("search-palette").classList.add("hidden");
    },

    populateResults: function (result) {
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
    },

    do: function(searchQuery) {
      if (searchQuery.length == 0) {
        document.getElementById("search-no-results").classList.add("hidden");
        document.getElementById("search-results").classList.add("hidden");
        document.getElementById("search-quick-actions").classList.remove("hidden");

        return;
      }

      if (searchQuery == doco.search.lastQuery) {
        return;
      }

      doco.search.lastQuery = searchQuery;
      if (doco.search.tm != null) {
        clearTimeout(doco.search.tm);
      }

      doco.search.tm = setTimeout(() => {
        console.log(doco.search.index)
        var result = doco.search.index.search(searchQuery);

        // hide the quick actions
        document.getElementById("search-quick-actions").classList.add("hidden");

        if (result.length == 0) {
          document.getElementById("search-no-results").classList.remove("hidden");
          document.getElementById("search-results").classList.add("hidden");
          return;
        }

        document.getElementById("search-no-results").classList.add("hidden");
        document.getElementById("search-results").classList.remove("hidden");

        doco.search.populateResults(result);
      }, 400);
    },
  },

  setupImages: function () {
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
  },

  setupMobileMenu: function () {
    // Adding the mobile menu toggle
    let toggles = document.querySelectorAll(".toggle-mobile-nav");
    toggles.forEach((toggle) => {
      toggle.addEventListener("click", () => {
        document.getElementById("mobile-menu").classList.toggle("hidden");
        document.querySelector("body").classList.toggle("overflow-hidden");
      });
    });
  },
}

document.addEventListener("DOMContentLoaded", () => {
  doco.search.setup();

  doco.formatCode();
  doco.replaceYear();
  doco.setupImages();
  doco.setupMobileMenu();
});
