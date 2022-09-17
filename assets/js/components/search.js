search = document.getElementById("search-icon")
bar = document.getElementById("search-bar")
const sicon = '<img class="nav-search-icon" alt="search" src="/assets/images/search-icon.png"/>'
const cicon = '<img class="nav-search-icon" alt="close" src="/assets/images/close.png"/>'

search.innerHTML = sicon
let b = true

search.addEventListener("click", function () {
    b ? search.innerHTML = cicon : search.innerHTML = sicon
    b ? bar.style.display = "flex" : bar.style.display = "none"
    b = !b
})
