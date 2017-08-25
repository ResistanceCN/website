const mix = require('laravel-mix');

mix.disableNotifications();
mix.js('assets-src/js/app.js', 'assets/js');
mix.sass('assets-src/scss/app.scss', 'assets/css');
