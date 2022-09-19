function load(path) {
    let head = document.getElementsByTagName('head')[0]
    let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = path
    head.appendChild(script)
}

load('/assets/js/components/navbar.js')
load('/assets/js/components/search.js')
load('/assets/js/components/search_list.js')
load('/assets/js/particles.js-master/particles.js')
load('/assets/js/particles.js-master/demo/js/app.js')
