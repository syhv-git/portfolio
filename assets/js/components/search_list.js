const initState = {
    loading: false,
    results: [],
    value: '',
}

function handleSubmit(state, action) {
    if (action.type === 'CLICK_SELECTION') {
        this.location = `/#/profile/${action.selection}`
    }
    if (action.type === 'SEARCH') {
        this.location = `/#/search/${action.submit}`
    }
}

function reducer(state, action) {
    switch (action.type) {
        case 'CLEAN_QUERY':
            return initState
        case 'START_SEARCH':
            return { ...state, loading: true, value: action.query }
        case 'FINISH_SEARCH':
            return { ...state, loading: false, results: action.results }
        case 'CLICK_SELECTION':
            handleSubmit(state, action)
            return initState
        case 'SEARCH':
            handleSubmit(state, action)
            return initState
        default:
            throw new Error()
    }
}

function SearchList() {
}