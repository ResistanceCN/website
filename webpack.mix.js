const mix = require('laravel-mix');

mix.options({
    processCssUrls: false,
});

mix.disableNotifications();
mix.js('assets-src/js/app.js', 'assets/js');
mix.sass('assets-src/scss/app.scss', 'assets/css');
// mix.copy('./node_modules/font-awesome/fonts/*.*',  './assets/fonts/');
