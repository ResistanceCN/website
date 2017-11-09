const mix = require('laravel-mix');

mix.options({
    processCssUrls: false,
});

mix.disableNotifications();

// MDL theme
mix.js('./assets-src/mdl/js/app.js', './assets/mdl/js');
mix.sass('./assets-src/mdl/scss/app.scss', './assets/mdl/css');

// AntD theme
// mix.js('./assets-src/ant/js/app.js', './assets/ant/js');
// mix.sass('./assets-src/ant/scss/app.scss', './assets/ant/css');
// mix.copy('./node_modules/font-awesome/fonts/*.*',  './assets/fonts/');
