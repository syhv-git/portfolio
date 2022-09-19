toggle = document.querySelector('.nav-menu');
menu = document.querySelector('.nav-pill');
links = document.querySelectorAll('.nav-link');

allListeners()

function allListeners() {
    toggle.addEventListener('click', clickMenu)
    links.forEach(e => e.addEventListener('click', clickLink))
}

function clickMenu() {
    menu.classList.toggle('open')
}

function clickLink() {
    if (menu.classList.contains('open')) {
        toggle.click()
    }
}
