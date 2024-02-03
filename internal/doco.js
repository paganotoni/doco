document.addEventListener('DOMContentLoaded', () => {
    // Add highligthing to the code blocks
    hljs.highlightAll();
    
    
    // Adding the mobile menu toggle
    let toggles = document.querySelectorAll(".toggle-mobile-nav")
    toggles.forEach(toggle => {
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
    document.querySelectorAll(".search-button").forEach(element => {
        element.addEventListener('click', showSearch);
    })

    document.getElementById("search-overlay").addEventListener('click', hideSearch);
    document.getElementById("close-search").addEventListener('click', hideSearch);
    document.getElementById("search-input").addEventListener('keyup', (e) => search(e.target.value));
    
    // Loading the index when the page it loaded so the search just uses it.
    fetch('/index.json')
    .then(response => response.json())
    .then(data => {
        window.searchIndex = new Fuse(data, {
            threshold: 0.8,
            tokenize: false,
            includeMatches: true,
            maxPatternLength: 32,
            minMatchCharLength: 1,
            location: 80_000,
            keys: ["content","title", "section_name"],
        });

    }).catch(error => console.error('Error:', error));

    let imageContainer = document.querySelector("#zoomed-image-overlay")
    imageContainer.addEventListener('click', () => {
        imageContainer.classList.add("hidden");
    });

    let zoomables = document.querySelectorAll("#htmlcontainer img");
    zoomables.forEach(zoomable => {
        zoomable.classList.add("cursor-zoom-in");
        zoomable.addEventListener('click', () => {        
            imageContainer.classList.remove("hidden");
            imageContainer.querySelector("img").src = zoomable.src;
        });
    });

    document.addEventListener('keydown', (e) => {
        if (e.keyCode === 27) {
            hideSearch();
            // hide image container
            imageContainer.classList.add("hidden");
        }

        if (e.keyCode >= 65 && e.keyCode <= 90) {
            let char = (e.metaKey ? '⌘-' : '') + String.fromCharCode(e.keyCode)
            if (char == "⌘-K") {
                showSearch();
            }
        }
    })

});


function showSearch() {
    document.getElementById("search-palette").classList.remove("hidden");
    document.getElementById("search-input").focus();
}

function hideSearch() { 
    document.getElementById("search-palette").classList.add("hidden");
}

let tm = null
function search(searchQuery) {
    if (searchQuery.length == 0) {
        document.getElementById("search-no-results").classList.add("hidden");
        document.getElementById("search-results").classList.add("hidden");
        document.getElementById("search-quick-actions").classList.remove("hidden");

        return
    }

    if (tm != null) {
        clearTimeout(tm)
    }

    tm = setTimeout(() => {
        var result = searchIndex.search(searchQuery);

        // hide the quick actions
        document.getElementById("search-quick-actions").classList.add("hidden");

        if (result.length == 0) {
            document.getElementById("search-no-results").classList.remove("hidden");
            document.getElementById("search-results").classList.add("hidden");
            return
        }

        document.getElementById("search-no-results").classList.add("hidden");
        document.getElementById("search-results").classList.remove("hidden");


        populateResults(result);
    }, 400);
}

function populateResults(result) {
    let template = document.getElementById('search-result-template').innerHTML;
    document.getElementById('search-results').innerHTML = "";

    result.forEach((value, index) => {
        if (value.item.title == null) {
            return
        }

        var output = render(template, {
            key: index,
            title: value.item.title,
            link: value.item.link,
        });

        document.getElementById('search-results').innerHTML += output;
    });
}

function render(templateString, data) {
    var conditionalMatches, conditionalPattern, copy;
    conditionalPattern = /\$\{\s*isset ([a-zA-Z]*) \s*\}(.*)\$\{\s*end\s*}/g;
    //since loop below depends on re.lastIndex, we use a copy to capture any manipulations whilst inside the loop
    copy = templateString;
    while ((conditionalMatches = conditionalPattern.exec(templateString)) !== null) {
        if (data[conditionalMatches[1]]) {
            //valid key, remove conditionals, leave contents.
            copy = copy.replace(conditionalMatches[0], conditionalMatches[2]);
        } else {
            //not valid, remove entire section
            copy = copy.replace(conditionalMatches[0], '');
        }
    }

    templateString = copy;
    //now any conditionals removed we can do simple substitution
    var key, find, re;
    for (key in data) {
        find = '\\$\\{\\s*' + key + '\\s*\\}';
        re = new RegExp(find, 'g');
        templateString = templateString.replace(re, data[key]);
    }
    return templateString;
}
