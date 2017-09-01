import $ from 'jquery';

export function toast(text, timeout) {
    let $el = $('.mdl-js-snackbar');

    $el[0].MaterialSnackbar.showSnackbar({
        message: text,
        timeout: timeout || 2750
    });
}

window.toast = toast;
