import $ from 'jquery';

let $snackbar = $('.mdl-js-snackbar')[0];

export function toast(text, timeout) {
    $snackbar.MaterialSnackbar.showSnackbar({
        message: text,
        timeout: timeout || 2750
    });
}

let testElement = document.createElement('p');

testElement.style.cssText = 'width: 100px; height: 1px; overflow-y: scroll;';

document.body.appendChild(testElement);
export const scrollbarWidth = testElement.offsetWidth - testElement.clientWidth;
document.body.removeChild(testElement);
