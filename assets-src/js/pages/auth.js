import $ from 'jquery';

if (window.gapi) {
    window.gapi.load('auth2', () => {
        window.gapi.auth2.init({
            client_id: '{{ google_client_id }}'
        });
        window.gapi.signin2.render('g-signin-button', {
            scope: 'profile email',
            height: 36,
            longtitle: true,
            theme: 'dark',
            onsuccess(user) {
                // TODO
            },
            onfailure(error) {
                // TODO
            }
        });
    });
}
