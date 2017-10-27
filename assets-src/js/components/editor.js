import $ from 'jquery';
import SimpleMDE from 'simplemde';
import { transformIcons } from '../utils';

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
            {
                name: "bold",
                action: SimpleMDE.toggleBold,
                className: "md-icon md-icon-format_bold",
                title: "Bold",
            },
            {
                name: "italic",
                action: SimpleMDE.toggleItalic,
                className: "md-icon md-icon-format_italic",
                title: "Italic",
            },
            {
                name: "strikethrough",
                action: SimpleMDE.toggleStrikethrough,
                className: "md-icon md-icon-format_strikethrough",
                title: "Strikethrough"
            },
            '|',
            {
                name: "heading",
                action: SimpleMDE.toggleHeadingSmaller,
                className: "md-icon md-icon-text_fields",
                title: "Heading",
            },
            {
                name: "horizontal-rule",
                action: SimpleMDE.drawHorizontalRule,
                className: "md-icon md-icon-remove",
                title: "Insert Horizontal Line"
            },
            '|',
            {
                name: "code",
                action: SimpleMDE.toggleCodeBlock,
                className: "md-icon md-icon-code",
                title: "Code"
            },
            {
                name: "quote",
                action: SimpleMDE.toggleBlockquote,
                className: "md-icon md-icon-format_quote",
                title: "Quote",
            },
            {
                name: "unordered-list",
                action: SimpleMDE.toggleUnorderedList,
                className: "md-icon md-icon-format_list_bulleted",
                title: "Generic List",
            },
            {
                name: "ordered-list",
                action: SimpleMDE.toggleOrderedList,
                className: "md-icon md-icon-format_list_numbered",
                title: "Numbered List",
            },
            '|',
            {
                name: "link",
                action: SimpleMDE.drawLink,
                className: "md-icon md-icon-link",
                title: "Create Link",
            },
            {
                name: "image",
                action: SimpleMDE.drawImage,
                className: "md-icon md-icon-insert_photo",
                title: "Insert Image",
            },
            {
                name: "table",
                action: SimpleMDE.drawTable,
                className: "md-icon md-icon-border_all",
                title: "Insert Table"
            },
            '|',
            {
                name: "undo",
                action: SimpleMDE.undo,
                className: "md-icon md-icon-undo no-disable",
                title: "Undo"
            },
            {
                name: "redo",
                action: SimpleMDE.redo,
                className: "md-icon md-icon-redo no-disable",
                title: "Redo"
            }
        ]
    });

    const blocks = [
        {
            name: 'code',
            startToken: '```',
            endToken: '```'
        },
        {
            name: 'comment',
            startToken: '<!--',
            endToken: '-->'
        }
    ];

    editor.codemirror.on('change', instance => {
        findBlocksAndAddClass(instance, blocks);
    });
    findBlocksAndAddClass(editor.codemirror, blocks);

    updateToolbarPos();
    $(window).resize(updateToolbarPos);
    $(window).scroll(updateToolbarPos);
    $('.mdl-layout').scroll(updateToolbarPos);

    transformIcons();
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

function findBlocksAndAddClass(codeMirror, blocks) {
    for (let block of blocks) {
        const className = 'block-' + block.name;

        // Indicate if the line is a part of a block in the iteration
        let inBlock = false;

        codeMirror.eachLine(line => {
            // Determine if a start token is found (may be the same with end token)
            let blockStart = line.text.startsWith(block.startToken);

            if (!inBlock && !blockStart) {
                codeMirror.removeLineClass(line, 'text', className);
                return;
            }

            codeMirror.removeLineClass(line, 'text');
            codeMirror.addLineClass(line, 'text', className);

            let searchPosition = 0;

            if (!inBlock) {
                // If inBlock is false, the line must be the start line of a block
                // Therefore, we should search for the end token with an offset because two token might be identical
                // If inBlock is true, the line is not the start line, then we can search in whole line
                searchPosition = block.startToken.length;
                codeMirror.addLineClass(line, 'text', 'block-start');
            }

            if (line.text.indexOf(block.endToken, searchPosition) === -1) {
                inBlock = true;
            } else {
                // If the line is the end line of a block, inBlock will be set to false
                inBlock = false;
                codeMirror.addLineClass(line, 'text', 'block-end');
            }
        });
    }
}
