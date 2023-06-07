document.addEventListener('DOMContentLoaded', function() {
    var audio = document.getElementById('myAudio');
    var video1 = document.getElementById('myVideo1');
    var video2 = document.getElementById('myVideo2');
    var video3 = document.getElementById('myVideo3');
    var video4 = document.getElementById('myVideo4');
    var video5 = document.getElementById('myVideo5');

    audio.volume = 0.2;
    video2.volume = 0.2;
    video3.volume = 0.2;

    audio.play();
    video1.play();
    video2.play();
    video3.play();
    video4.play();
    video5.play();
});
