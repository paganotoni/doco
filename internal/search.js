


// TODO: Keyboard up and down
// TODO: Keyboard enter to navigate to article
document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll(".search-button").forEach(element => {
        element.addEventListener('click', showSearch);
    })

    document.getElementById("search-overlay").addEventListener('click', hideSearch);
    document.getElementById("close-search").addEventListener('click', hideSearch);
    document.getElementById("search-input").addEventListener('keyup', (e) => search(e.target.value));

    document.addEventListener('keydown', (e) => {
        if (e.keyCode === 27) {
            hideSearch();
        }

        if (e.keyCode >= 65 && e.keyCode <= 90) {
            let char = (e.metaKey ? '⌘-' : '') + String.fromCharCode(e.keyCode)
            if (char == "⌘-K") {
                showSearch();
            }
        }
    })
    
    // Loading the index when the page it loaded so the search just uses it.
    fetch('/index.json')
    .then(response => response.json())
    .then(data => {
        window.searchIndex = new Fuse(data, {keys: ["title", "section_name", "content"]});
    }).catch(error => console.error('Error:', error));
});


function showSearch() {
    console.log("show search")
    document.getElementById("search-palette").classList.toggle("hidden");
    document.getElementById("search-input").focus();
}

function hideSearch() { 
    document.getElementById("search-palette").classList.toggle("hidden");
}

let tm = null
function search(searchQuery) {
    if (searchQuery.length == 0) {
        document.getElementById("search-no-results").classList.add("hidden");
        document.getElementById("search-results").classList.add("hidden");
        document.getElementById("search-quick-actions").classList.remove("hidden");

        return
    }

    let fuseOptions = {
        shouldSort: true,
        includeMatches: true,
        threshold: 0.0,
        tokenize: true,
        location: 0,
        distance: 100,
        maxPatternLength: 32,
        minMatchCharLength: 2,
        keys: [
            { name: "title", weight: 0.8 },
            { name: "content", weight: 0.5 },
            { name: "section_name", weight: 0.5 },
        ],
    };

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
    }, 200);
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
            snippet: value.item.contents
        });

        document.getElementById('search-results').innerHTML += output;
    });
}

function render(templateString, data) {
    var conditionalMatches, conditionalPattern, copy;
    conditionalPattern = /\$\{\s*isset ([a-zA-Z]*) \s*\}(.*)\$\{\s*end\s*}/g;
    //since loop below depends on re.lastInxdex, we use a copy to capture any manipulations whilst inside the loop
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
