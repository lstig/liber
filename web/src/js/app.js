window.htmx = require('htmx.org');

// disable making HX requests when the user hits the browser back button (instead the entire page will be returned)
window.htmx.config.historyRestoreAsHxRequest = false