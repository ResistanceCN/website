import $ from 'jquery';

let $snackbar = $('.mdl-js-snackbar')[0];

export function toast(text, timeout) {
    $snackbar.MaterialSnackbar.showSnackbar({
        message: text,
        timeout: timeout || 2750
    });
}

export function transformIcons() {
    $('.md-icon').each(function () {
        let classes = this.className.match(/md-icon-[a-z_]+/gi);
        let className = this.className.replace(/md-icon(-[a-z_]+)?/gi, '');

        this.className = className + ' material-icons';
        this.innerText = classes[classes.length - 1].substr(8);
    })
}

let testElement = document.createElement('p');

testElement.style.cssText = 'width: 100px; height: 1px; overflow-y: scroll;';

document.body.appendChild(testElement);
export const scrollbarWidth = testElement.offsetWidth - testElement.clientWidth;
document.body.removeChild(testElement);
