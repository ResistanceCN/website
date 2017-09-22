import $ from 'jquery';
import { toast } from "../utils";

if (window.gapi) {
    window.gapi.load('auth2', () => {
        window.gapi.auth2.init({
            client_id: window.config.google_client_id
        });
        window.gapi.signin2.render('google-signin-button', {
            scope: 'profile email',
            height: 36,
            longtitle: true,
            theme: 'dark',
            onsuccess(user) {
                console.log(user);
                login(user.getAuthResponse().id_token);
            },
            onfailure(error) {
                console.error(error);
                toast('Error signing in Google account!');
            }
        });
    });
}

function login(id_token) {
    let promise = $.ajax({
        url: '/login',
        type: 'POST',
        dataType: 'json',
        data: { id_token }
    });

    promise.done(user => {
        console.log(user);

        if (user.id) {
            toast('Login successful');
            location.href = '/';
        } else {
            $('.auth-card-title').text('Create your profile');

            $('#email').val(user.email)
                .parent('.mdl-textfield')
                .addClass('is-dirty');

            $('#id_token').val(id_token);

            $('#google-signin').hide();
            $('#register-form').fadeIn(100);
        }
    });

    promise.fail(function (xhr) {
        toast('Failed to login: ' + xhr.status);
        console.error(xhr.responseText);
    });
}
