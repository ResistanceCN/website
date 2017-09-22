import $ from 'jquery';
import Masonry from 'masonry-layout';

$(function () {
    let $grid = $('.grid');

    if (!$grid.length)
        return;

    const MIN_WIDTH = 320;

    function resetWidth() {
        let totalWidth = $('.grid').width();

        let n = Math.floor(totalWidth / MIN_WIDTH);

        if (n <= 1) {
            return $('.grid-item').css('width', totalWidth + 'px');
        }

        let boxWidth = Math.floor(totalWidth / n);

        $('.grid-item').css('width', boxWidth + 'px');
    }

    $(window).resize(resetWidth);
    resetWidth();

    new Masonry($grid[0], {
        itemSelector: '.grid-item',
        columnWidth: '.grid-item'
    });
});
