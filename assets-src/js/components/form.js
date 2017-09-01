import $ from "jquery";

import { toast } from "../utils";

$(function () {
    $('form').each(function () {
        $('<input>').attr('type', 'hidden')
            .addClass('submit-field')
            .appendTo($(this));
    });

    $(document).on('click', 'form [type="submit"]', function () {
        let $this = $(this);

        let name = $this.attr('name');
        let value = $this.val();

        if (name && value) {
            let $field = $this.parents('form').first().find('.submit-field');
            $field.attr('name', name).val(value);
        }
    });

    $(document).on('submit', '.ajax-form', function (event) {
        event.preventDefault();

        let $this = $(this);

        if ($this.is('[data-confirm]') && !window.confirm($this.data('confirm'))) {
            return;
        }

        let $btn = $this.find('[type="submit"]');
        $btn.prop('disabled', true);

        let data = {};

        /* $(this).find('input, select').each(function () {
            let $this = $(this);

            let name = $this.attr('name');
            let val = $this.val();

            if (name && val) data[$this.attr('name')] = val;
        }); */
        data = $this.serialize();

        let promise = $.ajax({
            url: $this.attr('action'),
            type: $this.attr('method'),
            dataType: 'json',
            data: data
        });

        promise.done(function (data) {
            if (data.redirect) {
                if (!data.msg) {
                    // If msg is empty, redirect immediately
                    return location.href = data.redirect;
                }

                setTimeout('location.href="' + data.redirect + '"', 1600);
            } else {
                $btn.prop('disabled', false);
            }

            toast(data.msg);
        });

        promise.fail(function (xhr) {
            try {
                data = JSON.parse(xhr.responseText);
                toast("Request status: " + data.msg);
            } catch (e) {
                toast("Unknown error. Please contact administrators")
            }

            $btn.prop('disabled', false);
        });
    });
});
