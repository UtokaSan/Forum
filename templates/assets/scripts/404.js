window.onload = function() {
    var videos = document.getElementsByTagName('video');
    for (var i = 0; i < videos.length; i++) {
        videos[i].autoplay = true;
        videos[i].load();
    }
};