function load(path) {
    let head = document.getElementsByTagName('head')[0]
    let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = path
    head.appendChild(script)
}

load('compile.js')
load('/assets/js/particles.js-master/particles.js')
load('/assets/js/particles.js-master/demo/js/app.js')
