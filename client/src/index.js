$(document).ready(function() {
    game = $.ajax({
        url: 'https://unityapi.joycehuang.me/summary',
            type: 'GET',
            dataType: 'json',
            success: game => {
                $('.card-image').append('<a class="gameUrl" target="_blank" href=" ' + game["gameUrl"] + '">');
                $('.gameUrl').append('<img src=' + game["imageUrl"] + '></a>');
                $('.card-image').append('<span class="card-title">' + game["title"] + '</span>');
                $('.card-content').append('<p id="desc">' + game["description"] + '</p>');
                console.log(game["title"]);
            },
            beforeSend: function(r) {
                r.setRequestHeader(
                    "Authorization",
                    JSON.parse(localStorage.getItem("auth"))
                );
            },
            error: function(xhr) {
                $('.heading').append('<h3 class="message">' + xhr.responseText + '</h3>');
            }
    })
});