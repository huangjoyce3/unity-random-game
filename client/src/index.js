$(document).ready(function() {
    game = $.ajax({
        url: 'https://unityapi.joycehuang.me/summary',
            type: 'GET',
            dataType: 'json',
            success: game => {
                $('.card-image').append('<a class="gameUrl" target="_blank" href=" ' + game["gameUrl"] + '">');
                $('.gameUrl').append('<img src=' + game["imageUrl"] + '></a>');
                $('.card-image').append('<span class="card-title">' + game["title"] + '</span>');
                // $('.card-image').append('<span class="">' + game["developer"] + '</span>');
                $('.card-content').append('<p id="desc">' + game["description"] + '</p>');
            },
            beforeSend: function(r) {
                r.setRequestHeader(
                    "Authorization",
                    JSON.parse(localStorage.getItem("auth"))
                );
            },
            error: function(xhr) {
                $('#message').append('<h3 id="message">' + xhr.responseText + '</h3>');

                setTimeout(function() {
                    $('h3').remove();
                }, 2500);
            }
    })
    // fetch("https://unityapi.joycehuang.me/summary")
    //     .then(function(response) {
    //         return response.text();
    //     })
    //     .then(function(data) {
    //         var obj = JSON.parse(data)
    //         document.getElementById("title").innerHTML = obj.title;
    //         document.getElementById("desc").innerHTML = obj.description;
    //         document.getElementById("img").src = obj.image;
    //     })
});