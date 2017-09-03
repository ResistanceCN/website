import $ from 'jquery';

export function toast(text, timeout) {
    let $el = $('.mdl-js-snackbar');

    $el[0].MaterialSnackbar.showSnackbar({
        message: text,
        timeout: timeout || 2750
    });
}

window.toast = toast;

let testElement = document.createElement('p');
let testElementStyles = {
    width: '100px',
    height: '1px',
    overflowY: 'scroll'
};

for (let i in testElementStyles)
    testElement.style[i] = testElementStyles[i];

document.body.appendChild(testElement);
export const scrollbarWidth = testElement.offsetWidth - testElement.clientWidth;
document.body.removeChild(testElement);

export const windowWidth = () => document.body.offsetWidth + scrollbarWidth;
