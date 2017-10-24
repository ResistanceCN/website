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
    let inBlock = false;

    codemirror.eachLine(line => {
        let blockStart = line.text.startsWith(start);

        if (!inBlock && !blockStart) {
            codemirror.removeLineClass(line, 'text', className);
            return;
        }

        codemirror.addLineClass(line, 'text', className);

        let searchStart = (!inBlock) ? start.length : 0;
        inBlock = line.text.indexOf(end, searchStart) === -1;
    });
}
