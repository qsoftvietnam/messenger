new Vue({
    el: '#app',

    data: {
        gravatar: 'http://www.gravatar.com/avatar/',
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        chatHistory: '', // A running list of chat history messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">' +
                '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                +
                msg.username +
                '</div>' +
                emojione.toImage(msg.message) + '<br/>'; // Parse emojis

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });

        // Textarea is editable only when socket is opened.
        this.ws.onopen = function(e) {
            Materialize.toast('WebSocket server have opened', 2000);
        };

        this.ws.onclose = function(e) {
            Materialize.toast('WebSocket server have been closed', 2000);
        };
    },

    methods: {
        history: function() {
            var self = this;
            this.$http.get('/join').then(response => {
                // get body data
                var result = response.data;
                if (result != null) {
                    if (result.length > 0) {
                        self.chatContent = '';

                        result.forEach(function(msg) {
                            self.chatContent += '<div class="chip">' +
                                '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                                +
                                msg.username +
                                '</div>' +
                                emojione.toImage(msg.message) + '<br/>'; // Parse emojis
                        })

                        var element = document.getElementById('chat-messages');
                        element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                    }
                }
            }, response => {
                Materialize.toast('Have error happen in process', 2000);
            });
        },

        send: function() {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }));

                // Reset newMsg
                this.newMsg = '';
            }
        },

        join: function() {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }

            if (!this.validateEmail(this.email)) {
                Materialize.toast('Email is invalid', 2000);
                return
            }

            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }

            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();

            this.ws.send(
                JSON.stringify({
                    email: this.email,
                    username: this.username
                }));

            this.joined = true;
            this.history();
        },

        gravatarURL: function(email) {
            return this.gravatar + CryptoJS.MD5(email);
        },

        validateEmail: function (email) {
            var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return re.test(email);
        }
    }
});
