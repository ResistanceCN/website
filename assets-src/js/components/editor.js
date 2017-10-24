import $ from 'jquery';
import SimpleMDE from 'simplemde';

let $editor = $('#markdown-editor');

if ($editor.length !== 0) {
    let editor = new SimpleMDE({
        element: $editor[0],
        autoDownloadFontAwesome: false,
        indentWithTabs: false,
        tabSize: 4,
        promptURLs: true,
        styleSelectedText: false,
        toolbar: [
            'bold',
            'italic',
            'strikethrough',
            '|',
            'heading',
            'horizontal-rule',
            '|',
            'code',
            'quote',
            'unordered-list',
            'ordered-list',
            '|',
            'link',
            'image',
            'table',
            '|',
            'undo',
            'redo'
        ],
    });

    editor.codemirror.on('change', instance => {
        findBlocksAndAddClass(instance, '<!--', '-->', 'html-comments');
        findBlocksAndAddClass(instance, '```', '```', 'md-code-block');
    });
    findBlocksAndAddClass(editor.codemirror, '<!--', '-->', 'html-comments');
    findBlocksAndAddClass(editor.codemirror, '```', '```', 'md-code-block');

    updateToolbarPos();
    $(window).resize(updateToolbarPos);
    $(window).scroll(updateToolbarPos);
    $('.mdl-layout').scroll(updateToolbarPos);
}

function updateToolbarPos() {
    let rect = $('.editor')[0].getBoundingClientRect();

    if (rect.y <= 0) {
        $('.editor-toolbar').addClass('mdl-shadow--2dp fixed').css({
            width: rect.width + 'px',
            left: rect.x + 'px'
        });
    } else {
        $('.editor-toolbar').removeClass('mdl-shadow--2dp fixed').css({
            width: '100%',
            left: 0
        });
    }
}

function findBlocksAndAddClass(codemirror, start, end, className) {
    // Indicate if the line is a part of a block in the iteration
    let inBlock = false;

    codemirror.eachLine(line => {
        // Determine if a start token is found (may be the same with end token)
        let blockStart = line.text.startsWith(start);

        if (!inBlock && !blockStart) {
            codemirror.removeLineClass(line, 'text', className);
            return;
        }

        codemirror.addLineClass(line, 'text', className);

        // If inBlock is false, the line must be the start line of a block
        // Therefore, we should search for the end token with an offset because two token might be identical
        // If inBlock is true, the line is not the start line, then we can search in whole line
        let searchPosition = (!inBlock) ? start.length : 0;

        // If the line is the end line of a block, inBlock will be set to false
        inBlock = line.text.indexOf(end, searchPosition) === -1;
    });
}
